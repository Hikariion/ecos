package watcher

import (
	"context"
	"ecos/edge-node/infos"
	"ecos/messenger"
	"ecos/messenger/common"
	"ecos/utils/logger"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Monitor interface {
	MonitorServer
	Run()
	GetAllNodeReports() []*NodeStatusReport
	GetNodeReport(nodeId uint64) *NodeStatusReport
	GetEventChannel() <-chan *Event
	Register(name string, reporter Reporter) error
	stop()
}

type Event struct {
	Report *NodeStatusReport
}

type ReportType int32

const (
	ReportTypeADD ReportType = iota
	ReportTypeUPDATE
	ReportTypeDELETE
)

type Report struct {
	ReportType     ReportType
	NodeReport     *NodeStatusReport
	PipelineReport *PipelineReport
}

type Reporter interface {
	IsChanged() bool
	GetReports() []Report
}

type NodeMonitor struct {
	UnimplementedMonitorServer
	ctx    context.Context
	cancel context.CancelFunc
	timer  *time.Ticker

	nodeStatusMap sync.Map
	reportTimers  sync.Map

	selfNodeStatus *NodeStatusReport
	selfPipeline   map[uint64]*PipelineReport
	watcher        *Watcher

	reportersMap sync.Map
	eventChannel chan *Event
}

func (m *NodeMonitor) Register(name string, reporter Reporter) error {
	if _, ok := m.reportersMap.Load(name); ok {
		return errors.New("reporter already registered")
	}
	m.reportersMap.Store(name, reporter)
	return nil
}

// Report is a rpc func to get the node status report.
// it called by all node (include leader self).
func (m *NodeMonitor) Report(_ context.Context, report *NodeStatusReport) (*common.Result, error) {
	if !m.watcher.GetMoon().IsLeader() {
		// only leader can be runReport
		logger.Infof("node %v is not leader, can't runReport", m.watcher.GetSelfInfo().GetID())
		return nil, errors.New("not leader")
	}
	if val, ok := m.reportTimers.Load(report.NodeId); ok {
		logger.Tracef("reset timer for node %v", report.NodeId)
		t := val.(*time.Timer)
		if !t.Stop() {
			<-t.C
		}
		t.Reset(time.Second * 3)
	} else {
		logger.Debugf("create timer for node %v", report.NodeId)
		m.reportTimers.Store(report.NodeId, time.AfterFunc(time.Second*3, func() {
			select {
			case <-m.ctx.Done():
				return
			default:
			}
			logger.Warningf("get node status timeout, nodeId: %v", report.NodeId)
			v, _ := m.nodeStatusMap.Load(report.NodeId)
			r := v.(*NodeStatusReport)
			r.State = infos.NodeState_OFFLINE
			m.nodeStatusMap.Store(report.NodeId, r)
			m.eventChannel <- &Event{
				Report: r,
			}
		}))
	}
	if _, ok := m.nodeStatusMap.Load(report.NodeId); !ok {
		// first time online
		m.eventChannel <- &Event{
			Report: report,
		}
	}
	m.nodeStatusMap.Store(report.NodeId, report)
	return &common.Result{}, nil
}

// GetEventChannel returns the event channel.
// Event channel will send event when the node status changed. (like node online, offline, etc.)
func (m *NodeMonitor) GetEventChannel() <-chan *Event {
	return m.eventChannel
}

// GetAllNodeReports returns all node status reports.
func (m *NodeMonitor) GetAllNodeReports() []*NodeStatusReport {
	var nodeStatusList []*NodeStatusReport
	m.nodeStatusMap.Range(func(key, value interface{}) bool {
		nodeStatusList = append(nodeStatusList, value.(*NodeStatusReport))
		return true
	})
	sort.Slice(nodeStatusList, func(i, j int) bool {
		return nodeStatusList[i].NodeId < nodeStatusList[j].NodeId
	})
	return nodeStatusList
}

func (m *NodeMonitor) GetNodeReport(nodeID uint64) *NodeStatusReport {
	if val, ok := m.nodeStatusMap.Load(nodeID); ok {
		return val.(*NodeStatusReport)
	}
	return nil
}

func (m *NodeMonitor) GetClusterReport(context.Context, *emptypb.Empty) (*ClusterReport, error) {
	reports := m.GetAllNodeReports()
	return &ClusterReport{
		Nodes: reports,
	}, nil
}

func (m *NodeMonitor) collectReports() {
	m.reportersMap.Range(func(key, value interface{}) bool {
		reporter := value.(Reporter)
		if reporter.IsChanged() == false {
			return true
		}
		reports := reporter.GetReports()
		for _, report := range reports {
			if report.NodeReport != nil {
				m.selfNodeStatus = report.NodeReport
			}
			if report.PipelineReport != nil {
				if report.ReportType == ReportTypeDELETE {
					delete(m.selfPipeline, report.PipelineReport.PgId)
				} else {
					m.selfPipeline[report.PipelineReport.PgId] = report.PipelineReport
				}
			}
		}
		return true
	})
}

func (m *NodeMonitor) runReport() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-m.timer.C:
			m.collectReports()
			leaderID := m.watcher.GetMoon().GetLeaderID()
			if leaderID == 0 {
				continue
			}
			leaderInfo, err := m.watcher.GetMoon().GetInfoDirect(infos.InfoType_NODE_INFO, strconv.FormatUint(leaderID, 10))
			if err != nil || leaderInfo.BaseInfo().GetNodeInfo() == nil {
				logger.Warningf("node: %v get leader info: %v failed: %v", m.watcher.GetSelfInfo().GetID(), leaderID, err)
				continue
			}
			conn, _ := messenger.GetRpcConnByNodeInfo(leaderInfo.BaseInfo().GetNodeInfo())
			client := NewMonitorClient(conn)
			_, err = client.Report(m.ctx, m.selfNodeStatus)
			if err != nil {
				logger.Errorf("runReport node status failed: %v", err)
			}
		}
	}
}

func (m *NodeMonitor) Run() {
	m.timer = time.NewTicker(time.Second * 1)
	go m.runReport()
}

func (m *NodeMonitor) stop() {
	m.cancel()
	m.timer.Stop()
}

func NewMonitor(ctx context.Context, w *Watcher, rpcServer *messenger.RpcServer) Monitor {
	ctx, cancel := context.WithCancel(ctx)
	monitor := &NodeMonitor{
		ctx:           ctx,
		cancel:        cancel,
		nodeStatusMap: sync.Map{},
		watcher:       w,
		eventChannel:  make(chan *Event),
		selfPipeline:  make(map[uint64]*PipelineReport),
	}
	RegisterMonitorServer(rpcServer, monitor)
	return monitor
}

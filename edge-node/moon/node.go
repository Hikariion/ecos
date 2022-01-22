package moon

import (
	"context"
	"ecos/utils/logger"
	"encoding/json"
	"github.com/coreos/etcd/raft/raftpb"
	"go.etcd.io/etcd/Godeps/_workspace/src/github.com/golang/glog"
	"go.etcd.io/etcd/raft"
	"time"
)

type node struct {
	id          uint64 //raft节点的id
	selfInfo    *NodeInfo
	ctx         context.Context //context
	infoStorage NodeInfoStorage
	raftStorage *raft.MemoryStorage //raft需要的内存结构
	cfg         *raft.Config        //raft需要的配置
	raft        raft.Node           // raft node
	ticker      <-chan time.Time    //定时器，提供周期时钟源和超时触发能力
	recv        chan raftpb.Message
}

var (
	bcChans = []chan raftpb.Message{
		make(chan raftpb.Message),
		make(chan raftpb.Message),
		make(chan raftpb.Message),
		make(chan raftpb.Message),
	}
)

func NewNode(selfInfo *NodeInfo, leaderInfo *NodeInfo, groupInfo []*NodeInfo) *node {
	ctx := context.TODO()
	storage := raft.NewMemoryStorage()
	infoStorage := NewMemoryNodeInfoStorage()
	cfg := raft.Config{
		ID:              uint64(selfInfo.ID),
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	id := selfInfo.ID

	n := &node{
		id:          uint64(selfInfo.ID),
		selfInfo:    selfInfo,
		ctx:         ctx,
		infoStorage: infoStorage,
		raftStorage: storage,
		cfg:         &cfg,
		ticker:      time.Tick(time.Millisecond),
		recv:        bcChans[id-1],
	}

	var peers []raft.Peer
	for _, nodeInfo := range groupInfo {
		peers = append(peers, raft.Peer{
			ID:      uint64(nodeInfo.ID),
			Context: nil,
		})
	}

	n.raft = raft.StartNode(n.cfg, peers)
	return n
}

func (n *node) send(messages []raftpb.Message) {
	for _, m := range messages {
		glog.Infof(raft.DescribeMessage(m, nil))
		to := m.To
		ch := bcChans[to-1]
		glog.Infof("%d send to %v, type %v", n.id, m.To, m.Type)
		ch <- m
	}
}

func (n *node) process(entry raftpb.Entry) {
	if entry.Type == raftpb.EntryNormal && entry.Data != nil {
		var nodeInfo NodeInfo
		err := json.Unmarshal(entry.Data, &nodeInfo)
		if err != nil {
			logger.Errorf("Moon process nodeInfo err: %v", err.Error())
		}
		logger.Infof("Node %v: get node info %v", n.id, nodeInfo)
		_ = n.infoStorage.UpdateNodeInfo(&nodeInfo)
	}
}

func (n *node) run() {
	for {
		select {
		case <-n.ticker:
			n.raft.Tick()
		case rd := <-n.raft.Ready():
			_ = n.raftStorage.Append(rd.Entries)
			go n.send(rd.Messages)
			for _, entry := range rd.CommittedEntries {
				n.process(entry)
				if entry.Type == raftpb.EntryConfChange {
					var cc raftpb.ConfChange
					_ = cc.Unmarshal(entry.Data)
					n.raft.ApplyConfChange(cc)
				}
			}
			n.raft.Advance()
		case m := <-n.recv:
			glog.Infof("%d got message from %v to %v, type %v", n.id, m.From, m.To, m.Type)
			_ = n.raft.Step(n.ctx, m)
			glog.Infof("%d status: %v", n.id, n.raft.Status().RaftState)
		}
	}
}

func (n *node) reportSelfInfo() {
	info := n.selfInfo
	js, _ := json.Marshal(info)
	err := n.raft.Propose(n.ctx, js)
	if err != nil {
		logger.Errorf("report self info err: %v", err.Error())
	}
}

package gaia

import (
	"context"
	"ecos/edge-node/watcher"
	"ecos/messenger"
	"ecos/messenger/common"
	"ecos/utils/errno"
	"ecos/utils/logger"
	"io"
)

// Gaia save object block data
// 盖亚存储和查询对象 block 数据
// 地势坤，厚德载物
type Gaia struct {
	UnimplementedGaiaServer

	ctx     context.Context
	cancel  context.CancelFunc
	watcher *watcher.Watcher

	config *Config
}

func (g *Gaia) UploadBlockData(stream Gaia_UploadBlockDataServer) error {
	transporter := &PrimaryCopyTransporter{}
	for {
		select {
		case <-g.ctx.Done():
			return stream.SendAndClose(&common.Result{
				Status:  common.Result_FAIL,
				Code:    errno.CodeGaiaClosed,
				Message: errno.GaiaClosedErr.Error(),
			})
		default:
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch payload := r.Payload.(type) {
		case *UploadBlockRequest_Message:
			err = g.processControlMessage(payload, transporter, stream)
		case *UploadBlockRequest_Chunk:
			err = g.processChunk(payload, transporter, stream)
		case nil:
			// The field is not set.
			logger.Warningf("Received blank Payload")
			return stream.SendAndClose(&common.Result{
				Status: common.Result_FAIL,
			})
		default:
			logger.Errorf("UploadBlockRequest.Payload has unexpected type %T", payload)
			return stream.SendAndClose(&common.Result{
				Status: common.Result_FAIL,
			})
		}
		if err != nil {
			return err
		}
	}
}

// processControlMessage will modify transporter when receive ControlMessage_BEGIN
func (g *Gaia) processControlMessage(message *UploadBlockRequest_Message, transporter *PrimaryCopyTransporter,
	stream Gaia_UploadBlockDataServer) (err error) {
	msg := message.Message
	code := msg.Code
	p := msg.Pipeline
	switch code {
	case ControlMessage_BEGIN:
		// 建立与同组 Node 的连接，准备转发
		clusterInfo, err := g.watcher.GetClusterInfoByTerm(msg.Term)
		if err != nil {
			return err
		}
		t, err := NewPrimaryCopyTransporter(g.ctx, msg.Block, p, g.watcher.GetSelfInfo().RaftId,
			&clusterInfo, g.config.BasePath)
		if err != nil {
			return err
		}
		*transporter = *t
		logger.Infof("Gaia start receive block: %v", msg.Block.BlockId)
	case ControlMessage_EOF:
		// 确认转发成功，关闭连接
		err := transporter.Close()
		if err != nil {
			return stream.SendAndClose(&common.Result{
				Status:  common.Result_FAIL,
				Message: err.Error(),
			})
		}
		logger.Infof("Gaia save block: %v success", msg.Block.BlockId)
		return stream.SendAndClose(&common.Result{
			Status: common.Result_OK,
		})
	default:
		logger.Errorf("ControlMessage has unexpected code %v", code)
		return stream.SendAndClose(&common.Result{
			Status: common.Result_FAIL,
		})
	}
	return nil
}

func (g *Gaia) processChunk(chunk *UploadBlockRequest_Chunk, transporter *PrimaryCopyTransporter,
	stream Gaia_UploadBlockDataServer) (err error) {
	data := chunk.Chunk.Content
	if transporter == nil {
		return stream.SendAndClose(&common.Result{
			Status:  common.Result_FAIL,
			Code:    errno.NoTransporterErr.Code,
			Message: errno.NoTransporterErr.Error(),
		})
	}
	_, err = transporter.Write(data)
	if err != nil {
		return stream.SendAndClose(&common.Result{
			Status:  common.Result_FAIL,
			Code:    errno.CodeTransporterWriteFail,
			Message: errno.TransporterWriteFail.Error() + err.Error(),
		})
	}
	return nil
}

func NewGaia(ctx context.Context, rpcServer *messenger.RpcServer, watcher *watcher.Watcher,
	config *Config) *Gaia {
	ctx, cancel := context.WithCancel(ctx)
	g := Gaia{
		ctx:     ctx,
		cancel:  cancel,
		watcher: watcher,
		config:  config,
	}
	RegisterGaiaServer(rpcServer, &g)
	return &g
}

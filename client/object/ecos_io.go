package object

import (
	"context"
	"crypto/sha256"
	"ecos/client/config"
	info_agent "ecos/client/info-agent"
	clientNode "ecos/client/node"
	"ecos/edge-node/infos"
	"ecos/edge-node/object"
	"ecos/edge-node/pipeline"
	"ecos/edge-node/watcher"
	"ecos/messenger"
	"ecos/utils/common"
	"ecos/utils/logger"
	"io"
)

// EcosIOFactory Generates EcosWriter with ClientConfig
type EcosIOFactory struct {
	infoAgent   *info_agent.InfoAgent
	config      *config.ClientConfig
	clusterInfo *infos.ClusterInfo
	objPipes    []*pipeline.Pipeline
	blockPipes  []*pipeline.Pipeline
}

// NewEcosIOFactory Constructor for EcosIOFactory
//
// nodeAddr shall provide the address to get ClusterInfo from
func NewEcosIOFactory(config *config.ClientConfig) *EcosIOFactory {
	conn, err := messenger.GetRpcConn(config.NodeAddr, config.NodePort)
	if err != nil {
		return nil
	}
	watcherClient := watcher.NewWatcherClient(conn)
	reply, err := watcherClient.GetClusterInfo(context.Background(),
		&watcher.GetClusterInfoRequest{Term: 0})
	clusterInfo := reply.GetClusterInfo()
	if err != nil {
		logger.Errorf("get group info fail: %v", err)
		return nil
	}
	// TODO: Retry?
	clientNode.InfoStorage.SaveClusterInfoWithTerm(0, clusterInfo)
	const groupNum = 3
	// TODO: Get pgNum, groupNum from moon
	ret := &EcosIOFactory{
		infoAgent:   info_agent.NewInfoAgent(context.Background(), clusterInfo),
		clusterInfo: clusterInfo,
		config:      config,
		objPipes:    pipeline.GenPipelines(*clusterInfo, objPgNum, groupNum),
		blockPipes:  pipeline.GenPipelines(*clusterInfo, blockPgNum, groupNum),
	}
	return ret
}

func (f *EcosIOFactory) newLocalChunk() (io.Closer, error) {
	return &localChunk{freeSize: f.config.Object.ChunkSize}, nil
}

// GetEcosWriter provide a EcosWriter for object associated with key
func (f *EcosIOFactory) GetEcosWriter(key string) EcosWriter {
	maxChunk := uint(f.config.UploadBuffer / f.config.Object.ChunkSize)
	chunkPool, _ := common.NewPool(f.newLocalChunk, maxChunk, int(maxChunk))
	return EcosWriter{
		clusterInfo: f.clusterInfo,
		key:         key,
		config:      f.config,
		Status:      READING,
		chunks:      chunkPool,
		blocks:      map[int]*Block{},
		blockPipes:  f.blockPipes,
		meta: &object.ObjectMeta{
			ObjId:      "",
			ObjSize:    0,
			UpdateTime: nil,
			ObjHash:    "",
			PgId:       0,
			Blocks:     []*object.BlockInfo{},
		},
		objHash:        sha256.New(),
		objPipes:       f.objPipes,
		finishedBlocks: make(chan *Block),
	}
}

// GetEcosReader provide a EcosWriter for object associated with key
func (f *EcosIOFactory) GetEcosReader(key string) *EcosReader {
	maxChunkId := f.config.Object.BlockSize/f.config.Object.ChunkSize - 1
	return &EcosReader{
		infoAgent:         f.infoAgent,
		clusterInfo:       f.clusterInfo,
		key:               key,
		blockPipes:        nil,
		curBlockId:        0,
		meta:              nil,
		objPipes:          f.objPipes,
		curChunkIdInBlock: 0,
		maxChunkIdInBlock: maxChunkId,
		curChunk:          0,
		sizeOfChunk:       f.config.Object.ChunkSize,
		chunkOffset:       0,
		alreadyReadBytes:  0,
	}
}

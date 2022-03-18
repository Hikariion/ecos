package object

import (
	"ecos/client/config"
	"ecos/edge-node/alaya"
	"ecos/edge-node/gaia"
	"ecos/edge-node/moon"
	"ecos/edge-node/node"
	"ecos/messenger"
	"ecos/utils/common"
	"ecos/utils/timestamp"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net"
	"path"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestEcosWriter(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)
	type args struct {
		objectSize int
		nodeAddr   string
		port       int
		key        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"writer 8M object",
			args{
				1024 * 1024 * 8, // 8M
				"127.0.0.1",
				32801,
				"/path/8M_obj",
			},
			false,
		},
		{"writer 8.1M object",
			args{
				1024*1024*8 + 1024*100, // 8.1M
				"127.0.0.1",
				32801,
				"/path/8.1M_obj",
			},
			false,
		},
		{"writer 128M object",
			args{
				1024 * 1024 * 128, // 128M
				"127.0.0.1",
				32801,
				"/path/128M_obj",
			},
			false,
		},
	}
	basePath := "./ecos-data/"
	_ = common.InitAndClearPath(basePath)
	infos, moons, alayas, rpcServers, err := createServers(9, "", path.Join(basePath, "db", "moon"))
	if err != nil {
		t.Errorf("RpcServer error = %v", err)
	}
	for i := 0; i < 9; i++ {
		infoStorage := moons[i].InfoStorage
		gaia.NewGaia(rpcServers[i], infos[i], infoStorage,
			&gaia.Config{BasePath: path.Join(basePath, "gaia", strconv.Itoa(i+1))})
		i := i
		go func(rpc *messenger.RpcServer) {
			err := rpc.Run()
			if err != nil {
				t.Errorf("GaiaServer error = %v", err)
			}
		}(rpcServers[i])
		go moons[i].Run()
		go alayas[i].Run()
	}

	time.Sleep(10 * time.Second) // ensure rpcServer running

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := NewEcosWriterFactory(config.DefaultConfig, tt.args.nodeAddr, tt.args.port)
			writer := factory.GetEcosWriter(tt.args.key)
			var data = make([]byte, tt.args.objectSize)
			rand.Read(data)
			assert.NoError(t, err)
			writeSize, err := writer.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, tt.args.objectSize, writeSize)
			assert.NoError(t, writer.Close())
			t.Logf("Upload Finish!")
			if (err != nil) != tt.wantErr {
				t.Errorf("PutObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			time.Sleep(1 * time.Second)
		})
	}
}

func TestPortClose(t *testing.T) {
	for port := 32801; port < 32804; port++ {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:"+strconv.Itoa(port), time.Second)
		if err == nil && conn != nil {
			t.Errorf("port %v not close!", port)
		}
	}
	t.Log("All ports closed!")
}

func createServers(num int, sunAddr string, basePath string) ([]*node.NodeInfo, []*moon.Moon, []*alaya.Alaya, []*messenger.RpcServer, error) {
	err := common.InitAndClearPath(basePath)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	var infoStorages []node.InfoStorage
	var stableStorages []moon.Storage
	var rpcServers []*messenger.RpcServer
	var moons []*moon.Moon
	var nodeInfos []*node.NodeInfo
	var alayas []*alaya.Alaya

	for i := 0; i < num; i++ {
		raftID := uint64(i + 1)
		infoStorages = append(infoStorages, node.NewMemoryNodeInfoStorage())
		stableStorages = append(stableStorages, moon.NewStorage(path.Join(basePath, "/raft", strconv.Itoa(i+1))))
		rpcServers = append(rpcServers, messenger.NewRpcServer(32800+raftID))
		nodeInfos = append(nodeInfos, node.NewSelfInfo(raftID, "127.0.0.1", 32800+raftID))
		alayas = append(alayas, alaya.NewAlaya(nodeInfos[i], infoStorages[i], alaya.NewMemoryMetaStorage(), rpcServers[i]))
	}

	moonConfig := moon.DefaultConfig
	moonConfig.SunAddr = sunAddr
	moonConfig.GroupInfo = node.GroupInfo{
		GroupTerm:       &node.Term{Term: 0},
		LeaderInfo:      nil,
		UpdateTimestamp: timestamp.Now(),
	}

	for i := 0; i < num; i++ {
		if sunAddr != "" {
			moons = append(moons, moon.NewMoon(nodeInfos[i], &moonConfig, rpcServers[i], infoStorages[i],
				stableStorages[i]))
		} else {
			moonConfig.GroupInfo.NodesInfo = nodeInfos
			moons = append(moons, moon.NewMoon(nodeInfos[i], &moonConfig, rpcServers[i], infoStorages[i],
				stableStorages[i]))
		}
	}
	return nodeInfos, moons, alayas, rpcServers, nil
}

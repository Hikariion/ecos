package moon

import (
	"context"
	"ecos/cloud/sun"
	"ecos/edge-node/node"
	"ecos/messenger"
	"go.etcd.io/etcd/raft/v3"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"reflect"
	"testing"
	"time"
)

func TestRaft(t *testing.T) {

	// 先起三个节点
	groupInfo := []*node.NodeInfo{
		node.NewSelfInfo(0x01, "127.0.0.1", 32671),
		node.NewSelfInfo(0x02, "127.0.0.1", 32672),
		node.NewSelfInfo(0x03, "127.0.0.1", 32673),
	}

	rpcServer1 := messenger.NewRpcServer(32671)
	rpcServer2 := messenger.NewRpcServer(32672)
	rpcServer3 := messenger.NewRpcServer(32673)

	node1 := NewMoon(groupInfo[0], "", nil, groupInfo, rpcServer1)
	go rpcServer1.Run()
	go node1.Run()
	defer node1.Stop()
	defer rpcServer1.Stop()
	node2 := NewMoon(groupInfo[1], "", nil, groupInfo, rpcServer2)
	go rpcServer2.Run()
	go node2.Run()
	defer node2.Stop()
	defer rpcServer2.Stop()
	node3 := NewMoon(groupInfo[2], "", nil, groupInfo, rpcServer3)
	go rpcServer3.Run()
	go node3.Run()
	defer node3.Stop()
	defer rpcServer3.Stop()

	nodes := []*Moon{node1, node2, node3}

	// 等待选主
	leader := -1
	for {
		for i := 0; i < 3; i++ {
			if raft.StateLeader == nodes[i].raft.Status().RaftState {
				leader = i
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		if leader >= 0 {
			t.Logf("leader: %v", leader+1)
			break
		}
	}

	// Node4
	node4Info := node.NewSelfInfo(0x04, "127.0.0.1", 32674)
	nodes[leader].AddNodeToGroup(context.TODO(), node4Info)
	rpcServer4 := messenger.NewRpcServer(32674)
	node4 := NewMoon(node4Info, "", nil, groupInfo, rpcServer4)
	go rpcServer4.Run()

	// 集群提交增加节点请求
	_ = node1.raft.ProposeConfChange(node1.ctx, raftpb.ConfChange{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  0x04,
		Context: nil,
	})

	// 等待共识
	time.Sleep(2 * time.Second)

	// 启动 Node4
	go node4.Run()
	defer node4.Stop()

	nodes = append(nodes, node4)

	time.Sleep(2 * time.Second)

	// 判断集群是否达成共识
	info := nodes[0].InfoStorage.ListAllNodeInfo()
	t.Log(info)
	for i := 1; i < 4; i++ {
		anotherInfo := nodes[i].InfoStorage.ListAllNodeInfo()
		if !reflect.DeepEqual(info, anotherInfo) {
			t.Errorf("Node Info Not Equal")
		}
		t.Log(anotherInfo)
	}
	t.Log("Reach agreement success")
	rpcServer1.Stop()
	rpcServer2.Stop()
	rpcServer3.Stop()
	rpcServer4.Stop()
}

func TestMoon_Register(t *testing.T) {
	sunRpc := messenger.NewRpcServer(3260)
	sun.NewSun(sunRpc)
	go sunRpc.Run()
	time.Sleep(1 * time.Second)

	groupInfo := []*node.NodeInfo{
		node.NewSelfInfo(0x01, "127.0.0.1", 32671),
		node.NewSelfInfo(0x02, "127.0.0.1", 32672),
		node.NewSelfInfo(0x03, "127.0.0.1", 32673),
	}

	rpcServers := []*messenger.RpcServer{
		messenger.NewRpcServer(32671),
		messenger.NewRpcServer(32672),
		messenger.NewRpcServer(32673),
	}

	var moons []*Moon

	for i := 0; i < 3; i++ {
		moon := NewMoon(groupInfo[i], "127.0.0.1:3260", nil, nil,
			rpcServers[i])
		moons = append(moons, moon)
		go rpcServers[i].Run()
		defer rpcServers[i].Stop()
		go moon.Run()
		defer moon.Stop()
	}

	time.Sleep(2 * time.Second)

	leader := -1
	for {
		for i := 0; i < 3; i++ {
			if raft.StateLeader == moons[i].raft.Status().RaftState {
				leader = i
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		if leader >= 0 {
			t.Logf("leader: %v", leader+1)
			break
		}
	}

}

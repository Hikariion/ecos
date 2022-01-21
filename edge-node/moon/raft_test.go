package moon

import (
	"go.etcd.io/etcd/raft"
	"testing"
	"time"
)

func TestRaft(t *testing.T) {
	peers := []raft.Peer{{ID: 0x01}, {ID: 0x02}, {ID: 0x03}}
	node1 := newNode(0x01, peers)
	go node1.run()
	node2 := newNode(0x02, peers)
	go node2.run()
	node3 := newNode(0x03, peers)
	go node3.run()

	nodes := []*node{node1, node2, node3}

	for {
		leader := -1
		for i := 0; i < 3; i++ {
			if raft.StateLeader == nodes[i].raft.Status().RaftState {
				leader = i
				break
			}
			time.Sleep(100 * time.Microsecond)
		}
		if leader > 0 {
			t.Logf("leader: %v", leader+1)
			break
		}
	}
}

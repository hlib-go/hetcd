package hetcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"testing"
	"time"
)

func TestConn(t *testing.T) {

	// 修改hosts etcd
	c, err := NewClient()
	if err != nil {
		return
	}

	cluster := clientv3.NewCluster(c)
	member, err := cluster.MemberList(context.TODO())
	for _, item := range member.Members {
		t.Log(item.ID, item.Name, item.PeerURLs, item.ClientURLs)
	}

	var (
		lease   = clientv3.NewLease(c)
		kv      = clientv3.NewKV(c)
		watcher = clientv3.NewWatcher(c)
	)

	lgr, err := lease.Grant(context.TODO(), 10000)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = kv.Put(context.TODO(), "/tlease/ss", "{}", clientv3.WithLease(lgr.ID))
	if err != nil {
		t.Error(err)
		return
	}

	go func() {
		watchChan := watcher.Watch(context.TODO(), "/test/k")
		for true {
			v := <-watchChan
			for _, e := range v.Events {
				t.Log("e.Kv.Value=", string(e.Kv.Value))
			}
			t.Log("v=", v)
		}
	}()

	_, err = kv.Put(context.TODO(), "/kkkk", "vvvv")
	if err != nil {
		t.Log(err)
		return
	}

	time.Sleep(60 * time.Second)
}

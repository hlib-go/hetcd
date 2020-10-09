package hetcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
	"time"
)

/*
# 修改 hosts
# Etcd proxy dev server 端口：2379
etcd 101.133.221.239
*/
func NewClient() (*clientv3.Client, error) {
	log.Info("Etcd NewClient ... ")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"etcd:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	member, err := c.Cluster.MemberList(context.TODO())
	for _, item := range member.Members {
		fmt.Println("Etcd node: ", item.ID, item.Name, item.PeerURLs, item.ClientURLs)
	}
	return c, err
}

// 读取Key值，只返回第一个匹配key的值
func Get(c *clientv3.Client, k string) (value []byte, err error) {
	gr, err := c.Get(context.TODO(), k)
	if err != nil {
		return
	}
	for _, v := range gr.Kvs {
		value = v.Value
		if value == nil {
			continue
		}
		break
	}
	return
}

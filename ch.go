/* ----------------------------------
*  @author suyame 2022-08-29 21:03:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/
package loadbalance

import (
	"loadbalance/ConsistentHash"
	"log"
	"os"
)

// 一致哈希负载均衡算法，保证来自同一IP的请求分配到同一个服务器节点上

type CH struct {
	servers *ConsistentHash.HashCycle
}

func NewCH(servers []Server) (*CH, error) {
	hc := ConsistentHash.NewHashCycle()
	// Init log
	l := log.New(os.Stdout, "LB_CH: ", log.Ldate|log.Ltime)
	hc.SetLogger(l)
	// Use heartbeat
	go ConsistentHash.Heartbeat(hc, &map[ConsistentHash.HashAddrType]bool{})
	// Add new servers
	var err error
	for _, server := range servers {
		err = hc.AddServer(server.(*ConsistentHash.Server))
		if err != nil {
			return nil, err
		}
	}
	return &CH{
		servers: hc,
	}, nil
}

func (ch *CH) Do(request Request) (Server, error) {
	// 根据请求的ip地址分配服务器
	tgtServerip, _, err := ch.servers.FindTgtServer(request.ipstr)
	if err != nil {
		return nil, err
	}
	return ch.servers.GetServer(tgtServerip)
}

/* ----------------------------------
*  @author suyame 2022-08-29 21:32:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/
package loadbalance

import (
	"loadbalance/ConsistentHash"
	"testing"
)

func TestNewCH(t *testing.T) {
	_, err := NewCH(servers2)
	if err != nil {
		t.Error("NewCH error, ", err)
	}
}

func TestCH_Do(t *testing.T) {
	ch, err := NewCH(servers2)
	if err != nil {
		t.Error("NewCH error, ", err)
	}
	// 处理20个请求
	for _, request := range requests {
		server, err := ch.Do(request)
		if err != nil {
			t.Error("CH_Do error", err)
		}
		t.Log("request: ", request.ipstr, "allocated at Server: ", server.(*ConsistentHash.Server).IP())
	}
}

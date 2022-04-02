package natsUtils

import (
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"strings"
	"teacupapi/config"
	glog "teacupapi/logs"
	"teacupapi/utils"
)

var (
	natsConnect         *nats.Conn
	pubStreamingConnect stan.Conn
	subStreamingConnect stan.Conn
	subList             []stan.Subscription
)

func Init() error {
	addressList, err := utils.GetMacAddress()
	if err != nil {
		return err
	}

	if len(addressList) == 0 {
		return errors.New("获取本级mac地址为0")
	}
	macAddress := strings.ReplaceAll(addressList[0], ":", "")
	natsConfig := config.GetNatsConfig()
	pubClientId := natsConfig.PubClientId + macAddress
	subClientId := natsConfig.SubClientId + macAddress
	clusterID := natsConfig.ClusterID
	natsUrl := natsConfig.NatsUrl
	fmt.Println(natsUrl)
	natsc, err := nats.Connect(natsUrl)
	natsConnect = natsc
	if err != nil {
		glog.Errorf("连接nats error:%+v", err)
		return err
	}

	connOptions := []stan.Option{
		stan.NatsConn(natsConnect),
		stan.Pings(10, 100),
	}

	pubStreamingConnect, err = stan.Connect(clusterID, pubClientId,
		append(connOptions, stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			glog.Errorf("连接publish nats streaming lost, reason: %v", reason)
		}))...,
	)
	if err != nil {
		glog.Errorf("连接publish nats streaming err:%+v", err)
		return err
	}
	subStreamingConnect, err = stan.Connect(clusterID, subClientId,
		append(connOptions, stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			glog.Errorf("连接subscribe nats streaming lost, reason: %v", reason)
		}))...,
	)
	if err != nil {
		glog.Errorf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, natsUrl)
		return err
	}
	glog.Infof("Connected to %s clusterID: [%s] clientID: [%s]\n", natsUrl, clusterID, subClientId)
	return nil
}

func Close() {
	for i := range subList {
		err := subList[i].Close()
		if err != nil {
			glog.Errorf("关闭nats streaming err:%+v", err)
		}
	}
	if pubStreamingConnect != nil {
		err := pubStreamingConnect.Close()
		if err != nil {
			glog.Errorf("关闭nats streaming err:%+v", err)
		}
	}
	if subStreamingConnect != nil {
		err := subStreamingConnect.Close()
		if err != nil {
			glog.Errorf("关闭nats streaming err:%+v", err)
		}
	}
	if natsConnect != nil {
		natsConnect.Close()
	}
}

func SubConnected() bool {
	return subStreamingConnect != nil
}

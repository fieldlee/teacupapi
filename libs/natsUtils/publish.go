package natsUtils

import (
	"github.com/nats-io/stan.go"
	"sync"
	"teacupapi/config"
	glog "teacupapi/logs"
	"time"
)

func Publish(subject, message string, ackHandler stan.AckHandler) {
	subject = config.GetBaseConf().Env + "_" + subject
	ackChan := publish(subject, message, ackHandler)
	if ackChan != nil {
		select {
		case <-ackChan:
			break
		case <-time.After(10 * time.Second):
			glog.Errorf("消息发送失败，subject：%s,msg:%s,超过ack时间", subject, message)
		}
	}
}

func publish(subject, message string, ackHandler stan.AckHandler) chan bool {
	if pubStreamingConnect == nil {
		glog.Errorf("消息发送失败，subject：%s,msg:%s,nats streaming 连接不存在", subject, message)
		return nil
	}
	ackChan := make(chan bool)
	var glock sync.Mutex

	var guid string
	glock.Lock()
	if ackHandler == nil {
		ackHandler = func(lguid string, err error) {
			glock.Lock()
			glog.Infof("Received ACK for guid %s", lguid)
			defer glock.Unlock()

			if err != nil {
				glog.Errorf("Error in server ack for guid %s: %v", lguid, err)
			}
			if lguid != guid {
				glog.Errorf("Expected a matching guid in ack callback, got %s vs %s", lguid, guid)
			}
			ackChan <- true
		}
	}
	guid, err := pubStreamingConnect.PublishAsync(subject, []byte(message), ackHandler)

	if err != nil {
		glog.Errorf("消息发送失败，subject：%s,msg:%s", subject, message)
	}
	glock.Unlock()

	if guid == "" {
		glog.Errorf("Expected non-empty guid to be returned.")
	}
	glog.Infof("发送主题： [%s] ，消息： '%s' ，[guid: %s]", subject, message, guid)
	return ackChan
}

package apns

import (
	"errors"
	"github.com/edganiukov/apns"
	"io/ioutil"
	"teacupapi/config"
	glog "teacupapi/logs"
	"teacupapi/service/userService"
	"time"
)

var ApnsClient *apns.Client

func Init() error {
	data, err := ioutil.ReadFile(config.GetApnsConf().PEMFilePath)
	if err != nil {
		glog.Errorf("ApnsClient init error :%+v", err)
		return err
	}
	ApnsClient, err = apns.NewClient(
		apns.WithJWT(data, config.GetApnsConf().KeyId, config.GetApnsConf().Issuer),
		apns.WithBundleID(config.GetApnsConf().BundleId),
		apns.WithMaxIdleConnections(10),
		apns.WithTimeout(5*time.Second),
	)
	if err != nil {
		glog.Errorf("ApnsClient NewClient error :%+v", err)
		return err
	}
	return nil
}

func PushiOSNotification(userId int64, title, body string) error {

	var user userService.UserInfo
	err := user.GetUserById(userId)
	if err != nil {
		glog.Errorf("PushiOSNotification getUser err:%+v", err)
		return err
	}
	if user.UserPhoneType != 1 {
		glog.Errorf("PushiOSNotification getUser %d err:phone type not is ios", userId)
		return errors.New("phone type not is ios")
	}

	if user.UUID == "" {
		glog.Errorf("PushiOSNotification getUser %d err:user uuid is empty", userId)
		return errors.New("user uuid is empty")
	}
	resp, err := ApnsClient.Send(user.UUID, apns.Payload{
		APS: apns.APS{
			Alert: apns.Alert{
				Title: title,
				Body:  body,
			},
		},
	},
		apns.WithExpiration(10),
		apns.WithCollapseID(""),
		apns.WithPriority(5),
	)
	if err != nil {
		glog.Errorf("PushiOSNotification err:%+v", err)
		return err
	}
	if resp.Error != nil {
		glog.Errorf("PushiOSNotification err:%+v", err)
		return err
	}
	return nil
}

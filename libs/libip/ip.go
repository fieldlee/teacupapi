package libip

import (
	"github.com/ipipdotnet/ipdb-go"
	"teacupapi/config"
	glog "teacupapi/logs"
	"teacupapi/models"
)

var (
	ipdbInfo *ipdb.City
)

// 初始化 ip 库
func InitIP() error {
	var err error

	ipdbInfo, err = ipdb.NewCity(config.GetIPConfAddr())
	if err != nil {
		return err
	}

	return nil
}

func GetIPLoc(ip string) *models.IPLoc {
	var data models.IPLoc
	data.Country = "中国"

	if ipdbInfo == nil {
		glog.Errorf("ipdbInfo is nil")
		return &data
	}

	loc, err := ipdbInfo.FindInfo(ip, "CN")
	if err != nil {
		glog.Errorf("ip=%s query err: %v", ip, err)
		return &data
	}

	data.Country = loc.CountryName
	data.Province = loc.RegionName
	data.City = loc.CityName

	return &data
}

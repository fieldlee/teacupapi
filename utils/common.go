package utils

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"net"
	"strconv"
)

// GetEsFromSize 获取 es 的
func GetEsFromSize(pageStr, pageSizeStr string, limit int) (int, int, error) {
	from := 0
	size := 10

	if pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			tmpStr := fmt.Sprintf("pageSize=%s fmt int err: %v", pageSizeStr, err)
			return 0, 0, errors.New(tmpStr)
		}
		if pageSize > limit { // 每次不能超过 limit
			size = limit
		} else if pageSize > 0 {
			size = pageSize
		}
	}
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			tmpStr := fmt.Sprintf("page=%s fmt int err: %v", pageStr, err)
			return 0, 0, errors.New(tmpStr)
		}
		if page > 0 {
			from = (page - 1) * size
		}
	}

	return from, size, nil
}

// 生成转账的订单号
func GetTransferBillNo() string {
	return GetYMDHmSBillNo("T", RealRand(5))
}

// 生成 prefix 前缀 + 年月日时分秒 + 后缀的订单号
func GetYMDHmSBillNo(prefix, suffix string) string {
	return prefix + GetBjNowTime().Format(TimeFormatHMS) + suffix
}

func Cputicks() (t uint64)

func InArray(val string, array []string) (exists bool) {
	exists = false
	for _, v := range array {
		if val == v {
			exists = true
			return
		}
	}
	return
}

func GetID() (int64, error) {
	node, err := snowflake.NewNode(8)
	if err != nil {
		return 0, err
	}
	nid := node.Generate()
	return nid.Int64(), nil
}

//获取当前机器的mac地址
func GetMacAddress() (macAddress []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return macAddress, err
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddress = append(macAddress, macAddr)
	}
	return macAddress, nil
}

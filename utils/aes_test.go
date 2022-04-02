package utils

import (
	"fmt"
	"testing"
)

func TestGetEncodeString(t *testing.T) {
	str := GetEncodeString("D0C3590D050552F2F665EFA91956D3ED", "root:7JJljqC1%LS%%bvYeV6W@tcp(127.0.0.1:20532)/teacup?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println(str)
}

func TestGetRealString(t *testing.T) {
	str := GetRealString("D0C3590D050552F2F665EFA91956D3ED", "NQZIytInKJJQX3jo5e499+/nxFC17SukBL1HU48xo3tDRR4qg4i4xNZ+Zji+2zVvD2IOGriMZooxX4OCPwwcDmnjX0EL1Mrnlicd8ejBfQX8msvHwhrk1iOyqgQ9Ebq1")
	fmt.Println(str)
}

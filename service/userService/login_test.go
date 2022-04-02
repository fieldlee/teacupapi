package userService

import (
	"fmt"
	"teacupapi/config"
	"testing"
)

func TestGetTokenDataByGinContext(t *testing.T) {
	p := "../../config/app.dev.ini"
	err := config.InitConfig(&p)
	if err != nil {
		t.FailNow()
	}

	token := "jYKKzRCBQIG3JEcgEbIoTxmrRSDgB+O4JwHYYb5Bm6m/NWpzXbZliJVWQshqQGOdVZutf6HR8fjqcFSHVdsRKWPyclBe4dvjcC/aTQdn+spFMOco51/niq4WCSbDU3jGyZwITsQ0S+c93lOdNXy5z0YDPB+9w8bc8ycOii7zAPw="
	data, err := ParseToken(token)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("data:%v", data)
}

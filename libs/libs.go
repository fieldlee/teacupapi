package libs

import (
	"errors"
	"fmt"
	"teacupapi/libs/libip"
)

func Initlibs() error {
	var err error

	err = libip.InitIP()
	if err != nil {
		tmpStr := fmt.Sprintf("init ip err %v", err)
		return errors.New(tmpStr)
	}

	return nil
}

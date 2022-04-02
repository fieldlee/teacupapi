package db

import (
	"errors"
	"fmt"

	"teacupapi/db/redisdb"
	"teacupapi/db/sqldb"
)

// 初始化 db
func InitDB() error {
	var err error

	// mysql 的初始化
	err = sqldb.InitMysql()
	if err != nil {
		tmpStr := fmt.Sprintf("init mysql err: %v", err)
		return errors.New(tmpStr)
	}

	// 初始化 redis
	err = redisdb.InitRedis()
	if err != nil {
		tmpStr := fmt.Sprintf("init redis err: %v", err)
		return errors.New(tmpStr)
	}

	return nil
}

// Close close all db
func Close() error {
	return nil
}

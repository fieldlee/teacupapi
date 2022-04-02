package sqldb

import (
	"errors"
	"fmt"
	"teacupapi/utils"
	"time"

	"teacupapi/models"

	"teacupapi/config"
	glog "teacupapi/logs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Sqldb struct {
	tyDB *gorm.DB
	reDb *gorm.DB
}

var sqlDB *Sqldb

// InitMysql init mysql
func InitMysql() error {
	tyDB, err := initSQLDb(
		//config.GetTyDbConf().TSqlAddr,
		utils.GetRealString(config.GetSecretKey(), config.GetTyDbConf().TSqlAddr),
		config.GetTyDbConf().TDbLogEnable,
		config.GetTyDbConf().TDbMaxConnect,
		config.GetTyDbConf().TDbIdleConnect,
		config.GetTyDbConf().TDbMaxLifeTime,
	)
	if err != nil {
		tmpStr := fmt.Sprintf("init mysql db is err: %v", err)
		glog.Errorf(tmpStr)
		return errors.New(tmpStr)
	}
	sqlDB = &Sqldb{
		tyDB: tyDB,
	}
	reDB, err := initSQLDb(
		//config.GetReDbConf().RTSqlAddr,
		utils.GetRealString(config.GetSecretKey(), config.GetReDbConf().RTSqlAddr),
		config.GetReDbConf().RTDbLogEnable,
		config.GetReDbConf().RTDbMaxConnect,
		config.GetReDbConf().RTDbIdleConnect,
		config.GetReDbConf().RTDbMaxLifeTime,
	)
	if err != nil {
		tmpStr := fmt.Sprintf("init replication mysql db is err: %v", err)
		glog.Errorf(tmpStr)
		return errors.New(tmpStr)
	}
	sqlDB.reDb = reDB

	return nil
}

func GetRDb() *gorm.DB {
	return sqlDB.reDb
}
func GetTDb() *gorm.DB {
	return sqlDB.tyDB
}

func getDb(isSlave bool) *gorm.DB {
	if isSlave {
		return sqlDB.reDb
	}
	return sqlDB.tyDB
}

//初始化数据库
func initSQLDb(tsSqlAddr string, enableLog bool, maxConn, idleConn, maxLifeTime int) (*gorm.DB, error) {
	tmpDb, err := gorm.Open("mysql", tsSqlAddr)
	if err != nil {
		tmpStr := fmt.Sprintf("conect=%s the sqldb, err: %v", tsSqlAddr, err)
		glog.Errorf(tmpStr)
		return nil, fmt.Errorf(tmpStr)
	}

	tmpDb.DB().SetMaxOpenConns(maxConn)
	tmpDb.DB().SetMaxIdleConns(idleConn)
	tmpDb.DB().SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)

	if err = tmpDb.DB().Ping(); err != nil {
		tmpStr := fmt.Sprintf("Ping the db=%s err: %v", tsSqlAddr, err)
		glog.Error(tmpStr)
		return nil, fmt.Errorf(tmpStr)
	}

	tmpDb.LogMode(enableLog)
	tmpDb.SingularTable(true)
	env := config.GetEnv()
	if env == "dev" {
		// 自动建表, 测试环境能自动建表, 预发布以上环境不能建表
		tables := []interface{}{}
		tmpDb = tmpDb.AutoMigrate(tables...)
		for _, v := range tables {
			if !tmpDb.HasTable(v) {
				return nil, fmt.Errorf("build table %v failed", v)
			}
		}
	}
	return tmpDb, nil
}

// Close close mysql db
func Close() error {
	_ = sqlDB.tyDB.Close()
	_ = sqlDB.reDb.Close()
	return nil
}

// UpdateDb 更新单条数据
func UpdateTyDb(row interface{}) error {
	err := sqlDB.tyDB.Save(row).Error
	if err != nil {
		return err
	}

	return nil
}

// InsertDb 插入 dbs 数据库
func InsertTyDb(row interface{}) error {
	err := sqlDB.tyDB.Create(row).Error
	if err != nil {
		return err
	}

	return nil
}

// SingleUpdateTyb 更新某一张表的多个列
func SingleUpdateTyb(mods interface{}, upMap map[string]interface{}, query interface{}, args ...interface{}) error {
	err := sqlDB.tyDB.Model(mods).Where(query, args...).Updates(upMap).Error
	if err != nil {
		return err
	}

	return nil
}

// BatchUpdateTyb 批量更新某一张表
func BatchUpdateTyb(tableName string, upMap map[string]interface{}, query interface{}, args ...interface{}) error {
	err := sqlDB.tyDB.Table(tableName).Where(query, args...).Update(upMap).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchOne 取某一张表的多个列的第一条数据
func FetchOne(mods interface{}, tableName string, isSlave bool, cols []string, query interface{}, args ...interface{}) error {
	return getDb(isSlave).Table(tableName).Select(cols).Where(query, args...).First(mods).Error
}

// FetchAll 取某一张表的所有数据
func FetchAll(mods interface{}, tableName string, isSlave bool, cols []string, query interface{}, args ...interface{}) error {
	return getDb(isSlave).Table(tableName).Select(cols).Where(query, args...).Find(mods).Error
}

func FetchAllWithout(mods interface{}, tableName string, isSlave bool, cols []string) error {
	return getDb(isSlave).Table(tableName).Select(cols).Find(mods).Error
}

// 原生 sql
func ExecSQLScan(mods interface{}, sql string, isSlave bool, args ...interface{}) error {
	err := getDb(isSlave).Raw(sql, args...).Scan(mods).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { // 没有该记录
			return models.ErrNotFound
		}
		return err
	}

	return nil
}

// 原生 sql
func ExecSQL(sql string, isSlave bool) error {
	err := getDb(isSlave).Exec(sql).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { // 没有该记录
			return models.ErrNotFound
		}
		return err
	}

	return nil
}

func Count(tableName string, isSlave bool, query interface{}, args ...interface{}) (int, error) {
	var total int
	err := getDb(isSlave).Table(tableName).Where(query, args...).Count(&total).Error
	return total, err
}

func SingleUpdate(tableName string, query interface{}, upMap map[string]interface{}, args ...interface{}) error {
	return sqlDB.tyDB.Table(tableName).Where(query, args...).Updates(upMap).Error
}

func ExecSQLRepScan(mods interface{}, sql string, args ...interface{}) error {
	return sqlDB.reDb.Raw(sql, args...).Scan(mods).Error
}

// SingleUpdateSc 更新某一张表的单列
func SingleUpdateSc(mods, query interface{}, oldValue interface{}, field, fieldExpression string, newValue interface{}) error {
	err := sqlDB.tyDB.Model(mods).Where(query, oldValue).Update(field, gorm.Expr(fieldExpression, newValue)).Error
	if err != nil {
		return err
	}

	return nil
}

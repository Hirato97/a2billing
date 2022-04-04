package mysql

import (

	/* "database/sql"
	"fmt"
	"time"

	IMS "a2billing-go-api/internal/sqldb/mysql"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger" */

	"database/sql"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type IMySqlConnector interface {
	GetDB() *bun.DB
}

type MySqlConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	PingInterval int
	MaxOpenConns int
	MaxIdleConns int
	ReadTimeout  int
	WriteTimeout int
}

type MySqlConnector struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	PingInterval int
	MaxOpenConns int
	MaxIdleConns int
	ReadTimeout  int
	WriteTimeout int
	DB           *bun.DB
	SqlDB        *sql.DB
	MySqlConfig
}

// GetDB implements mysql.IMySqlConnector

// Ping implements mysql.IMySqlConnector

func NewMySqlConnector(config MySqlConfig) IMySqlConnector {
	MySql := &MySqlConnector{
		Host:         config.Host,
		Port:         config.Port,
		User:         config.User,
		Password:     config.Password,
		Database:     config.Database,
		Charset:      config.Charset,
		PingInterval: config.PingInterval,
		MaxOpenConns: config.MaxOpenConns,
		MaxIdleConns: config.MaxIdleConns,
	}
	if err := MySql.Connect(); err != nil {
		log.Fatal("SqlClient", "NewSqlClient", err)
		return nil
	}
	if err := MySql.DB.Ping(); err != nil {
		log.Fatal("SqlClient", "NewSqlClient", err)
		return nil
	}
	// IMS.MySqlConnector = MySql
	return MySql
}

/*
func (conn *MySqlConnector) GetConn() *gorm.DB {
	return conn.DB
}

func (conn *MySqlConnector) Connect() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&autocommit=true", conn.User, conn.Password, conn.Host, conn.Port, conn.Database, conn.Charset)
	var err error
	conn.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dataSourceName,
		DefaultStringSize:         256,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Error("Connect MySQL failed!", err.Error())
		return err
	}
	sqlDB, err := conn.DB.DB()
	sqlDB.SetMaxIdleConns(conn.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conn.MaxOpenConns)
	conn.SqlDB = sqlDB
	return nil
}

// Retry Check Connection
func (conn *MySqlConnector) RetrySql(config MySqlConfig) {
	for {
		time.Sleep(time.Duration(config.PingInterval) * time.Second)
		err := conn.Connect()
		if err != nil {
			log.Error("Connect MySQL failed! ", err.Error())
			continue
		}
		log.Info("Connect MySQL Succcess!")
		break
	}
}

func (conn *MySqlConnector) Ping() error {
	sqlDB, err := conn.DB.DB()
	err = sqlDB.Ping()
	if err == nil {
		log.Info("Connect database mysql successful :", conn.Host)
	} else {
		log.Error("Connect database mysql fail :", err)
	}
	return err
}
*/
func (c *MySqlConnector) Connect() (err error) {

	//username:password@protocol(address)/dbname?param=value
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	sqldb, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("SqlClient", "Connect", err)
		return err
	}
	sqldb.SetMaxIdleConns(c.MaxIdleConns)
	sqldb.SetMaxOpenConns(c.MaxOpenConns)
	db := bun.NewDB(sqldb, mysqldialect.New(), bun.WithDiscardUnknownColumns())
	c.DB = db
	return nil
}

func (c *MySqlConnector) GetDB() *bun.DB {
	return c.DB
}

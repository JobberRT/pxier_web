package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type pxier struct {
	*echo.Echo
	readDB  *gorm.DB
	writeDB *gorm.DB
}

func NewPixer() *pxier {
	p := &pxier{
		Echo: echo.New(),
	}
	p.readDB = newReadDB()
	p.writeDB = newWriteDB()
	p.registerMiddleware()
	p.registerRoute()
	return p
}

// registerMiddleware will register needed middlewares for *echo.Echo
func (p *pxier) registerMiddleware() {
	rateLimit := viper.GetInt("echo.rate_limit")
	if rateLimit == 0 {
		logrus.Warn("rate_limit is 0, set to 3")
		rateLimit = 3
	}
	p.Use(middleware.Recover())
	p.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	p.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(rateLimit))))
	p.Use(logger)
	p.Use(checkRequireProxyParam)
	p.Use(checkReportErrorParam)
}

// registerRoute will register routes for *echo.Echo
func (p *pxier) registerRoute() {
	p.GET("/require", p.GetProxy)
	p.GET("/report", p.ReportErrorProxy)
}

func newReadDB() *gorm.DB {
	logrus.Info("start read mysql")
	url := viper.GetString("database.read_db")
	if len(url) == 0 {
		logrus.Panic("mysql url is empty")
	}
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		logrus.WithError(err).Panic("failed to create db")
	}
	if err := db.AutoMigrate(&proxy{}); err != nil {
		logrus.WithError(err).Panic("failed to migrate model")
	}
	d, _ := db.DB()
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)
	d.SetConnMaxLifetime(time.Hour)
	return db
}

func newWriteDB() *gorm.DB {
	logrus.Info("start write mysql")
	url := viper.GetString("database.write_db")
	if len(url) == 0 {
		logrus.Panic("mysql url is empty")
	}
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		logrus.WithError(err).Panic("failed to create db")
	}
	if err := db.AutoMigrate(&proxy{}); err != nil {
		logrus.WithError(err).Panic("failed to migrate model")
	}
	d, _ := db.DB()
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)
	d.SetConnMaxLifetime(time.Hour)
	return db
}

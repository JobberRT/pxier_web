package core

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

func (p *pxier) GetProxy(c echo.Context) error {
	num := c.Get("num").(int)
	providers := c.Get("providers").([]string)
	eachProviderAverage := num / len(providers)
	if eachProviderAverage == 0 {
		eachProviderAverage = 1
	}
	res := make([]*proxy, 0)
	for _, pvd := range providers {
		temp := make([]*proxy, 0)
		if err := p.readDB.Limit(eachProviderAverage).
			Where("err_times < ?", viper.GetInt("database.max_err")).
			Order("RAND()").
			Find(&temp).Error; err != nil {
			logrus.WithError(err).WithField("provider", pvd).Error("failed to get proxy")
			return c.JSON(http.StatusOK, map[string]any{
				"code": httpFailed,
				"data": err.Error(),
			})
		}
		res = append(res, temp...)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"code": httpSuccess,
		"data": res,
	})
}

func (p *pxier) ReportErrorProxy(c echo.Context) error {
	id := c.QueryParam("id")
	p.writeDB.Model(&proxy{}).Where("id = ?", id).UpdateColumn("err_times", gorm.Expr("err_times + ?", 1))
	return c.JSON(http.StatusOK, map[string]any{
		"code": httpSuccess,
		"data": "success",
	})
}

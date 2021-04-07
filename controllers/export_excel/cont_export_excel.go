package exportexcel

import (
	"context"
	iexportexcel "nuryanto2121/cukur_in_web/interface/export_excel"

	"github.com/labstack/echo/v4"
)

type contExportExcel struct {
	useExportExcel iexportexcel.Usecase
}

func NewContExportExcel(e *echo.Echo, a iexportexcel.Usecase) {
	controller := &contExportExcel{
		useExportExcel: a,
	}

	r := e.Group("/web/export_excel")

	r.GET("/transaksi", controller.GetTrx)
}

func (c *contExportExcel) GetTrx(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	return e.File("")
}

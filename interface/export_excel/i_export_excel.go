package iexportexcel

import (
	"nuryanto2121/cukur_in_web/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Repository interface {
	GetTrxApplist(source int) (result []*models.TrxApp, err error)
}

type Usecase interface {
	GetTrxAppExcel(source int) *excelize.File
}

package contredemteguk

import (
	"context"
	"fmt"
	"net/http"
	iimportredem "nuryanto2121/cukur_in_web/interface/import_redem"
	midd "nuryanto2121/cukur_in_web/middleware"
	"nuryanto2121/cukur_in_web/models"
	tool "nuryanto2121/cukur_in_web/pkg/tools"
	"time"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

// ContFileUpload :
type ContFileUpload struct {
	useRedemTeguk iimportredem.Usecase
}

// NewContFileUpload :
func NewContFileUpload(e *echo.Echo, useRedemTeguk iimportredem.Usecase) {
	cont := &ContFileUpload{
		useRedemTeguk: useRedemTeguk,
	}

	e.Static("/wwwroot", "wwwroot")
	r := e.Group("/api/import_redem")
	// Configure middleware with custom claims
	r.Use(midd.JWT)
	// r.Use(midd.Versioning)
	r.POST("", cont.ImportRedem)

}

// ImportRedem :
// @Summary Import Redem
// @Security ApiKeyAuth
// @Description Upload file excel data redem
// @Tags FileUpload
// @Accept  multipart/form-data
// @Produce json
// @Param import_redem formData file true "Data Redem"
// @Success 200 {object} tool.ResponseModel
// @Router /web-service/api/import_redem [post]
func (u *ContFileUpload) ImportRedem(e echo.Context) (err error) {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		appE      = tool.Res{R: e}
		ListRedem []*models.RedemTeguk

		// size          int64
		// logger        = logging.Logger{}
	)

	form, err := e.MultipartForm()
	if err != nil {
		return err
	}
	images := form.File["import_redem"]

	//directory api
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Print(dir)

	for i, image := range images {
		// Source
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		_ = i
		f, err := excelize.OpenReader(src)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Get value from cell by given worksheet name and axis.
		// cell, err := f.GetCellValue("Sheet1", "B2")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// fmt.Println(cell)
		// Get all the rows in the Sheet1.
		sheet1Name := "Sheet1"
		rows, err := f.GetRows("Sheet1")
		fmt.Print(len(rows))
		// for _, row := range rows {
		// 	for _, colCell := range row {
		// 		fmt.Print(colCell, "\t")
		// 	}
		// 	fmt.Println()
		// }

		for i := 2; i <= len(rows); i++ {
			mdl := models.Model{}
			mdl.UserEdit = "Martin"
			mdl.UserInput = "Martin"

			RedemCD, _ := f.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i))
			ExpireDateString, _ := f.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i))
			ExpireDate, _ := time.Parse("2006-01-02", ExpireDateString)
			row := &models.RedemTeguk{
				RedemCd:     RedemCD,
				ExpiredDate: ExpireDate,
				IsUsed:      false,
				OrderID:     0,
				Model:       mdl,
			}
			ListRedem = append(ListRedem, row)
			// rows = append(rows, row)
		}
		fmt.Print(ListRedem)
		if len(ListRedem) > 0 {
			err = u.useRedemTeguk.CreateRedemTeguk(ctx, ListRedem)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

	}
	return appE.Response(http.StatusOK, "Ok", ListRedem)

	// return c.JSON(http.StatusCreated, models.ResponseImage(http.StatusCreated, imageFormList))
}

package contpatnermaster

import (
	"context"
	"fmt"
	"net/http"
	ipatnermaster "nuryanto2121/cukur_in_web/interface/patner_master"
	midd "nuryanto2121/cukur_in_web/middleware"
	"nuryanto2121/cukur_in_web/models"
	app "nuryanto2121/cukur_in_web/pkg"
	"nuryanto2121/cukur_in_web/pkg/logging"
	tool "nuryanto2121/cukur_in_web/pkg/tools"
	util "nuryanto2121/cukur_in_web/pkg/utils"
	"strconv"

	"github.com/labstack/echo/v4"

	_ "nuryanto2121/cukur_in_web/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type contPatnerMaster struct {
	usePatnerMaster ipatnermaster.Usecase
}

func NewContPatnerMaster(e *echo.Echo, a ipatnermaster.Usecase) {
	controller := &contPatnerMaster{
		usePatnerMaster: a,
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health_check", controller.HealthCheck)

	r := e.Group("/web/patner/patner_master")
	r.Use(midd.JWT)
	// r.Use(midd.Versioning)
	r.GET("/:id", controller.GetDataBy)
	r.GET("", controller.GetList)
	r.POST("", controller.Create)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}

func (u *contPatnerMaster) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}

// GetDataByID :
// @Summary GetById
// @Security ApiKeyAuth
// @Tags PatnerMaster
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /web/patner/patner_master/{id} [get]
func (u *contPatnerMaster) GetDataBy(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = tool.Res{R: e} // wajib
		id     = e.Param("id")  //kalo bukan int => 0
	)
	ID, err := strconv.Atoi(id)
	logger.Info(ID)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	data, err := u.usePatnerMaster.GetDataBy(ctx, claims, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", data)
}

// GetList :
// @Summary GetList PatnerMaster
// @Security ApiKeyAuth
// @Tags PatnerMaster
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /web/patner/patner_master [get]
func (u *contPatnerMaster) GetList(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{}
		appE         = tool.Res{R: e}     // wajib
		paramquery   = models.ParamList{} // ini untuk list
		responseList = models.ResponseModelList{}
		err          error
	)

	httpCode, errMsg := app.BindAndValid(e, &paramquery)
	// logger.Info(util.Stringify(paramquery))
	if httpCode != 200 {
		return appE.ResponseErrorList(http.StatusBadRequest, errMsg, responseList)
	}
	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseErrorList(http.StatusBadRequest, fmt.Sprintf("%v", err), responseList)
	}

	responseList, err = u.usePatnerMaster.GetList(ctx, claims, paramquery)
	if err != nil {
		return appE.ResponseErrorList(tool.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	return appE.ResponseList(http.StatusOK, "", responseList)
}

// CreatePatnerMaster :
// @Summary Add PatnerMaster
// @Security ApiKeyAuth
// @Tags PatnerMaster
// @Produce json
// @Param req body models.AddPatnerMaster true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /web/patner/patner_master [post]
func (u *contPatnerMaster) Create(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		form   models.AddPatnerMaster
	)

	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	err = u.usePatnerMaster.Create(ctx, claims, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", nil)
}

// UpdatePatnerMaster :
// @Summary Rubah PatnerMaster
// @Security ApiKeyAuth
// @Tags PatnerMaster
// @Produce json
// @Param id path string true "ID"
// @Param req body models.AddPatnerMaster true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /web/patner/patner_master/{id} [put]
func (u *contPatnerMaster) Update(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		err    error

		id   = e.Param("id") //kalo bukan int => 0
		form = models.AddPatnerMaster{}
	)

	SchoolID, err := strconv.Atoi(id)
	logger.Info(SchoolID)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	// form.UpdatedBy = claims.PatnerMasterName
	err = u.usePatnerMaster.Update(ctx, claims, SchoolID, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusCreated, "Ok", nil)
}

// DeletePatnerMaster :
// @Summary Delete PatnerMaster
// @Security ApiKeyAuth
// @Tags PatnerMaster
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /web/patner/patner_master/{id} [delete]
func (u *contPatnerMaster) Delete(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = tool.Res{R: e} // wajib
		id     = e.Param("id")
	)
	ID, err := strconv.Atoi(id)
	logger.Info(ID)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	err = u.usePatnerMaster.Delete(ctx, claims, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", nil)
}

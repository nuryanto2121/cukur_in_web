package routes

import (

	// sqlxposgresdb "nuryanto2121/cukur_in_web/pkg/postgresqlxdb"
	postgresgorm "nuryanto2121/cukur_in_web/pkg/postgregorm"

	_contPatnerMaster "nuryanto2121/cukur_in_web/controllers/patner_master"
	_repoPatnerMaster "nuryanto2121/cukur_in_web/repository/patner_master"
	_usePatnerMaster "nuryanto2121/cukur_in_web/usecase/patner_master"

	_contImportRedem "nuryanto2121/cukur_in_web/controllers/import_redem"
	_useImportRedem "nuryanto2121/cukur_in_web/usecase/import_redem"

	"nuryanto2121/cukur_in_web/pkg/setting"
	"time"

	"github.com/labstack/echo/v4"
)

//Echo :
type EchoRoutes struct {
	E *echo.Echo
}

func (e *EchoRoutes) InitialRouter() {
	timeoutContext := time.Duration(setting.FileConfigSetting.Server.ReadTimeout) * time.Second

	repoPatnerMaster := _repoPatnerMaster.NewRepoPatnerMaster(postgresgorm.Conn)
	usePatnerMaster := _usePatnerMaster.NewUsePatnerMaster(repoPatnerMaster, timeoutContext)
	_contPatnerMaster.NewContPatnerMaster(e.E, usePatnerMaster)

	useImportRedem := _useImportRedem.NewImportRedem(postgresgorm.Conn, timeoutContext)
	_contImportRedem.NewContFileUpload(e.E, useImportRedem)

}

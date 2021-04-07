package main

import (
	"fmt"
	"log"
	"nuryanto2121/cukur_in_web/pkg/logging"
	"nuryanto2121/cukur_in_web/pkg/monggodb"
	"nuryanto2121/cukur_in_web/routes"

	postgresgorm "nuryanto2121/cukur_in_web/pkg/postgregorm"
	// postgresgorm "job_cukur-in/pkg/postgregorm"
	"nuryanto2121/cukur_in_web/pkg/setting"
	"nuryanto2121/cukur_in_web/redisdb"
	useredem "nuryanto2121/cukur_in_web/usecase/use_redem"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	setting.Setup()
	// postgresdb.Setup()
	postgresgorm.Setup()
	redisdb.Setup()
	monggodb.Setup()
	logging.Setup()

}

// @title Capster Cukur-in
// @version 1.0
// @description REST API for Capter Cukur-in

// @contact.name Nuryanto
// @contact.url https://www.linkedin.com/in/nuryanto-1b2721156/
// @contact.email nuryantofattih@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// e.Use(midd.MiddlewareOne)
	// e.Use(jwt.JWT(e))
	// e.Debug = false

	R := routes.EchoRoutes{E: e}

	R.InitialRouter()

	sPort := fmt.Sprintf(":%d", setting.FileConfigSetting.Server.HTTPPort)
	// maxHeaderBytes := 1 << 60
	// s := &http.Server{
	// 	Addr:           sPort,
	// 	ReadTimeout:    1000,  //setting.FileConfigSetting.Server.ReadTimeout,
	// 	WriteTimeout:   10000, //setting.FileConfigSetting.Server.WriteTimeout,
	// 	MaxHeaderBytes: maxHeaderBytes,
	// }
	// log.Fatal(e.StartServer(s))
	//s.ListenAndServe()
	go runCrond()
	log.Fatal(e.Start(sPort))

	//log.Fatal(e.StartServer(s))

}

func runCrond() {
	//sc := setting.FileConfigSetting.Server.Seconds
	gocron.Every(1).Minutes().Do(RedemTegukin)

	<-gocron.Start()
}

func RedemTegukin() {
	useredem.ProsesRedem()
}

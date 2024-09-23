package main

import (
	"context"
	"fmt"
	"log"
	"nuryanto2121/cukur_in_web/pkg/logging"
	"nuryanto2121/cukur_in_web/routes"
	"time"

	postgresgorm "nuryanto2121/cukur_in_web/pkg/postgregorm"
	// postgresgorm "job_cukur-in/pkg/postgregorm"
	"nuryanto2121/cukur_in_web/pkg/setting"
	"nuryanto2121/cukur_in_web/redisdb"

	_repoNotification "nuryanto2121/cukur_in_web/repository/notification"
	_repoOrder "nuryanto2121/cukur_in_web/repository/order"
	_useNotification "nuryanto2121/cukur_in_web/usecase/notification"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	setting.Setup()
	// postgresdb.Setup()
	postgresgorm.Setup()

	// monggodb.Setup()
	logging.Setup()

}

// @title Web Cukur-in
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

	redisHandler := redisdb.New()
	R := routes.EchoRoutes{E: e}

	R.InitialRouter(redisHandler)

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
	go runCrond(redisHandler)
	log.Fatal(e.Start(sPort))

	//log.Fatal(e.StartServer(s))

}

func runCrond(redis *redisdb.RedisHandler) {
	var logger = logging.Logger{}
	repoOrder := _repoOrder.NewRepoOrder(postgresgorm.Conn)
	repoNotif := _repoNotification.NewRepoNotification(postgresgorm.Conn)
	useNotif := _useNotification.NewUseNotification(repoNotif, repoOrder, redis)
	//sc := setting.FileConfigSetting.Server.Seconds
	// gocron.Every(1).Minutes().Do(RedemTegukin(redis))
	// // useNotif.NotifArriveOnTimeUser(context.Background())
	s := gocron.NewScheduler(time.UTC)

	// Schedule the runJob function to run every minute
	_, err := s.Every(1).Minutes().Do(func() {
		//check order_date yang mendekati 30 dan 15 menit sebelum
		useNotif.NotifArriveOnTimeUser(context.Background())
	})
	if err != nil {
		// log.Fatal(err)
		logger.Error("[Cron-Job][SendNotifOnTimeUser]", err)
	}

	// Start the scheduler asynchronously
	s.StartAsync()
}

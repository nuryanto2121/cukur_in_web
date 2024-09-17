package usenotification

import (
	"context"
	"fmt"
	inotification "nuryanto2121/cukur_in_web/interface/notification"
	iorder "nuryanto2121/cukur_in_web/interface/order"
	"nuryanto2121/cukur_in_web/models"
	"nuryanto2121/cukur_in_web/pkg/logging"
	util "nuryanto2121/cukur_in_web/pkg/utils"
	"nuryanto2121/cukur_in_web/redisdb"
	"strconv"
	"time"

	fcmgetway "nuryanto2121/cukur_in_web/pkg/fcm"

	"github.com/mitchellh/mapstructure"
)

type useNotification struct {
	repoNotification inotification.Repository
	repoOrder        iorder.Repository
	redis            *redisdb.RedisHandler
}

func NewUseNotification(repoNotif inotification.Repository, repoOrder iorder.Repository, redis *redisdb.RedisHandler) inotification.Usecase {
	return &useNotification{
		repoNotification: repoNotif,
		repoOrder:        repoOrder,
		redis:            redis,
	}
}

// Create implements inotification.Usecase.
func (u *useNotification) Create(ctx context.Context, Token string, data *models.AddNotification) (err error) {

	var (
		mNotification = models.Notification{}
		queryParam    = models.ParamList{}
		TokenFCM      []string
		logger        = logging.Logger{}
	)

	// mapping to struct model saRole
	err = mapstructure.Decode(data, &mNotification.AddNotification)
	if err != nil {
		logger.Error("failed mapstructure", err)
		return err
	}

	mNotification.UserEdit = "system.cron"
	mNotification.UserInput = "system.cron"

	err = u.repoNotification.Create(ctx, &mNotification)
	if err != nil {
		logger.Error("failed create notif", err)
		return err
	}
	// send notif to user

	TokenFCM = []string{Token}
	queryParam.InitSearch = fmt.Sprintf("user_id = %d and notification_status = 'N' ", data.UserId)
	cntNotif, err := u.repoNotification.Count(ctx, queryParam)
	if err != nil {
		return err
	}

	fcm := &fcmgetway.SendFCM{
		Title:       mNotification.Title,
		Body:        mNotification.Descs,
		JumlahNotif: cntNotif,
		DeviceToken: TokenFCM,
	}

	// go fcm.SendPushNotification()
	err = fcm.SendPushNotification()
	if err != nil {
		logger.Error(fmt.Sprintf("failed send notification :[TokenFCM :%s]", TokenFCM), err)
		return err
	}
	//end send notif

	return nil
}

// GetCountNotif implements inotification.Usecase.
func (u *useNotification) GetCountNotif(ctx context.Context) (result interface{}, err error) {
	panic("unimplemented")
}

// NotifArriveOnTimeUser implements inotification.Usecase.
func (u *useNotification) NotifArriveOnTimeUser(ctx context.Context) error {
	//get data order notification status New and filter Close time 30 and 15 minute
	var (
		logger = logging.Logger{}
	)

	orders, err := u.repoOrder.GetDataOrderStatusArriveOnTime(ctx)
	if err != nil {
		return err
	}

	for _, val := range orders {
		key := fmt.Sprintf("notif:[arriveOnTimeUser:%d][order:%d]", val.UserID, val.OrderID)
		mutex, err := u.redis.CreateAndLockMutex(ctx, key, time.Minute*5)
		if err != nil {
			logger.Error("[Cron-Job][NotifArriveOnTimeUser]", err)
		}
		//process send notif
		userFCM := fmt.Sprintf("%v", u.redis.GetSession(ctx, strconv.Itoa(val.UserID)+"_fcm"))
		err = u.Create(ctx, userFCM, &models.AddNotification{
			Title:              "30 Menit Lagi: Waktunya Gaya di Barber!",
			Descs:              "Persiapkan dirimu, waktu booking di barber sebentar lagi! Kamu akan tampil keren dalam 30 menit. Segera datang dan siap-siap jadi pusat perhatian!",
			NotificationStatus: "N",
			NotificationType:   "I",
			LinkId:             val.OrderID,
			NotificationDate:   util.GetTimeNow(),
			UserId:             val.UserID,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("[Cron-Job][NotifArriveOnTimeUser][failedSendNotif][user:%d][order:%d]", val.UserID, val.OrderID), err)
		}

		mutex.Unlock()

	}
	// if len(errList) > 0 {

	// }
	return nil
}

// PushNotif implements inotification.Usecase.
func (u *useNotification) PushNotif(ctx context.Context, ID int) error {
	panic("unimplemented")
}

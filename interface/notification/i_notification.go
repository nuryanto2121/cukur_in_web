package inotification

import (
	"context"
	"nuryanto2121/cukur_in_web/models"
)

type Repository interface {
	GetDataBy(ctx context.Context, ID int) (result *models.Notification, err error)
	// GetList(ctx context.Context,UserID int, queryparam models.ParamListGeo) (result []*models.NotificationList, err error)
	Create(ctx context.Context, data *models.Notification) (err error)
	Update(ctx context.Context, ID int, data map[string]interface{}) (err error)
	Delete(ctx context.Context, ID int) (err error)
	Count(ctx context.Context, queryparam models.ParamList) (result int64, err error)
}

type Usecase interface {
	// GetDataBy(ctx context.Context, Claims util.Claims, ID int) (result *models.Notification, err error)
	// GetList(ctx context.Context, Claims util.Claims, queryparam models.ParamListGeo) (result models.ResponseModelList, err error)
	Create(ctx context.Context, TokenFCM string, data *models.AddNotification) (err error)
	// Update(ctx context.Context, Claims util.Claims, ID int, data *models.StatusNotification) (err error)
	// Delete(ctx context.Context, Claims util.Claims, ID int) (err error)
	GetCountNotif(ctx context.Context) (result interface{}, err error)
	PushNotif(ctx context.Context, ID int) error
	NotifArriveOnTimeUser(ctx context.Context) error
}

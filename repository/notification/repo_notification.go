package reponotification

import (
	"context"
	inotification "nuryanto2121/cukur_in_web/interface/notification"
	"nuryanto2121/cukur_in_web/models"
	"nuryanto2121/cukur_in_web/pkg/logging"
	"nuryanto2121/cukur_in_web/pkg/setting"

	"gorm.io/gorm"
)

type repoNotification struct {
	Conn *gorm.DB
}

func NewRepoNotification(Conn *gorm.DB) inotification.Repository {
	return &repoNotification{Conn}
}

func (db *repoNotification) GetDataBy(ctx context.Context, ID int) (result *models.Notification, err error) {
	var (
		logger        = logging.Logger{}
		mNotification = &models.Notification{}
	)
	query := db.Conn.WithContext(ctx).Where("notification_id = ? ", ID).Find(mNotification)
	err = query.Error
	if err != nil {
		logger.Error("[GetDataBy] data not found ", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return mNotification, nil
}

func (db *repoNotification) GetList(ctx context.Context, queryparam models.ParamList) (result []*models.Notification, err error) {

	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = queryparam.SortField
		query    *gorm.DB
	)
	// pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}
	//end pagination

	// Order
	if queryparam.SortField != "" {
		orderBy = queryparam.SortField
	}
	//end Order by

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and (lower(notification_status) LIKE ?)"
		} else {
			sWhere += "(lower(notification_status) LIKE ?)"
		}
		query = db.Conn.WithContext(ctx).Where(sWhere, queryparam.Search).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	} else {
		query = db.Conn.WithContext(ctx).Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}

	//cath to log query string
	err = query.Error

	if err != nil {
		logger.Error("[GetList] data not found ", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}

func (db *repoNotification) Create(ctx context.Context, data *models.Notification) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.WithContext(ctx).Create(data)
	//cath to log query string
	err = query.Error
	if err != nil {
		logger.Error("[Create] failed ", err.Error())
		return err
	}
	return nil
}
func (db *repoNotification) Update(ctx context.Context, ID int, data map[string]interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.WithContext(ctx).Model(models.Notification{}).Where("link_id = ?", ID).Updates(data)
	//cath to log query string
	err = query.Error
	if err != nil {
		logger.Error("[Update] failed ", err.Error())
		return err
	}
	return nil
}

func (db *repoNotification) Delete(ctx context.Context, ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.WithContext(ctx).Where("notification_id = ?", ID).Delete(&models.Notification{})
	//cath to log query string
	err = query.Error
	if err != nil {
		logger.Error("[Delete] failed ", err.Error())
		return err
	}
	return nil
}

func (db *repoNotification) Count(ctx context.Context, queryparam models.ParamList) (result int64, err error) {
	var (
		sWhere = ""
		logger = logging.Logger{}
		query  *gorm.DB
	)
	result = 0

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and (lower(notification_status) LIKE ? )" //+ queryparam.Search
		} else {
			sWhere += "(lower(notification_status) LIKE ? )" //queryparam.Search
		}
		query = db.Conn.WithContext(ctx).Model(&models.Notification{}).Where(sWhere, queryparam.Search).Count(&result)
	} else {
		query = db.Conn.WithContext(ctx).Model(&models.Notification{}).Where(sWhere).Count(&result)
	}
	// end where

	//cath to log query string
	err = query.Error
	if err != nil {
		logger.Error("[count] failed ", err.Error())
		return 0, err
	}

	return result, nil
}

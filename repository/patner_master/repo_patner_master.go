package repopatnermaster

import (
	"fmt"
	ipatnermaster "nuryanto2121/cukur_in_web/interface/patner_master"
	"nuryanto2121/cukur_in_web/models"
	"nuryanto2121/cukur_in_web/pkg/logging"
	"nuryanto2121/cukur_in_web/pkg/setting"

	"gorm.io/gorm"
)

type RepoPatnerMaster struct {
	Conn *gorm.DB
}

func NewRepoPatnerMaster(Conn *gorm.DB) ipatnermaster.Repository {
	return &RepoPatnerMaster{Conn}
}

func (db *RepoPatnerMaster) GetDataBy(ID int) (result *models.PatnerMaster, err error) {
	var (
		logger        = logging.Logger{}
		mPatnerMaster = &models.PatnerMaster{}
	)
	query := db.Conn.Where("patner_master_id = ? ", ID).Find(mPatnerMaster)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return mPatnerMaster, nil
}

func (db *RepoPatnerMaster) GetList(queryparam models.ParamList) (result []*models.PatnerMaster, err error) {

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
			sWhere += " and (lower(name) LIKE ?)"
		} else {
			sWhere += "(lower(name) LIKE ?)"
		}
		query = db.Conn.Where(sWhere, queryparam.Search).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	} else {
		query = db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}

	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}
func (db *RepoPatnerMaster) GetListN(lat float64, long float64) (result []*models.PatnerMaster, err error) {

	sSql := `
	select * 
	from patner_master pm 
	order by fn_distance(pm.latitude,pm.longitude,?,?)
	limit 10
	`
	query := db.Conn.Raw(sSql, lat, long).Find(&result)
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (db *RepoPatnerMaster) Create(data *models.PatnerMaster) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *RepoPatnerMaster) Update(ID int, data map[string]interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Model(models.PatnerMaster{}).Where("patnermaster_id = ?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *RepoPatnerMaster) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Where("patner_master_id = ?", ID).Delete(&models.PatnerMaster{})
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *RepoPatnerMaster) Count(queryparam models.ParamList) (result int, err error) {
	var (
		sWhere  = ""
		logger  = logging.Logger{}
		query   *gorm.DB
		_result int64 = 0
	)
	result = 0

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and (lower(name) LIKE ? )" //+ queryparam.Search
		} else {
			sWhere += "(lower(name) LIKE ? )" //queryparam.Search
		}
		query = db.Conn.Model(&models.PatnerMaster{}).Where(sWhere, queryparam.Search).Count(&_result)
	} else {
		query = db.Conn.Model(&models.PatnerMaster{}).Where(sWhere).Count(&_result)
	}
	// end where

	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}

	result = int(_result)
	return result, nil
}

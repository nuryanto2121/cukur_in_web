package reporedem

import (
	"fmt"
	"log"
	"nuryanto2121/cukur_in_web/models"
	"nuryanto2121/cukur_in_web/pkg/logging"

	"gorm.io/gorm"
)

type RepoRedem struct {
	Conn *gorm.DB
}

func (db *RepoRedem) RedemCode() string {
	var (
		result string = ""
	)

	query := db.Conn.Model(&models.RedemTeguk{}).Where(`order_id is null`).Select(`MAX(redem_cd)`).Row()
	log.Printf(fmt.Sprintf("%v", query))
	query.Scan(&result)
	// err := query.Error
	// if err != nil {
	// 	return ""
	// }

	return (result)
}

func (db *RepoRedem) CountRedem() int {
	var (
		result int64 = 0
	)

	query := db.Conn.Model(&models.RedemTeguk{}).Where(`order_id = 0`).Count(&result)
	log.Printf(fmt.Sprintf("%v", query))
	err := query.Error
	if err != nil {
		return 0
	}

	return int(result)
}

func (db *RepoRedem) Update(RedemCd string, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Model(models.RedemTeguk{}).Where("redem_cd = ?", RedemCd).Updates(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *RepoRedem) Insert(data []*models.RedemTeguk) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(&data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *RepoRedem) Delete() error {

	var (
		logger = logging.Logger{}
		err    error
	)

	query := db.Conn.Exec(`DELETE FROM redem_teguk`)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

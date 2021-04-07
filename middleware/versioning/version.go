package version

import "gorm.io/gorm"

type SsVersion struct {
	VersionID int    `json:"version_id" gorm:"PRIMARY_KEY"`
	OS        string `json:"os" gorm:"type:varchar(20)"`
	Version   int    `json:"version" gorm:"type:integer"`
}

func (V *SsVersion) GetVersion(Conn *gorm.DB) (result SsVersion, err error) {
	err = Conn.Where("os = ? AND apps = 'barber' ", V.OS).First(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

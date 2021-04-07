package sendredem

import (
	"fmt"
	templateemail "nuryanto2121/cukur_in_web/pkg/email"
	postgresgorm "nuryanto2121/cukur_in_web/pkg/postgregorm"
	util "nuryanto2121/cukur_in_web/pkg/utils"
	repopatnermaster "nuryanto2121/cukur_in_web/repository/patner_master"
	"strings"
)

type SendRedem struct {
	Email     string  `json:"email"`
	Name      string  `json:"name"`
	RedemCd   string  `json:"redem_cd"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (R *SendRedem) SendEmail() error {
	subjectEmail := "Selamat kamu mendapatkan voucher gratis teguk"
	// fmt.Printf(subjectEmail)
	err := util.SendEmail(R.Email, subjectEmail, getVerifyBody(R))
	if err != nil {
		return err
	}
	return nil
}

func getVerifyBody(R *SendRedem) string {

	repoPatner := &repopatnermaster.RepoPatnerMaster{
		Conn: postgresgorm.Conn,
	}
	//  <li style="text-align: left;">Teguk A : <a href="https://www.google.com/maps/dir//-6.1784428,107.0077443" rel="nofollow">Maps</a></li>
	var stOutlet = ""
	dataPatner, _ := repoPatner.GetListN(R.Latitude, R.Longitude)
	for _, data := range dataPatner {
		fmt.Print(data)
		stOutlet += fmt.Sprintf(`<li style="text-align: left;"> %s : <a href="https://www.google.com/maps/dir//%g,%g" rel="nofollow">%s</a></li>`, data.Name, data.Latitude, data.Longitude, data.Address)
		fmt.Print(stOutlet)
	}

	redemHTML := templateemail.EmailRedem

	redemHTML = strings.ReplaceAll(redemHTML, `{Name}`, R.Name)
	redemHTML = strings.ReplaceAll(redemHTML, `{RedemCd}`, R.RedemCd)
	redemHTML = strings.ReplaceAll(redemHTML, `{Outlet}`, stOutlet)
	return redemHTML
}

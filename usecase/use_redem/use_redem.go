package useredem

import (
	"fmt"
	"log"
	postgresgorm "nuryanto2121/cukur_in_web/pkg/postgregorm"
	reporedem "nuryanto2121/cukur_in_web/repository/redem"
	repoorder "nuryanto2121/cukur_in_web/repository/repo_order"
	sendredem "nuryanto2121/cukur_in_web/usecase/send_redem"
)

func ProsesRedem() {
	rpRedem := &reporedem.RepoRedem{
		Conn: postgresgorm.Conn,
	}
	// _ = rpRedem

	rpOrder := &repoorder.RepoOrder{
		Conn: postgresgorm.Conn,
	}
	_ = rpOrder

	CntRedem := rpRedem.CountRedem()
	if CntRedem > 0 { // check masih ada redem apa gk
		OrderList, err := rpOrder.DataOrder()
		if err != nil {
			fmt.Printf("%v", err)
			log.Fatalln(err)
		}

		for _, data := range OrderList { // looping transaksi yg telah finish
			DataRedemCd, err := rpRedem.FirstGetData()
			if err != nil {
				fmt.Printf("%v", err)
				// log.Fatalln(err)
			}
			// _ = RedemCd
			if DataRedemCd == nil || DataRedemCd.RedemCd == "" {
				continue
			}

			// send emaail here
			MailService := &sendredem.SendRedem{
				Email:       data.Email,
				Name:        data.Name,
				RedemCd:     DataRedemCd.RedemCd,
				Latitude:    data.Latitude,
				Longitude:   data.Longitude,
				ExpiredDate: DataRedemCd.ExpiredDate,
			}

			// go MailService.SendEmail()
			err = MailService.SendEmail()
			if err != nil {
				fmt.Printf("%v", err)
				log.Fatalln(err)
			}

			UpdateRedem := map[string]interface{}{
				"order_id":  data.OrderID,
				"is_used":   true,
				"barber_id": data.BarberID,
				// "user_edit": data.Email,
			}
			err = rpRedem.Update(DataRedemCd.RedemCd, UpdateRedem)
			if err != nil {
				fmt.Printf("%v", err)
				log.Fatalln(err)
			}
			// _ = UpdateRedem

		}
	}
}

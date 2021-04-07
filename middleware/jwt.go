package midd

import (
	"fmt"
	"net/http"
	version "nuryanto2121/cukur_in_web/middleware/versioning"
	postgresgorm "nuryanto2121/cukur_in_web/pkg/postgregorm"
	tool "nuryanto2121/cukur_in_web/pkg/tools"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			code  = http.StatusOK
			msg   = ""
			data  interface{}
			token = e.Request().Header.Get("Authorization")
		)
		fmt.Println(e.Request().Header)
		data = map[string]string{
			"token": token,
		}

		if token == "" {
			code = http.StatusNetworkAuthenticationRequired
			msg = "Auth Token Required"
		} else {
			if strings.ToLower(token) != "martin" {
				code = http.StatusUnauthorized
				msg = "Token Failed"
			}
			// existToken := redisdb.GetSession(token)
			// if existToken == "" {
			// 	code = http.StatusUnauthorized
			// 	msg = "Token Failed"
			// }
			// claims, err := util.ParseToken(token)
			// if err != nil {
			// 	code = http.StatusUnauthorized
			// 	switch err.(*jwt.ValidationError).Errors {
			// 	case jwt.ValidationErrorExpired:
			// 		msg = "Token Expired"
			// 	default:
			// 		msg = "Token Failed"
			// 	}
			// } else {
			// 	var issuer = setting.FileConfigSetting.App.Issuer
			// 	valid := claims.VerifyIssuer(issuer, true)
			// 	if !valid {
			// 		code = http.StatusUnauthorized
			// 		msg = "Issuer is not valid"
			// 	}
			// 	e.Set("claims", claims)
			// }
		}

		if code != http.StatusOK {
			resp := tool.ResponseModel{
				Msg:  msg,
				Data: data,
			}
			return e.JSON(code, resp)

			// return nil
		}
		return next(e)
	}
}
func Versioning(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			OS    = e.Request().Header.Get("OS")
			Versi = e.Request().Header.Get("Version")
		)
		Version, err := strconv.Atoi(Versi)

		if Version == 0 {
			resp := tool.ResponseModel{
				Msg:  "Please Set Header Version",
				Data: nil,
			}
			return e.JSON(http.StatusBadRequest, resp)
		}

		verService := &version.SsVersion{
			OS:      OS,
			Version: Version,
		}
		dataVersion, err := verService.GetVersion(postgresgorm.Conn) //sqlxposgresdb.GetVersion(OS)
		if err != nil {
			resp := tool.ResponseModel{
				Msg:  fmt.Sprintf("Versioning : %v", err),
				Data: nil,
			}
			return e.JSON(http.StatusBadRequest, resp)
		}

		if dataVersion.Version > Version {
			resp := tool.ResponseModel{
				Msg:  "Please Update Your Apps",
				Data: dataVersion.Version,
			}
			return e.JSON(http.StatusHTTPVersionNotSupported, resp)
		}
		if dataVersion.Version <= Version {
			// resp := tool.ResponseModel{
			// 	Msg:  "Version Not Support",
			// 	Data: dataVersion.Version,
			// }
			// return e.JSON(http.StatusHTTPVersionNotSupported, resp)
			return next(e)
		}

		//end

		return next(e)
	}
}

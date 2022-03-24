package auditoria

import (
  "github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
  "github.com/udistrital/utils_oas/request"
)

// InterceptMidRequest ...
//
// Deprecated: Movida-a/Traer-de utils_oas/auditoria
func InterceptMidRequest(ctx *context.Context) {
	end_point := ctx.Request.URL.String()
	if end_point != "/" {
		defer func () {
			//Catch
			if r := recover(); r != nil {
			}
		}()
		// try
		request.SetHeader(ctx.Request.Header["Authorization"][0])
	}

}

// InitInterceptor ...
//
// Deprecated: Movida-a/Traer-de utils_oas/auditoria
func InitInterceptor() {
	 beego.InsertFilter("*", beego.BeforeExec, InterceptMidRequest, false)
}

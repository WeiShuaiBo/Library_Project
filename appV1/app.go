package appV1

import (
	"library/appV1/model"
	"library/appV1/router"
	"library/appV1/tools"
)

func Start() {

	defer func() {
		model.Close()
	}()

	model.New()
	tools.NewToken("")
	//SendEmail.SendEmail("406624873@qq.com", "thouqwhtgioue")
	r := router.New()
	_ = r.Run(":8080")

}

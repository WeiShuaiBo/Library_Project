package appV0

import (
	"library/appV0/model"
	"library/appV0/router"
	"library/appV0/tools"
)

func Start() {
	defer func() {
		model.Close()
	}()

	model.New()
	tools.NewToken("")

	r := router.New()
	_ = r.Run(":8080")
}

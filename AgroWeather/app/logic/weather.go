package logic

import (
	"AgroWeather/app/goquery"
	"AgroWeather/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetW godoc
//
//	@Summary		查询
//	@Description	根据需求查询
//	@Tags			weather
//	@Accept			multipart/form-data
//	@Produce		json
//	@response		200,400,500	{object}	tools.HttpCode
//	@Router			/weather/do [post]
func GetW(c *gin.Context) {

	preprocess := goquery.Preprocess()

	fmt.Println(preprocess)

	//json
	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
		Data: preprocess,
	})

	return
}

package lib

import (
	"github.com/gin-gonic/gin"
)

//BuildResp will build response model
func BuildResp(c *gin.Context, data interface{}, err error) {
	httpCode := getStatusCode(err)
	c.JSON(httpCode, gin.H{
		"data":  data,
		"error": err,
	})
}

func getStatusCode(err error) (code int) {
	//TODO: compare error to provide accurate status code
	if err != nil {
		code = 500
	} else {
		code = 200
	}
	return
}

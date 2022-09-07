package auth

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/v1"
	"gohub/app/requests"
	"gohub/app/utils"
	"net/http"
)

type SignupController struct {
	v1.BaseController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": utils.IsPhoneExist(request.Phone),
	})
}

package auth

import (
	"github.com/gin-gonic/gin"
	"gohub/app/requests"
	"gohub/app/utils"
	"net/http"
)

type SignupController struct{}

func (s *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": utils.IsDataExist("phone", request.Phone),
	})
}

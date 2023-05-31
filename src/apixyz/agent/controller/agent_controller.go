package controller

import (
	"apixyz/src/apixyz/domain"
	"apixyz/src/apixyz/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerService domain.CustomerService
}

func CreateCustomerController(r *gin.Engine, customerService domain.CustomerService) {
	customerController := CustomerController{
		CustomerService: customerService,
	}

	v1public := r.Group("/api/v1/sks/public")
	{
		v1public.POST("/list/customer", customerController.GetDataCustomer)
	}
}

func (u *CustomerController) GetDataCustomer(c *gin.Context) {
	var request domain.AgentData
	idRequest := request.Id
	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}

	result, errorResponse := u.CustomerService.GetListCustomer(idRequest)
	if errorResponse.Message != nil {
		if errorResponse.Status == http.StatusNotFound {
			util.HandleError(c, errorResponse.Status, errorResponse.Status, errorResponse.Message.Error(), errorResponse.Message, errorResponse.Message.Error())
			return
		}
		util.HandleError(c, errorResponse.Status, errorResponse.Status, util.ERR_GENERAL, errorResponse.Message, errorResponse.Message.Error())
		return
	}
	util.HandleSuccess(c, http.StatusOK, 200, "Success Get Customer", result, "Success Get Customer", "Success Get Customer")

}

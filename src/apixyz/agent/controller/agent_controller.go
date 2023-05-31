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
		v1public.GET("/list/customer", customerController.GetDataCustomer)
		v1public.POST("/register/customer", customerController.InsertAgent)
		v1public.POST("/insert/limit", customerController.InsertTenor)
		v1public.POST("/list/limit", customerController.GetDataLimit)
		v1public.POST("/transaction", customerController.TransactionData)
		v1public.POST("/update/balance", customerController.UpdateBalance)
		v1public.POST("/payments", customerController.Payments)

	}
}

func (u *CustomerController) Payments(c *gin.Context) {
	var request domain.Requestpayments

	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}

	response, err := u.CustomerService.Paymenst(request.CustomerId, request.ContractNo, request.Installment, request.Ammount)

	if err.Message != nil {
		util.HandleError(c, err.Status, err.Status, util.ERR_GENERAL, err.Message, err.Message.Error())
		return
	}

	util.HandleSuccess(c, http.StatusOK, 200, "Update Success", response, "Update Success", "Update Success")
}

func (u *CustomerController) UpdateBalance(c *gin.Context) {
	var request domain.RequestUpdateAgentLimit
	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}

	response, err := u.CustomerService.UpdateCustomerLimitAmount(request.Customer, request.Tenor, request.LimitAmmount, request.Total)

	if err.Message != nil {
		util.HandleError(c, err.Status, err.Status, util.ERR_GENERAL, err.Message, err.Message.Error())
		return
	}

	util.HandleSuccess(c, http.StatusOK, 200, "Update Success", response, "Update Success", "Update Success")

}

func (u *CustomerController) TransactionData(c *gin.Context) {
	var request domain.TransactionsReq
	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}
	response, err := u.CustomerService.Transaction(request.CustomerId, request.ContractNo, request.Tenor, request.Otr, request.Assetname, request.TransactionType)

	if err.Message != nil {
		util.HandleError(c, err.Status, err.Status, util.ERR_GENERAL, err.Message, err.Message.Error())
		return
	}

	util.HandleSuccess(c, http.StatusOK, 200, "Transaction Success", response, "Transaction Success", "Transaction Success")
}

func (u *CustomerController) GetDataLimit(c *gin.Context) {
	var request domain.RequestIdTenor

	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}

	result, errorResponse := u.CustomerService.GetDataListTenor(request.Id)
	if errorResponse.Message != nil {
		if errorResponse.Status == http.StatusNotFound {
			util.HandleError(c, errorResponse.Status, errorResponse.Status, errorResponse.Message.Error(), errorResponse.Message, errorResponse.Message.Error())
			return
		}
		util.HandleError(c, errorResponse.Status, errorResponse.Status, util.ERR_GENERAL, errorResponse.Message, errorResponse.Message.Error())
		return
	}
	util.HandleSuccess(c, http.StatusOK, 200, "Success Get Limit Tenor", result, "Success Get Limit Tenor", "Success Get Limit Tenor")

}

func (u *CustomerController) InsertTenor(c *gin.Context) {
	var request domain.RequestIdTenor

	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}

	result, errorResponse := u.CustomerService.GetDataLimit(request.Id)
	if errorResponse.Message != nil {
		if errorResponse.Status == http.StatusNotFound {
			util.HandleError(c, errorResponse.Status, errorResponse.Status, errorResponse.Message.Error(), errorResponse.Message, errorResponse.Message.Error())
			return
		}
		util.HandleError(c, errorResponse.Status, errorResponse.Status, util.ERR_GENERAL, errorResponse.Message, errorResponse.Message.Error())
		return
	}
	util.HandleSuccess(c, http.StatusOK, 200, "Success Get Limit", result, "Success Get Limit", "Success Get Limit")

}

func (u *CustomerController) InsertAgent(c *gin.Context) {
	var request domain.AgentData
	if err := c.BindJSON(&request); err != nil {
		util.HandleError(c, http.StatusBadRequest, 400, util.ERR_BAD_REQUEST, err, util.ERR_BAD_REQUEST)
		return
	}
	response, err := u.CustomerService.InsertUpdateCustomer(request.Id, request.NIK, request.FullName, request.LegalName, request.PlaceOfBirth, request.BirthDate, request.Gaji, request.FotoKTP, request.FotoSelfie)

	if err.Message != nil {
		util.HandleError(c, err.Status, err.Status, util.ERR_GENERAL, err.Message, err.Message.Error())
		return
	}

	util.HandleSuccess(c, http.StatusOK, 200, "Success Insert/Update Agent", response, "Success Insert/Update Agent", "Success Insert/Update Agent")
}

func (u *CustomerController) GetDataCustomer(c *gin.Context) {

	result, errorResponse := u.CustomerService.GetListCustomer()
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

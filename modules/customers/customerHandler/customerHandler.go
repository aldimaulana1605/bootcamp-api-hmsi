package customerHandler

import (
	"net/http"

	"github.com/aldimaulana1605/bootcamp-api-hmsi/models"
	"github.com/aldimaulana1605/bootcamp-api-hmsi/modules/customers"
	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomerUsecase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomerUsecase) {
	handler := customerHandler{UC}
	
	r.GET("/customers", handler.GetAll)
	r.POST("/customers/insert", handler.Insert)
}

func (h *customerHandler) GetAll(c *gin.Context) {
	result, err := h.UC.FindAll()

	if err!= nil {
			c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": err.Error(),
			"data": []string{},
		})
		return
	}
    c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": result,
	})
}

func (h *customerHandler) Insert(c *gin.Context) {
	var request models.RequestInsertCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	err := h.UC.Insert(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Inserted successfully",
		"data":    []string{},
	})
}

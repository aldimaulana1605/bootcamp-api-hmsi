package customerHandler

import (
	"net/http"

	"github.com/aldimaulana1605/bootcamp-api-hmsi/modules/customers"
	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomerUsecase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomerUsecase) {
	handler := customerHandler{UC}
	
	r.GET("/customers", handler.GetAll)
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

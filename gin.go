package main

// Sending & Receiving Custom JSON
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define request struct
type OrderRequest struct { // binding:"required" : gin will validate it
	Stock    string  `json:"stock" binding:"required"`
	Type     string  `json:"type" binding:"required,oneof=buy sell"`
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

// Define response struct
type OrderResponse struct {
	FilledQty int `json:"filledqty"`
}

func processOrder(c *gin.Context) {
	var order OrderRequest

	// Bind incoming JSON to Order request struct & runs the validation : binding, oneof
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err) // Returns 400 status code along with error
	}

	filledOty := order.Quantity

	// Response that will be returned of type OrderResponse with value filledQty
	response := OrderResponse{
		FilledQty: filledOty,
	}

	c.JSON(http.StatusOK, response) // Send 200 status code along with response
}

func sendRequest2() {
	r := gin.Default() // Create gin router

	// Post request endpoint
	r.POST("/order", processOrder)

	// Port it should run at
	r.Run(":8080")
}

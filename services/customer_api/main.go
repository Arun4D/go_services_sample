package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// customer represents data about a record customer.
type customer struct {
	ID     string  `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName string  `json:"last_name"`
	DateOfBirth  string `json:"date_of_birth"`
}

// customers slice to seed record customer data.
var customers = []customer{
	{ID: "1", FirstName: "Blue Train", LastName: "John Coltrane", DateOfBirth: "01-01-1988"},
	{ID: "2", FirstName: "Jeru", LastName: "Gerry Mulligan", DateOfBirth: "01-01-1988"},
	{ID: "3", FirstName: "Sarah Vaughan and Clifford Brown", LastName: "Sarah Vaughan", DateOfBirth: "01-01-1988"},
}

// getcustomers responds with the list of all customers as JSON.
func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

// postcustomers adds an customer from JSON received in the request body.
func postCustomers(c *gin.Context) {
	var newCustomer customer

	// Call BindJSON to bind the received JSON to
	// newcustomer.
	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	// Add the new customer to the slice.
	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

// getcustomerByID locates the customer whose ID value matches the id
// parameter sent by the client, then returns that customer as a response.
func getCustomerByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of customers, looking for
	// an customer whose ID value matches the parameter.
	for _, a := range customers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}

func main() {
	router := gin.Default()
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomerByID)

	router.POST("/customers", postCustomers)

	router.Run("localhost:8080")
}

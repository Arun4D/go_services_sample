package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type CustomerEntity struct {
	ID            string `json:"id"`
	FIRST_NAME    string `json:"first_name"`
	LAST_NAME     string `json:"last_name"`
	DATE_OF_BIRTH string `json:"date_of_birth"`
}

var db *sql.DB

func db_connect() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "customer",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func getCustomers(c *gin.Context) {
	cus, err := getAllCustomers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
	} else {
		c.IndentedJSON(http.StatusOK, cus)
	}
}

func getAllCustomers() ([]CustomerEntity, error) {
	var Customers []CustomerEntity

	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, fmt.Errorf("getAllCustomers : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cus CustomerEntity
		if err := rows.Scan(&cus.ID, &cus.FIRST_NAME, &cus.LAST_NAME, &cus.DATE_OF_BIRTH); err != nil {
			return nil, fmt.Errorf("CustomersByArtist: %v", err)
		}
		Customers = append(Customers, cus)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("CustomersByArtist: %v", err)
	}
	return Customers, nil
}

func postCustomers(c *gin.Context) {
	var newCustomer CustomerEntity

	c.BindJSON(&newCustomer)
	fmt.Println("newCustomer!", newCustomer)
	cus, err := addCustomer(newCustomer)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Unable to insert customer"})
	} else {
		c.IndentedJSON(http.StatusOK, cus)
	}
}

func addCustomer(cus CustomerEntity) (int64, error) {
	result, err := db.Exec("INSERT INTO customer (FIRST_NAME, LAST_NAME, DATE_OF_BIRTH) VALUES (?, ?, ?)", cus.FIRST_NAME, cus.LAST_NAME, cus.DATE_OF_BIRTH)
	if err != nil {
		return 0, fmt.Errorf("addCustomer: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addCustomer: %v", err)
	}
	return id, nil
}

func getCustomerByID(c *gin.Context) {
	id := c.Param("id")

	cus, err := getCustomerDbByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
	} else {
		c.IndentedJSON(http.StatusOK, cus)
	}

}

func getCustomerDbByID(id string) (CustomerEntity, error) {
	var cus CustomerEntity

	row := db.QueryRow("SELECT * FROM customer WHERE id = ?", id)

	if err := row.Scan(&cus.ID, &cus.FIRST_NAME, &cus.LAST_NAME, &cus.DATE_OF_BIRTH); err != nil {
		if err == sql.ErrNoRows {
			return cus, fmt.Errorf("customerById %s: no such customer", id)
		}
		return cus, fmt.Errorf("customerById %s: %v", id, err)
	}
	return cus, nil

}

func main() {
	db_connect()
	router := gin.Default()
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomerByID)

	router.POST("/customers", postCustomers)

	router.Run("localhost:8080")
}

package main

import (
	"net/http" // Package for HTTP client and server implementations

	"github.com/gin-gonic/gin" // Web framework for building RESTful APIs
	"gorm.io/driver/sqlite"    // SQLite database driver for GORM
	"gorm.io/gorm"             // ORM (Object-Relational Mapping) library for Go
)

// User represents the user model in the database
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"` // Primary key field
	Name  string `json:"name"`                 // User's name
	Email string `json:"email"`                // User's email
}

func main() {
	r := gin.Default() // Create a default Gin router instance

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		// Panic if database connection fails
		panic("failed to connect to database")
	}

	// AutoMigrate creates or updates the database table schema for the User model
	db.AutoMigrate(&User{})

	// Route to fetch all users
	r.GET("/users", func(c *gin.Context) {
		var users []User             // Slice to hold users
		db.Find(&users)              // Fetch all users from the database
		c.JSON(http.StatusOK, users) // Return the users in JSON format
	})

	// Route to fetch a user by ID
	r.GET("/users/:id", func(c *gin.Context) {
		var user User
		// Find the user by ID; if not found, return a 404 error
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Return the found user in JSON format
		c.JSON(http.StatusOK, user)
	})

	// Route to create a new user
	r.POST("/users", func(c *gin.Context) {
		var user User
		// Bind the JSON payload to the User struct
		if err := c.ShouldBindJSON(&user); err != nil {
			// Return a 400 error if binding fails
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Save the new user to the database
		db.Create(&user)
		// Return the created user with a 201 status code
		c.JSON(http.StatusCreated, user)
	})

	// Route to update an existing user by ID
	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		// Find the user by ID; if not found, return a 404 error
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Bind the JSON payload to the existing user struct
		if err := c.ShouldBindJSON(&user); err != nil {
			// Return a 400 error if binding fails
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Update the user in the database
		db.Save(&user)
		// Return the updated user
		c.JSON(http.StatusOK, user)
	})

	// Route to delete a user by ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		// Delete the user by ID; if an error occurs, return a 500 error
		if err := db.Delete(&User{}, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		// Return a 204 No Content status if deletion is successful
		c.Status(http.StatusNoContent)
	})

	// Start the server on port 8080
	r.Run(":8080")
}

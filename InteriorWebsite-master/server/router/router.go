package router

import (
	"main/db"
	"main/tables"
	"net/http"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", CreateUser)
	r.POST("/login", LoginUser)
	r.GET("/projects", Projects)
	return r
}

func CreateUser(c *gin.Context) {
	var values map[string]string
	if err := c.ShouldBindJSON(&values); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	var user tables.User

	result := db.DB.Where("email = ?", values["email"]).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		user = tables.User{
			Fname:    values["fname"],
			Lname:    values["lname"],
			Email:    values["email"],
			Password: values["password"],
			Phone:    values["phone"],
		}
		db.DB.Create(&user)
		c.JSON(http.StatusAccepted, gin.H{
			"email": user.Email,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exist",
		})
	}

}

func LoginUser(c *gin.Context) {
	var values map[string]string
	if err := c.ShouldBindJSON(&values); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	var user tables.User

	result := db.DB.Where("email = ?", values["email"]).Where("password = ?", values["password"]).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"msg": "Login succesful",
		"email": values["email"]})

}

func Projects(c *gin.Context) {

	designer := tables.Designer{
		Name:  "Richard Frank",
		Email: "Richard.f@gmail.com",
	}

	db.DB.Create(&designer)

	value := tables.Projects{
		Name:        "Home Decore",
		Designer_id: 1,
	}

	db.DB.Create(&value)

	var projects []tables.Projects

	db.DB.Find(&projects)

	c.JSON(http.StatusAccepted, gin.H{
		"projetcs": projects,
	})
}

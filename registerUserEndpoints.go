package basego

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (a *Api) registerUserEndpoints(g *gin.RouterGroup) {

	g.POST("/create-user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.CreateAt = time.Now()

		if err := a.repo.CreateNewUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})

	})

	g.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		user, err := a.repo.GetUserByID(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})

	})

	g.PATCH("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := a.repo.UpdateUserByID(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})

	})

	g.DELETE("/delete-user/:id", func(c *gin.Context) {
		id := c.Param("id")

		if err := a.repo.DeleteUserByID(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Delete user",
		})
	})
}

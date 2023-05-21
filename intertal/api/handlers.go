package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *api) SendEmail(c *gin.Context) {
	req := &CreateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	err := a.app.SendEmail(c, req.Contact, "Title", "Contact")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, nil)
}

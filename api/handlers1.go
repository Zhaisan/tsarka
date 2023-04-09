package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (server *Server) checkEmailsHandler(ctx *gin.Context) {
	var reqBody struct {
		Message string `json:"message"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	re := regexp.MustCompile(`Email:\s*[\w.%+-]+@[\w.-]+\.[a-zA-Z]{2,6}`)
	emailsStr := re.FindAllString(reqBody.Message, -1)

	emails := make([]Email, 0, len(emailsStr))
	for _, emailStr := range emailsStr {
		email := Email{Address: emailStr}
		emails = append(emails, email)
	}

	ctx.JSON(http.StatusOK, gin.H{"emails": emails})
}

type Email struct {
	Address string `json:"address"`
}
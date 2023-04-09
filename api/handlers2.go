package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

func (server *Server) checkIINHandler(ctx *gin.Context) {
	var reqBody struct {
		Text string `json:"text"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	re := regexp.MustCompile(`^(0[1-9]|[1-2][0-9]|3[0-1])(0[1-9]|1[0-2])(\d{2})(\d{3})(\d)$`)
	matches := re.FindAllString(reqBody.Text, -1)

	var iins []string
	for _, match := range matches {
		if isValidIIN(match) {
			iins = append(iins, match)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"iins": iins})
}

func isValidIIN(iin string) bool {
	re := regexp.MustCompile(`^(0[1-9]|[1-2][0-9]|3[0-1])(0[1-9]|1[0-2])(\d{2})(\d{3})(\d)$`)
	if !re.MatchString(iin) {
		return false
	}

	// Проверяем, что первые 6 цифр соответствуют дате рождения
	day, _ := strconv.Atoi(iin[0:2])
	month, _ := strconv.Atoi(iin[2:4])
	year, _ := strconv.Atoi(iin[4:6])
	if !isDateValid(day, month, year) {
		return false
	}

	// Проверяем, что код региона рождения корректный
	regionCode, _ := strconv.Atoi(iin[6:8])
	if regionCode < 1 || regionCode > 20 {
		return false
	}

	// Проверяем, что порядковый номер корректный
	orderNumber, _ := strconv.Atoi(iin[8:11])
	if orderNumber < 1 || orderNumber > 999 {
		return false
	}


	return true
}

func isDateValid(day, month, year int) bool {
	// Проверяем, что год состоит из 2 цифр
	if year < 0 || year > 99 {
		return false
	}

	// Проверяем, что месяц от 1 до 12
	if month < 1 || month > 12 {
		return false
	}

	// Проверяем, что день от 1 до 31 и соответствует месяцу и году
	if day < 1 || day > 31 {
		return false
	}
	if month == 2 && day > 29 {
		return false
	}
	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return false
	}

	if (year % 4 == 0 && day > 29) || (year % 4 != 0 && day > 28) {
		return false
	}

	return true
}


package api

import (
	db "github.com/Zhaisan/tsarka_test/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func findMaxSubstring(s string) string {
	var maxSubstring string
	var currentSubstring string
	charMap := make(map[rune]int)

	for i, c := range s {
		if _, ok := charMap[c]; ok && charMap[c] >= i - len(currentSubstring) {
			currentSubstring = s[charMap[c] + 1 : i]
		}
		currentSubstring += string(c)
		charMap[c] = i
		if len(currentSubstring) > len(maxSubstring) {
			maxSubstring = currentSubstring
		}
	}
	return maxSubstring
}

func (server *Server) findSubstring(ctx *gin.Context) {
	var requestBody struct {
		String string `json:"string"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	maxSubstring := findMaxSubstring(requestBody.String)
	ctx.JSON(http.StatusOK, gin.H{"max_substring": maxSubstring})
}

func (server *Server) createString(ctx *gin.Context) {
	var requestBody struct {
		String string `json:"string"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	maxSubstring := findMaxSubstring(requestBody.String)
	result, err := server.db.CreateString(ctx, db.CreateStringParams{
		String:       requestBody.String,
		MaxSubstring: maxSubstring,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) deleteString(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := server.db.DeleteString(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) getString(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := server.db.GetString(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"max_substring": result})
}

func (server *Server) listStrings(ctx *gin.Context) {
	var requestBody struct {
		Limit  int32 `json:"limit"`
		Offset int32 `json:"offset"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := server.db.ListStrings(ctx, db.ListStringsParams{
		Limit:  requestBody.Limit,
		Offset: requestBody.Offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) updateString(ctx *gin.Context) {
	var requestBody struct {
		MaxSubstring string `json:"max_substring"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := server.db.UpdateString(ctx, db.UpdateStringParams{
		MaxSubstring: requestBody.MaxSubstring,
		ID:           int32(id),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

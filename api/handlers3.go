package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func (server *Server) findIdentifiersHandler(ctx *gin.Context) {
	substring := ctx.Param("str")

	identifiers := []string{}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			bytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			re := regexp.MustCompile(fmt.Sprintf(`\b\w*%s\w*\b`, substring))
			matches := re.FindAllString(string(bytes), -1)

			identifiers = append(identifiers, matches...)
		}

		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"identifiers": identifiers})
}

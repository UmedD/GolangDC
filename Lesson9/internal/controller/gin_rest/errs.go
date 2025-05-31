package gin_rest

import (
	errors2 "GolangDC/Lesson9/internal/errors"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerErr(c *gin.Context, err error) {
	if err == nil {
		return
	} else if errors.As(err, &errors2.ErrValidationFailed) || errors.Is(err, errors2.ErrInvalidOperationType) ||
		errors.Is(err, errors2.ErrNotEnoughBalance) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errors2.ErrAccountNotFound) || errors.Is(err, errors2.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("something went wrong: %s", err.Error()),
		})
	}
}

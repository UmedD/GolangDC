package gin_rest

import (
	errors2 "GolangDC/Lesson9/internal/errors"
	"GolangDC/Lesson9/internal/models"
	"GolangDC/Lesson9/internal/repository/MemoryRepo"
	"GolangDC/Lesson9/internal/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "start server",
	})
}

func GetAllAccounts(c *gin.Context) {
	accounts, err := service.GetAll()
	if err != nil {
		HandlerErr(c, err)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func GetAccountByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errors2.ErrValidationFailed)
		HandlerErr(c, err)
		return
	}
	account, err := service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func CreateAccount(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	idSeq := MemoryRepo.GetMaxID()
	idSeq++
	note.ID = idSeq
	note.CreatedAt = time.Now()
	service.Create(note)
	c.JSON(http.StatusBadRequest, gin.H{"succes": "new todo add"})

}

func DeleteAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errors2.ErrValidationFailed)
		HandlerErr(c, err)
		return
	}

	service.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprintf("Note with ID %d deleted", id),
		"id":      id,
		"status":  "success",
	}
	c.JSON(http.StatusOK, response)
}

func UpdateAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errors2.ErrValidationFailed)
		HandlerErr(c, err)
		return
	}

	var account models.Note
	account, err = MemoryRepo.GetByID(id)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errors2.ErrUserNotFound)
		HandlerErr(c, err)
		return
	}

	var input struct {
		Done bool `json:"done"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		HandlerErr(c, err)
		return
	}

	account.Done = input.Done
	service.UpdateAccount(account)
	c.JSON(http.StatusOK, gin.H{
		"note": account,
	})
}

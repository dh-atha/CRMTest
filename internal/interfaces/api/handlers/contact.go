package handlers

import (
	"net/http"
	"strconv"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	service service.ContactServiceInterface
}

func NewContactHandler(ContactService service.ContactServiceInterface) *ContactHandler {
	return &ContactHandler{service: ContactService}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var emp model.Contact
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ContactID, err := h.service.Create(c, &emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success create Contact", gin.H{
		"Contact_id": ContactID,
	}))
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Contact ID"})
		return
	}

	var emp model.Contact
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.ContactID = id

	if err := h.service.Update(c, &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success update contact", nil))
}

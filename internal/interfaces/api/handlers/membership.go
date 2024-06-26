package handlers

import (
	"fmt"
	"net/http"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type MembershipHandler struct {
	service service.MembershipServiceInterface
}

func NewMembershipHandler(MembershipService service.MembershipServiceInterface) *MembershipHandler {
	return &MembershipHandler{service: MembershipService}
}

func (h *MembershipHandler) CreateMembership(c *gin.Context) {
	var emp model.Membership
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(h)

	membership, err := h.service.Create(c, &emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success create membership", membership))
}

func (h *MembershipHandler) GetAllMemberships(c *gin.Context) {
	ctx := c.Request.Context()
	Memberships, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Membership", Memberships))
}

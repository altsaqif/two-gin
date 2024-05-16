package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"enigma.com/two-gin/model/dto"
	"enigma.com/two-gin/model/dto/commonResponse"
	"enigma.com/two-gin/service"
)

// 1. Struct
// 2. Routing
// 3. Method
// 4. Function

type authorController struct {
	authorUseCase service.AuthorUseCase
	router        *gin.RouterGroup
}

func (a *authorController) listHandler(c *gin.Context) {
	page, er := strconv.Atoi(c.Query("page"))
	if er != nil {
		commonResponse.SendErrorResponse(c, http.StatusBadRequest, er.Error())
	}

	size, er2 := strconv.Atoi(c.Query("size"))
	if er2 != nil {
		commonResponse.SendErrorResponse(c, http.StatusBadRequest, er.Error())
	}

	listData, paging, err := a.authorUseCase.FindAll(page, size)
	if err != nil {
		commonResponse.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var data []any
	for _, b := range listData {
		data = append(data, b)
	}

	commonResponse.SendManyResponse(c, data, paging, "Ok")
}

func (a *authorController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	data, err := a.authorUseCase.FindById(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, dto.SingleResponse{
		Status: dto.Status{
			Code:    http.StatusOK,
			Message: "Ok",
		},
		Data: data,
	})
}

func (a *authorController) Routing() {
	a.router.GET("/authors", a.listHandler)
	a.router.GET("/authors/:id", a.getByIdHandler)
}

func NewAuthorController(authorUc service.AuthorUseCase, rg *gin.RouterGroup) *authorController {
	return &authorController{
		authorUseCase: authorUc,
		router:        rg,
	}
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"enigma.com/two-gin/model"
	"enigma.com/two-gin/model/dto"
	"enigma.com/two-gin/model/dto/commonResponse"
	"enigma.com/two-gin/service"
)

// 1. Struct
// 2. Routing
// 3. Method
// 4. Function

type tasksController struct {
	tasksService service.TasksService
	router       *gin.RouterGroup
}

func (a *tasksController) createHandlerTasks(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := a.tasksService.CreateTasks(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTask)
}

func (a *tasksController) listHandlerTasks(c *gin.Context) {
	page, er := strconv.Atoi(c.Query("page"))
	if er != nil {
		commonResponse.SendErrorResponse(c, http.StatusBadRequest, er.Error())
	}

	size, er2 := strconv.Atoi(c.Query("size"))
	if er2 != nil {
		commonResponse.SendErrorResponse(c, http.StatusBadRequest, er.Error())
	}

	listData, paging, err := a.tasksService.FindAllTasks(page, size)
	if err != nil {
		commonResponse.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var data []any
	for _, b := range listData {
		data = append(data, b)
	}

	commonResponse.SendManyResponse(c, data, paging, "Ok")
}

func (a *tasksController) getByIdHandlerTasks(c *gin.Context) {
	id := c.Param("id")
	data, err := a.tasksService.FindByIdTasks(id)
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

func (a *tasksController) RoutingTasks() {
	a.router.GET("/tasks", a.listHandlerTasks)
	a.router.GET("/tasks/:id", a.getByIdHandlerTasks)
	a.router.POST("/tasks", a.createHandlerTasks)
}

func NewTasksController(tasksUc service.TasksService, rg *gin.RouterGroup) *tasksController {
	return &tasksController{
		tasksService: tasksUc,
		router:       rg,
	}
}

package config

import (
	"database/sql"
	"fmt"

	"enigma.com/two-gin/controller"
	"enigma.com/two-gin/middleware"
	"enigma.com/two-gin/repository"
	"enigma.com/two-gin/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	tasksUC  service.TasksService
	authorUC service.AuthorUseCase
	engine   *gin.Engine
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")
	rg.Use(middleware.LogMiddleware())
	controller.NewAuthorController(s.authorUC, rg).Routing()
	controller.NewTasksController(s.tasksUC, rg).RoutingTasks()
}

func (s *Server) Run() {
	s.initRoute()
	s.engine.Run(":3000")
}

func NewServer() *Server {
	c, _ := NewConfig()

	urlConnect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.DbPort, c.DbUser, c.DbPassword, c.DbName)

	database, err := sql.Open(c.Driver, urlConnect)
	if err != nil {
		panic("Connection error")
	}

	authorRepo := repository.NewAuthorRepo(database)
	tasksRepo := repository.NewTasksRepo(database)
	authorUC := service.NewAuthorUseCase(authorRepo)
	tasksUC := service.NewTasksService(tasksRepo)

	return &Server{
		authorUC: authorUC,
		tasksUC:  tasksUC,
		engine:   gin.Default(),
	}
}

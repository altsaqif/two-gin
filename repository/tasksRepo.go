package repository

import (
	"database/sql"
	"log"
	"math"
	"time"

	"enigma.com/two-gin/model"
	"enigma.com/two-gin/model/dto"
)

// 1. Buat struct
// 2. Interface => kontrak
// 3. Method
// 4. Function

type tasksRepo struct {
	db *sql.DB
}

// CreateTasks implements TasksRepo.
func (a *tasksRepo) CreateTasks(task model.Task) (model.Task, error) {
	stmt, err := a.db.Prepare("INSERT INTO trx_tasks(title, content, author_id) VALUES($1, $2, $3) RETURNING id")
	if err != nil {
		return model.Task{}, err
	}
	defer stmt.Close()

	var taskId string
	err = stmt.QueryRow(task.Title, task.Content, task.AuthorId).Scan(&taskId)
	if err != nil {
		return model.Task{}, err
	}

	task.Id = taskId
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return task, nil
}

// findAll implements TaskRepo.
func (a *tasksRepo) FindAllTasks(page int, size int) ([]model.Task, dto.Paging, error) {
	var listData []model.Task
	var rows *sql.Rows
	var err error

	// Rumus Pagiation
	offset := (page - 1) * size

	rows, err = a.db.Query("SELECT * FROM trx_tasks limit $1 offset $2", size, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	totalRows := 0
	err = a.db.QueryRow("SELECT COUNT(*) FROM trx_tasks").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var tasks model.Task

		err := rows.Scan(&tasks.Id, &tasks.Title, &tasks.Content, &tasks.AuthorId, &tasks.CreatedAt, &tasks.UpdatedAt)

		if err != nil {
			log.Println(err.Error())
		}

		listData = append(listData, tasks)
	}

	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return listData, paging, nil
}

// findById implements AuthorRepo.
func (a *tasksRepo) FindByIdTasks(id string) (model.Task, error) {
	var tasks model.Task

	err := a.db.QueryRow("SELECT * FROM trx_tasks WHERE id = $1", id).Scan(&tasks.Id, &tasks.Title, &tasks.Content, &tasks.AuthorId, &tasks.CreatedAt, &tasks.UpdatedAt)

	if err != nil {
		return model.Task{}, err
	}

	return tasks, nil
}

type TasksRepo interface {
	FindAllTasks(page int, size int) ([]model.Task, dto.Paging, error)
	FindByIdTasks(id string) (model.Task, error)
	CreateTasks(model.Task) (model.Task, error)
}

// Construction => Gerbang untuk mengakses repository
func NewTasksRepo(database *sql.DB) TasksRepo {
	return &tasksRepo{db: database}
}

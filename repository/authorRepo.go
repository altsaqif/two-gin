package repository

import (
	"database/sql"
	"log"
	"math"

	"enigma.com/two-gin/model"
	"enigma.com/two-gin/model/dto"
)

// 1. Buat struct
// 2. Interface => kontrak
// 3. Method
// 4. Function

type authorRepo struct {
	db *sql.DB
}

// findAll implements AuthorRepo.
func (a *authorRepo) FindAll(page int, size int) ([]model.Author, dto.Paging, error) {
	var listData []model.Author
	var rows *sql.Rows
	var err error

	// Rumus Pagiation
	offset := (page - 1) * size

	rows, err = a.db.Query("SELECT * FROM mst_authors limit $1 offset $2", size, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	totalRows := 0
	err = a.db.QueryRow("SELECT COUNT(*) FROM mst_authors").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var author model.Author

		err := rows.Scan(&author.Id, &author.Fullname, &author.Email, &author.Password, &author.CreatedAt, &author.UpdatedAt, &author.Role)

		if err != nil {
			log.Println(err.Error())
		}

		listData = append(listData, author)
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
func (a *authorRepo) FindById(id string) (model.Author, error) {
	var author model.Author

	err := a.db.QueryRow("SELECT * FROM mst_authors WHERE id = $1", id).Scan(&author.Id, &author.Fullname, &author.Email, &author.Password, &author.CreatedAt, &author.UpdatedAt, &author.Role)

	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

type AuthorRepo interface {
	FindAll(page int, size int) ([]model.Author, dto.Paging, error)
	FindById(id string) (model.Author, error)
}

// Construction => Gerbang untuk mengakses repository
func NewAuthorRepo(database *sql.DB) AuthorRepo {
	return &authorRepo{db: database}
}

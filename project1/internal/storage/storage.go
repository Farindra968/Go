package storage

import "github.com/Farindra968/go_project1/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int, password string) (int64, error)
	GetStudentByID(id string) (types.Student, error)
}
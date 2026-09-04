package storage

type Storage interface {
	CreateStudent(name string, email string, age int, password string) (int64, error)
}
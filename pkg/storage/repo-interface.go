package storage

import (
	"github.com/gal16v8d/app-registry.git/internal/domain"
)

type RepoStorageInterface interface {
	GetAll() ([]domain.Repo, error)
	GetById(id int) (domain.Repo, error)
	CreateRepo(repo domain.Repo) error
	UpdateRepo(repo domain.Repo) error
	DeleteRepo(id int) error
}

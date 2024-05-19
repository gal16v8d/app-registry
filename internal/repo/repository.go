package odontologo

import (
	"errors"

	"github.com/gal16v8d/app-registry.git/internal/domain"
	"github.com/gal16v8d/app-registry.git/pkg/storage"
)

type Repository interface {
	GetById(id int) (domain.Repo, error)
	CreateRepo(repo domain.Repo) (domain.Repo, error)
	UpdateRepo(id int, repo domain.Repo) (domain.Repo, error)
	DeleteRepo(id int) (string, error)
}

type repository struct {
	storage storage.RepoStorageInterface
}

func NewRepository(storage storage.RepoStorageInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetById(id int) (domain.Repo, error) {
	repo, err := r.storage.GetById(id)
	if err != nil {
		return domain.Repo{}, errors.New("repo not found")
	}
	return repo, nil
}

func (r *repository) CreateRepo(repo domain.Repo) (domain.Repo, error) {
	err := r.storage.CreateRepo(repo)
	if err != nil {
		return domain.Repo{}, errors.New("Error creating repo")
	}
	return repo, nil
}

func (r *repository) UpdateRepo(id int, repo domain.Repo) (domain.Repo, error) {
	err := r.storage.UpdateRepo(repo)
	if err != nil {
		return domain.Repo{}, errors.New("Error updating repo")
	}
	return repo, nil
}

func (r *repository) DeleteRepo(id int) (string, error) {
	err := r.storage.DeleteRepo(id)
	if err != nil {
		return "", err
	}
	return "", nil
}

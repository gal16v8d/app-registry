package odontologo

import (
	"github.com/gal16v8d/app-registry.git/internal/domain"
)

type Service interface {
	GetById(id int) (domain.Repo, error)
	CreateRepo(repo domain.Repo) (domain.Repo, error)
	UpdateRepo(id int, repo domain.Repo) (domain.Repo, error)
	DeleteRepo(id int) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

// @Summary Get a repo by id
// @Description Get a repo by id
// @Param id path int true "repo id"
// @Success 200 {object} domain.Repo
// @Router /repos/{id} [get]
func (s *service) GetById(id int) (domain.Repo, error) {
	d, err := s.r.GetById(id)
	if err != nil {
		return domain.Repo{}, err
	}
	return d, nil
}

// @Summary Create a new repo
// @Description Create a new repo
// @Param repo body domain.Repo true "Repo object"
// @Success 201 {object} domain.Repo
// @Router /repos [post]
func (s *service) CreateRepo(repo domain.Repo) (domain.Repo, error) {
	repo, err := s.r.CreateRepo(repo)
	if err != nil {
		return domain.Repo{}, err
	}
	return repo, nil
}

// @Summary Update a repo
// @Description Update a repo
// @Param id path int true "Repo id"
// @Param repo body domain.Repo true "Updated repo object"
// @Success 200 {object} domain.Repo
// @Router /repos/{id} [put]
func (s *service) UpdateRepo(id int, repo domain.Repo) (domain.Repo, error) {

	dbData, err := s.r.GetById(id)
	if err != nil {
		return domain.Repo{}, err
	}
	if repo.Name != "" {
		dbData.Name = repo.Name
	}
	if repo.MainTech != "" {
		dbData.MainTech = repo.MainTech
	}

	dbData, err = s.r.UpdateRepo(id, dbData)

	if err != nil {
		return domain.Repo{}, err
	}
	return dbData, nil
}

// @Summary Delete a repo
// @Description Delete a repo by id
// @Param id path int true "Repo id"
// @Success 204
// @Router /repos/{id} [delete]
func (s *service) DeleteRepo(id int) (string, error) {
	message, err := s.r.DeleteRepo(id)
	if err != nil {
		return "id doesn't exists ", err
	}
	return message, nil
}

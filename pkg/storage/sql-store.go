package storage

import (
	"database/sql"
	"fmt"

	"github.com/gal16v8d/app-registry.git/internal/domain"
)

type sqlStorage struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) RepoStorageInterface {
	return &sqlStorage{
		db: db,
	}
}

func (s *sqlStorage) GetAll() ([]domain.Repo, error) {
	var repos []domain.Repo
	query := "SELECT * FROM repo;"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var repo domain.Repo
		err := rows.Scan(&repo.Id, &repo.Name, &repo.MainTech)
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return repos, nil
}

func (s *sqlStorage) GetById(id int) (domain.Repo, error) {
	var repo domain.Repo
	query := "SELECT * FROM repo WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&repo.Id, &repo.Name, &repo.MainTech)
	if err != nil {
		return domain.Repo{}, err
	}
	return repo, nil
}

func (s *sqlStorage) CreateRepo(repo domain.Repo) error {
	query := "INSERT INTO repo (name, main_tech) VALUES (?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(repo.Name, repo.MainTech)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStorage) UpdateRepo(repo domain.Repo) error {
	query := "UPDATE repo SET name = ?, main_tech = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(repo.Name, repo.MainTech, repo.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStorage) DeleteRepo(id int) error {
	return s.DeleteById(id, "repo")
}

func (s *sqlStorage) DeleteById(id int, table string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?;", table)
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

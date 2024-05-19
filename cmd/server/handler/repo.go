package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gal16v8d/app-registry.git/internal/domain"
	"github.com/gal16v8d/app-registry.git/internal/repo"
	"github.com/gin-gonic/gin"
)

type repoHandler struct {
	s repo.Service
}

func NewRepoHandler(s repo.Service) *repoHandler {
	return &repoHandler{
		s: s,
	}
}

func validateEmptyRepo(repo *domain.Repo) (bool, error) {
	if repo.Name == "" || repo.MainTech == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func (h *repoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		repos, err := h.s.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving repos"})
			return
		}
		c.JSON(http.StatusOK, repos)
	}
}

func (h *repoHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		repo, err := h.s.GetById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "repo not found"})
			return
		}
		c.JSON(http.StatusOK, repo)
	}
}

func (h *repoHandler) CreateRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var repo domain.Repo
		err := ctx.ShouldBindJSON(&repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid repo"})
			return
		}
		valid, err := validateEmptyRepo(&repo)
		if !valid {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		r, err := h.s.CreateRepo(repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, r)
	}
}

func (h *repoHandler) UpdateRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var repo domain.Repo
		err = ctx.ShouldBindJSON(&repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid repo"})
			return
		}

		valid, err := validateEmptyRepo(&repo)
		if !valid {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		r, err := h.s.UpdateRepo(id, repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, r)
	}
}

func (h *repoHandler) DeleteRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		_, err = h.s.DeleteRepo(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"msg": "repo deleted"})
	}
}

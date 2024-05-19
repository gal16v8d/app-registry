package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gal16v8d/app-registry.git/internal/domain"
	"github.com/gal16v8d/app-registry.git/internal/repo"
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

func (h *repoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		repo, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "repo not found"})
			return
		}
		c.JSON(200, repo)
	}
}

func (h *repoHandler) CreateRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var repo domain.Repo
		err := ctx.ShouldBindJSON(&repo)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid repo"})
			return
		}
		valid, err := validateEmptyRepo(&repo)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		r, err := h.s.CreateRepo(repo)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, r)
	}
}

func (h *repoHandler) UpdateRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		var repo domain.Repo
		err = ctx.ShouldBindJSON(&repo)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid repo"})
			return
		}

		valid, err := validateEmptyRepo(&repo)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		r, err := h.s.UpdateRepo(id, repo)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, r)
	}
}

func (h *repoHandler) DeleteRepo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		_, err = h.s.DeleteRepo(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(204, gin.H{"msg": "repo deleted"})
	}
}
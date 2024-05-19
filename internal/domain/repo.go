package domain

type Repo struct {
	Id        int     `json: "id"`
	Name      string  `json: "name" binding:"required"`
	MainTech  string  `json: "mainTech" binding:"required"`
}
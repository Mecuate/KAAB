package models

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Content2 string `json:"content2"`
	Content3 string `json:"content3"`
}

type Tasks []Task

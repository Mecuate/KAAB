package fixtures

import "go_server/src/models"

var TasksData = models.Tasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

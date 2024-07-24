package model

type Task struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Priority          string `json:"priority"`
	State             string `json:"state"`
	ResponsiblePerson string `json:"responsible_person"`
	Project           string `json:"project"`
	CreatedAt         string `json:"created_at"`
	FinishedAt        string `json:"finished_at"`
}

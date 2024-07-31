package model

type Task struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Priority          string `json:"priority"`
	State             string `json:"state"`
	ResponsiblePerson string `json:"responsible_person"`
	Project           string `json:"project"`
	CreatedAt         string `json:"created_at" datetime_format:"YYYY-MM-DD"`
	FinishedAt        string `json:"finished_at" datetime_format:"YYYY-MM-DD"`
}

type TaskInputResponse struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Priority          string `json:"priority"`
	State             string `json:"state"`
	ResponsiblePerson string `json:"responsible_person"`
	Project           string `json:"project"`
	CreatedAt         string `json:"created_at" datetime_format:"YYYY-MM-DD"`
	FinishedAt        string `json:"finished_at" datetime_format:"YYYY-MM-DD"`
}

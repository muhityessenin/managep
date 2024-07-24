package model

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	FinishDate  string `json:"finishDate"`
	Manager     string `json:"manager"`
}

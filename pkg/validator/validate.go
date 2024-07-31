package validator

import (
	"managep/pkg/model"
	"regexp"
	"time"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateUserInput(user model.User) bool {
	if user.FullName == "" {
		return false
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(user.Email) {
		return false
	}
	t, err := time.Parse("2006-01-02", user.RegistrationDate)
	if err != nil {
		return false
	}
	if t.Before(time.Now()) {
		return false
	}
	if user.Role == "" {
		return false
	}
	return true
}

func (v *Validator) ValidateTaskInput(task model.Task) bool {
	if task.Name == "" {
		return false
	}
	validPriorities := map[string]bool{"Low": true, "Medium": true, "High": true, "low": true, "medium": true, "high": true}
	if !validPriorities[task.Priority] {
		return false
	}
	validStates := map[string]bool{"New": true, "In Progress": true, "Completed": true, "new": true, "in progress": true, "completed": true}
	if !validStates[task.State] {
		return false
	}
	if task.ResponsiblePerson == "" {
		return false
	}
	if task.Project == "" {
		return false
	}
	t, err := time.Parse("2006-01-02", task.CreatedAt)
	if err != nil {
		return false
	}
	if t.Before(time.Now()) {
		return false
	}
	return true
}

func (v *Validator) ValidateProjectInput(project model.Project) bool {
	if project.Name == "" {
		return false
	}
	if project.Manager == "" {
		return false
	}
	startDate, err := time.Parse("2006-01-02", project.StartDate)
	if err != nil {
		return false
	}
	if project.FinishDate != "" {
		finishDate, err := time.Parse("2006-01-02", project.FinishDate)
		if err != nil {
			return false
		}
		if finishDate.Before(startDate) {
			return false
		}
	}
	if startDate.Before(time.Now()) {
		return false
	}
	return true
}

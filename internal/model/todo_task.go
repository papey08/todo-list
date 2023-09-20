package model

// TodoTask is a struct for planning task
type TodoTask struct {
	Id           int
	Title        string
	Description  string
	PlanningDate Date
	Status       bool
}

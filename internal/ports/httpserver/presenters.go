package httpserver

type addTaskRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PlanningDate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"planning_date"`
	Status bool `json:"status"`
}

type getTaskByTextRequest struct {
	Text string `json:"text"`
}

type updateTaskRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PlanningDate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"planning_date"`
	Status bool `json:"status"`
}

type getTasksByStatusRequest struct {
	Status bool `json:"status"`
	Offset int  `json:"offset"`
	Limit  int  `json:"limit"`
}

type getTasksByDateAndStatusRequest struct {
	PlanningDate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"planning_date"`
	Status bool `json:"status"`
}

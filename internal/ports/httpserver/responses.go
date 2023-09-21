package httpserver

import (
	"todo-list/internal/model"
)

type taskData struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PlanningDate struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"planning_date"`
	Status bool `json:"status"`
}

type taskResponse struct {
	Data *taskData `json:"data"`
	Err  *string   `json:"error"`
}

type tasksResponse struct {
	Data []taskData `json:"data"`
	Err  *string    `json:"error"`
}

func taskSuccessResponse(t model.TodoTask) taskResponse {
	return taskResponse{
		Data: &taskData{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			PlanningDate: struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			}{
				Year:  t.PlanningDate.Year,
				Month: int(t.PlanningDate.Month),
				Day:   t.PlanningDate.Day,
			},
			Status: t.Status,
		},
		Err: nil,
	}
}

func tasksSuccessResponse(tasks []model.TodoTask) tasksResponse {
	resp := make([]taskData, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, taskData{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			PlanningDate: struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			}{
				Year:  t.PlanningDate.Year,
				Month: int(t.PlanningDate.Month),
				Day:   t.PlanningDate.Day,
			},
			Status: t.Status,
		})
	}
	return tasksResponse{
		Data: resp,
		Err:  nil,
	}
}

func errorResponse(err error) taskResponse {
	errStr := err.Error()
	return taskResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func deleteSuccessResponse() taskResponse {
	return taskResponse{
		Data: nil,
		Err:  nil,
	}
}

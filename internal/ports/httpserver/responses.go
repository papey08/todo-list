package httpserver

import (
	"github.com/gin-gonic/gin"
	"todo-list/internal/model"
)

type taskResponse struct {
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

type tasksResponse []taskResponse

func taskSuccessResponse(t model.TodoTask) *gin.H {
	return &gin.H{
		"data": taskResponse{
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
		"error": nil,
	}
}

func tasksSuccessResponse(tasks []model.TodoTask) *gin.H {
	resp := make(tasksResponse, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, taskResponse{
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
	return &gin.H{
		"data":  resp,
		"error": nil,
	}
}

func errorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func deleteSuccessResponse() *gin.H {
	return &gin.H{
		"data":  nil,
		"error": nil,
	}
}

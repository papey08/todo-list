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
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
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
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			}{
				Day:   t.PlanningDate.Day,
				Month: int(t.PlanningDate.Month),
				Year:  t.PlanningDate.Year,
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
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			}{
				Day:   t.PlanningDate.Day,
				Month: int(t.PlanningDate.Month),
				Year:  t.PlanningDate.Year,
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

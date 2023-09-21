package httpserver

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"todo-list/internal/app"
	"todo-list/internal/model"
)

func addTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req addTaskRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		t, err := a.AddTask(c, model.TodoTask{
			Title:       req.Title,
			Description: req.Description,
			PlanningDate: model.Date{
				Day:   req.PlanningDate.Day,
				Month: time.Month(req.PlanningDate.Month),
				Year:  req.PlanningDate.Year,
			},
			Status: req.Status,
		})

		switch {
		case errors.Is(err, model.ErrInvalidTask):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidTask))
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, taskSuccessResponse(t))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

func getTaskById(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		t, err := a.GetTaskById(c, id)

		switch {
		case errors.Is(err, model.ErrTaskNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrTaskNotFound))
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, taskSuccessResponse(t))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}

	}
}

func getTaskByText(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getTaskByTextRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		tasks, err := a.GetTaskByText(c, req.Text)

		switch {
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, tasksSuccessResponse(tasks))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

func updateTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		var req updateTaskRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		t, err := a.UpdateTask(c, id, model.TodoTask{
			Title:       req.Title,
			Description: req.Description,
			PlanningDate: model.Date{
				Day:   req.PlanningDate.Day,
				Month: time.Month(req.PlanningDate.Month),
				Year:  req.PlanningDate.Year,
			},
			Status: req.Status,
		})

		switch {
		case errors.Is(err, model.ErrTaskNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrTaskNotFound))
		case errors.Is(err, model.ErrInvalidTask):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidTask))
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, taskSuccessResponse(t))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

func deleteTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		err = a.DeleteTask(c, id)

		switch {
		case errors.Is(err, model.ErrTaskNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrTaskNotFound))
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, deleteSuccessResponse())
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

func getTasksByStatus(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getTasksByStatusRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		tasks, err := a.GetTasksByStatus(c, req.Status, req.Offset, req.Limit)

		switch {
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, tasksSuccessResponse(tasks))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

func getTasksByDate(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getTasksByDateRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		tasks, err := a.GetTasksByDate(c, model.Date{
			Day:   req.PlanningDate.Day,
			Month: time.Month(req.PlanningDate.Month),
			Year:  req.PlanningDate.Year,
		}, req.Offset, req.Limit)

		switch {
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, tasksSuccessResponse(tasks))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

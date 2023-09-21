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

// @Summary		Добавление новой задачи
// @Description	Возвращает добавленную задачу с её id в postgres
// @Produce		json
// @Param			input body addTaskRequest true "Новая задача в JSON"
// @Success		200		{object}	taskResponse	"Успешное добавление"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 400 {object} taskResponse "Неверный формат входных данных"
// @Router			/task [post]
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
				Year:  req.PlanningDate.Year,
				Month: time.Month(req.PlanningDate.Month),
				Day:   req.PlanningDate.Day,
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

// @Summary		Поиск задачи по её id в postgres
// @Description	Возвращает задачу с заданным id
// @Produce		json
// @Param			input body getTaskByTextRequest true "id задачи"
// @Success		200		{object}	taskResponse	"Успешное получение"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 400 {object} taskResponse "Неверный формат входных данных"
// @Failure 404 {object} taskResponse "Задача с заданным id не найдена"
// @Router			/task/{id} [get]
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

// @Summary		Поиск задачи по тексту заголовка или описания
// @Description	Возвращает задачу с вхождением данной строки в заголовке или описании
// @Produce		json
// @Param			id path int true "Текст в JSON"
// @Success		200		{object}	tasksResponse	"Успешное получение задач"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 400 {object} taskResponse "Неверный формат входных данных"
// @Router			/task [get]
func getTaskByText(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getTaskByTextRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		tasks, err := a.GetTaskByText(c, req.Text)

		switch {
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, tasksSuccessResponse(tasks))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

// @Summary		Обновление полей задачи по её id в postgres
// @Description	Возвращает задачу с заданным id и изменёнными полями
// @Produce		json
// @Param			input body updateTaskRequest true "Новые поля задачи"
// @Param id path int true "id изменяемой задачи"
// @Success		200		{object}	taskResponse	"Успешное обновление"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 400 {object} taskResponse "Неверный формат входных данных"
// @Failure 404 {object} taskResponse "Задача с заданным id не найдена"
// @Router			/task/{id} [put]
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
				Year:  req.PlanningDate.Year,
				Month: time.Month(req.PlanningDate.Month),
				Day:   req.PlanningDate.Day,
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

// @Summary		Удаление задачи по её id в postgres
// @Description	Удаляет задачу с заданным id
// @Produce		json
// @Param 		id path int true "id удаляемой задачи"
// @Success		200		{object}	taskResponse	"Успешное удаление"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 	400 {object} taskResponse "Неверный формат входных данных"
// @Failure 	404 {object} taskResponse "Задача с заданным id не найдена"
// @Router		/task/{id} [delete]
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

// @Summary		Получение списка задач с фильтром по статусу и пагинацией
// @Description	Возвращает список задач
// @Produce		json
// @Param		input body getTasksByStatusRequest true "Статус и пагинация"
// @Success		200		{object}	taskResponse	"Успешное получение задач"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 	400 {object} taskResponse "Неверный формат входных данных"
// @Router		/task/by_status [get]
func getTasksByStatus(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getTasksByStatusRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		tasks, err := a.GetTasksByStatus(c, req.Status, req.Offset, req.Limit)

		switch {
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrTaskRepo):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrTaskRepo))
		case err == nil:
			c.JSON(http.StatusOK, tasksSuccessResponse(tasks))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrUnknown))
		}
	}
}

// @Summary		Получение списка задач с фильтром по дате и статусу
// @Description	Возвращает список задач
// @Produce		json
// @Param		input body getTasksByDateAndStatusRequest true "Дата и статус"
// @Success		200		{object}	taskResponse	"Успешное получение задач"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure 	400 {object} taskResponse "Неверный формат входных данных"
// @Router		/task/by_date [get]
func getTasksByDateAndStatus(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getTasksByDateAndStatusRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		}

		tasks, err := a.GetTasksByDateAndStatus(c, model.Date{
			Year:  req.PlanningDate.Year,
			Month: time.Month(req.PlanningDate.Month),
			Day:   req.PlanningDate.Day,
		}, req.Status)

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

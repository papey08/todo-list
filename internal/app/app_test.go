package app

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	"todo-list/internal/app/mocks"
	"todo-list/internal/model"
)

type appTestSuite struct {
	suite.Suite
	taskRepo *mocks.TaskRepo
	a        App
}

func (s *appTestSuite) SetupSuite() {
	s.taskRepo = new(mocks.TaskRepo)
	s.a = New(s.taskRepo)
}

type addTaskMock struct {
	givenTask  model.TodoTask
	returnTask model.TodoTask
	returnErr  error
}

type addTaskTest struct {
	description  string
	givenTask    model.TodoTask
	expectedTask model.TodoTask
	expectedErr  error
}

func (s *appTestSuite) TestAddTask() {
	addTestMocks := []addTaskMock{
		{
			givenTask: model.TodoTask{
				Title:       "Successfully added task",
				Description: "Description of successfully added task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnTask: model.TodoTask{
				Id:          1,
				Title:       "Successfully added task",
				Description: "Description of successfully added task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnErr: nil,
		},
		{
			givenTask: model.TodoTask{
				Title:       "Not added task",
				Description: "Description of not added task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnTask: model.TodoTask{},
			returnErr:  model.ErrTaskRepo,
		},
	}

	tests := []addTaskTest{
		{
			description: "test of successful adding of task",
			givenTask: model.TodoTask{
				Title:       "Successfully added task",
				Description: "Description of successfully added task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedTask: model.TodoTask{
				Id:          1,
				Title:       "Successfully added task",
				Description: "Description of successfully added task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedErr: nil,
		},
		{
			description: "test of not adding of task",
			givenTask: model.TodoTask{
				Title:       "Not added task",
				Description: "Description of not added task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedTask: model.TodoTask{},
			expectedErr:  model.ErrTaskRepo,
		},
		{
			description: "test of adding of invalid task",
			givenTask: model.TodoTask{
				Title:       "",
				Description: "Description of invalid task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedTask: model.TodoTask{},
			expectedErr:  model.ErrInvalidTask,
		},
	}

	for _, m := range addTestMocks {
		s.taskRepo.On("AddTask", mock.Anything, m.givenTask).Return(m.returnTask, m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			task, err := s.a.AddTask(ctx, test.givenTask)
			assert.Equal(t, test.expectedTask, task)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

type getTaskByIdMock struct {
	givenId    int
	returnTask model.TodoTask
	returnErr  error
}

type getTaskByIdTest struct {
	description  string
	givenId      int
	expectedTask model.TodoTask
	expectedErr  error
}

func (s *appTestSuite) TestGetTaskById() {
	getTaskByIdMocks := []getTaskByIdMock{
		{
			givenId: 1,
			returnTask: model.TodoTask{
				Id:          1,
				Title:       "title",
				Description: "description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnErr: nil,
		},
		{
			givenId:    2,
			returnTask: model.TodoTask{},
			returnErr:  model.ErrTaskRepo,
		},
		{
			givenId:    46447,
			returnTask: model.TodoTask{},
			returnErr:  model.ErrTaskNotFound,
		},
	}

	tests := []getTaskByIdTest{
		{
			description: "test of successful getting of the task by id",
			givenId:     1,
			expectedTask: model.TodoTask{
				Id:          1,
				Title:       "title",
				Description: "description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedErr: nil,
		},
		{
			description:  "test of error with database",
			givenId:      2,
			expectedTask: model.TodoTask{},
			expectedErr:  model.ErrTaskRepo,
		},
		{
			description:  "test of non existing id",
			givenId:      46447,
			expectedTask: model.TodoTask{},
			expectedErr:  model.ErrTaskNotFound,
		},
	}

	for _, m := range getTaskByIdMocks {
		s.taskRepo.On("GetTaskById", mock.Anything, m.givenId).Return(m.returnTask, m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			task, err := s.a.GetTaskById(ctx, test.givenId)
			assert.Equal(t, test.expectedTask, task)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

type getTaskByTextMock struct {
	givenText   string
	returnTasks []model.TodoTask
	returnErr   error
}

type getTaskByTextTest struct {
	description   string
	givenText     string
	expectedTasks []model.TodoTask
	expectedErr   error
}

func (s *appTestSuite) TestGetTaskByText() {
	getTasksByTextMocks := []getTaskByTextMock{
		{
			givenText: "abc",
			returnTasks: []model.TodoTask{
				{
					Id:          1,
					Title:       "title",
					Description: "abc",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          2,
					Title:       "other title",
					Description: "abcdef",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          3,
					Title:       "xyzABC",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
			},
			returnErr: nil,
		},
		{
			givenText:   "non existing text",
			returnTasks: []model.TodoTask{},
			returnErr:   nil,
		},
		{
			givenText:   "text calling error in database",
			returnTasks: nil,
			returnErr:   model.ErrTaskRepo,
		},
	}

	tests := []getTaskByTextTest{
		{
			description: "test of finding ads with given substring in title or description",
			givenText:   "abc",
			expectedTasks: []model.TodoTask{
				{
					Id:          1,
					Title:       "title",
					Description: "abc",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          2,
					Title:       "other title",
					Description: "abcdef",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          3,
					Title:       "xyzABC",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
			},
			expectedErr: nil,
		},
		{
			description:   "test of making query with non existing text",
			givenText:     "non existing text",
			expectedTasks: []model.TodoTask{},
			expectedErr:   nil,
		},
		{
			description:   "test of occurring error in database",
			givenText:     "text calling error in database",
			expectedTasks: nil,
			expectedErr:   model.ErrTaskRepo,
		},
		{
			description:   "test of finding an empty string",
			givenText:     "",
			expectedTasks: nil,
			expectedErr:   model.ErrInvalidInput,
		},
	}

	for _, m := range getTasksByTextMocks {
		s.taskRepo.On("GetTaskByText", mock.Anything, m.givenText).Return(m.returnTasks, m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			task, err := s.a.GetTaskByText(ctx, test.givenText)
			assert.Equal(t, test.expectedTasks, task)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

type updateTaskMock struct {
	givenId    int
	givenTask  model.TodoTask
	returnTask model.TodoTask
	returnErr  error
}

type updateTaskTest struct {
	description  string
	givenId      int
	givenTask    model.TodoTask
	expectedTask model.TodoTask
	expectedErr  error
}

func (s *appTestSuite) TestUpdateTask() {
	updateTaskMocks := []updateTaskMock{
		{
			givenId: 1,
			givenTask: model.TodoTask{
				Title:       "other title",
				Description: "other description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnTask: model.TodoTask{
				Id:          1,
				Title:       "other title",
				Description: "other description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnErr: nil,
		},
		{
			givenId: 46447,
			givenTask: model.TodoTask{
				Title:       "other title",
				Description: "other description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			returnTask: model.TodoTask{},
			returnErr:  model.ErrTaskNotFound,
		},
	}

	tests := []updateTaskTest{
		{
			description: "test of successful updating of the task",
			givenId:     1,
			givenTask: model.TodoTask{
				Title:       "other title",
				Description: "other description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedTask: model.TodoTask{
				Id:          1,
				Title:       "other title",
				Description: "other description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedErr: nil,
		},
		{
			description: "test of updating non existing task",
			givenId:     46447,
			givenTask: model.TodoTask{
				Title:       "other title",
				Description: "other description",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedTask: model.TodoTask{},
			expectedErr:  model.ErrTaskNotFound,
		},
		{
			description: "test of updating task to invalid",
			givenId:     2,
			givenTask: model.TodoTask{
				Title:       "",
				Description: "description of invalid task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: time.January,
					Day:   1,
				},
				Status: false,
			},
			expectedTask: model.TodoTask{},
			expectedErr:  model.ErrInvalidTask,
		},
	}

	for _, m := range updateTaskMocks {
		s.taskRepo.On("UpdateTask", mock.Anything, m.givenId, m.givenTask).Return(m.returnTask, m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			task, err := s.a.UpdateTask(ctx, test.givenId, test.givenTask)
			assert.Equal(t, test.expectedTask, task)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

type deleteTaskMock struct {
	givenId   int
	returnErr error
}

type deleteTaskTest struct {
	description string
	givenId     int
	expectedErr error
}

func (s *appTestSuite) TestDeleteTask() {
	deleteTaskMocks := []deleteTaskMock{
		{
			givenId:   1,
			returnErr: nil,
		},
		{
			givenId:   2,
			returnErr: model.ErrTaskRepo,
		},
		{
			givenId:   46447,
			returnErr: model.ErrTaskNotFound,
		},
	}

	tests := []deleteTaskTest{
		{
			description: "test of successful deleting of the task",
			givenId:     1,
			expectedErr: nil,
		},
		{
			description: "test of occurring error in the database",
			givenId:     2,
			expectedErr: model.ErrTaskRepo,
		},
		{
			description: "test of deleting non existing task",
			givenId:     46447,
			expectedErr: model.ErrTaskNotFound,
		},
	}

	for _, m := range deleteTaskMocks {
		s.taskRepo.On("DeleteTask", mock.Anything, m.givenId).Return(m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			err := s.a.DeleteTask(ctx, test.givenId)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

type getTasksByStatusMock struct {
	givenStatus bool
	givenLimit  int
	givenOffset int
	returnTasks []model.TodoTask
	returnErr   error
}

type getTasksByStatusTest struct {
	description   string
	givenStatus   bool
	givenLimit    int
	givenOffset   int
	expectedTasks []model.TodoTask
	expectedErr   error
}

func (s *appTestSuite) TestGetTasksByStatus() {
	getTaskByStatusMocks := []getTasksByStatusMock{
		{
			givenStatus: true,
			givenLimit:  1,
			givenOffset: 0,
			returnTasks: []model.TodoTask{
				{
					Id:          1,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: true,
				},
			},
			returnErr: nil,
		},
		{
			givenStatus: false,
			givenLimit:  2,
			givenOffset: 1,
			returnTasks: []model.TodoTask{
				{
					Id:          2,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          3,
					Title:       "other title",
					Description: "other description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
			},
			returnErr: nil,
		},
		{
			givenStatus: false,
			givenLimit:  0,
			givenOffset: 0,
			returnTasks: []model.TodoTask{},
			returnErr:   nil,
		},
		{
			givenStatus: true,
			givenLimit:  0,
			givenOffset: 0,
			returnTasks: nil,
			returnErr:   model.ErrTaskRepo,
		},
	}

	tests := []getTasksByStatusTest{
		{
			description: "test of successful getting slice of done tasks",
			givenStatus: true,
			givenLimit:  1,
			givenOffset: 0,
			expectedTasks: []model.TodoTask{
				{
					Id:          1,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: true,
				},
			},
			expectedErr: nil,
		},
		{
			description: "test of successful getting slice of undone tasks",
			givenStatus: false,
			givenLimit:  2,
			givenOffset: 1,
			expectedTasks: []model.TodoTask{
				{
					Id:          2,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          3,
					Title:       "other title",
					Description: "other description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
			},
			expectedErr: nil,
		},
		{
			description:   "test of getting an empty slice of tasks",
			givenStatus:   false,
			givenLimit:    0,
			givenOffset:   0,
			expectedTasks: []model.TodoTask{},
			expectedErr:   nil,
		},
		{
			description:   "test of occurring an error in the database",
			givenStatus:   true,
			givenLimit:    0,
			givenOffset:   0,
			expectedTasks: nil,
			expectedErr:   model.ErrTaskRepo,
		},
		{
			description:   "test of getting a slice of tasks with invalid limit and offset",
			givenStatus:   true,
			givenLimit:    -1,
			givenOffset:   -1,
			expectedTasks: nil,
			expectedErr:   model.ErrInvalidInput,
		},
	}

	for _, m := range getTaskByStatusMocks {
		s.taskRepo.On("GetTasksByStatus", mock.Anything, m.givenStatus, m.givenLimit, m.givenOffset).Return(m.returnTasks, m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			tasks, err := s.a.GetTasksByStatus(ctx, test.givenStatus, test.givenLimit, test.givenOffset)
			assert.Equal(t, test.expectedTasks, tasks)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

type getTasksByDateAndStatusMock struct {
	givenDate   model.Date
	givenStatus bool
	returnTasks []model.TodoTask
	returnErr   error
}

type getTasksByDateAndStatusTest struct {
	description   string
	givenDate     model.Date
	givenStatus   bool
	expectedTasks []model.TodoTask
	expectedErr   error
}

func (s *appTestSuite) TestGetTasksByDateAndStatus() {
	getTaskByDateAndStatusMocks := []getTasksByDateAndStatusMock{
		{
			givenDate: model.Date{
				Year:  2024,
				Month: time.January,
				Day:   1,
			},
			givenStatus: true,
			returnTasks: []model.TodoTask{
				{
					Id:          1,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: true,
				},
			},
			returnErr: nil,
		},
		{
			givenDate: model.Date{
				Year:  2024,
				Month: time.January,
				Day:   1,
			},
			givenStatus: false,
			returnTasks: []model.TodoTask{
				{
					Id:          2,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          3,
					Title:       "other title",
					Description: "other description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
			},
			returnErr: nil,
		},
		{
			givenDate: model.Date{
				Year:  2025,
				Month: time.January,
				Day:   1,
			},
			givenStatus: true,
			returnTasks: nil,
			returnErr:   model.ErrTaskRepo,
		},
	}

	tests := []getTasksByDateAndStatusTest{
		{
			description: "test of successful getting slice of done tasks by date",
			givenDate: model.Date{
				Year:  2024,
				Month: time.January,
				Day:   1,
			},
			givenStatus: true,
			expectedTasks: []model.TodoTask{
				{
					Id:          1,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: true,
				},
			},
			expectedErr: nil,
		},
		{
			description: "test of successful getting slice of undone tasks by date",
			givenDate: model.Date{
				Year:  2024,
				Month: time.January,
				Day:   1,
			},
			givenStatus: false,
			expectedTasks: []model.TodoTask{
				{
					Id:          2,
					Title:       "title",
					Description: "description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
				{
					Id:          3,
					Title:       "other title",
					Description: "other description",
					PlanningDate: model.Date{
						Year:  2024,
						Month: time.January,
						Day:   1,
					},
					Status: false,
				},
			},
			expectedErr: nil,
		},
		{
			description: "test of occurring an error in the database",
			givenDate: model.Date{
				Year:  2025,
				Month: time.January,
				Day:   1,
			},
			givenStatus:   true,
			expectedTasks: nil,
			expectedErr:   model.ErrTaskRepo,
		},
		{
			description: "test of getting slice of tasks with invalid date",
			givenDate: model.Date{
				Year:  2024,
				Month: time.February,
				Day:   30,
			},
			givenStatus:   true,
			expectedTasks: nil,
			expectedErr:   model.ErrInvalidInput,
		},
	}

	for _, m := range getTaskByDateAndStatusMocks {
		s.taskRepo.On("GetTasksByDateAndStatus", mock.Anything, m.givenDate, m.givenStatus).Return(m.returnTasks, m.returnErr).Once()
	}

	ctx := context.Background()

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			tasks, err := s.a.GetTasksByDateAndStatus(ctx, test.givenDate, test.givenStatus)
			assert.Equal(t, test.expectedTasks, tasks)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(appTestSuite))
}

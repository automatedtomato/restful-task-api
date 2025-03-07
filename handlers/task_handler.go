package handlers

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/automatedtomato/restful-task-api/models"
	"github.com/gin-gonic/gin"
)

var (
	tasks = make(map[string]models.Task)
	mutex = &sync.Mutex{} // protect data from parallel access
)

const notFound = "Task not found"

// Initiate test data
func init() {
	task := models.Task{
		ID:          "1",
		Title:       "Sample Task",
		Description: "This is a sample description",
		Status:      "Not Complete",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks[task.ID] = task
}

// GetTasks godoc
// @Summary 	get task list
// @Description return list of all tasks
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Success 	200 {array} models.Task
// @Router 		/tasks [get]
func GetTasks(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	c.JSON(http.StatusOK, taskList)
}

// GetTask godoc
// @Summary	get particular task
// @Description	get task by ID
// @Tags		tasks
// @Accept		json
// @Produce		json
// @Param		id    path      string  true  "Task ID"
// @Success		200   {object}  models.Task
// @Failure		404   {object}  map[string]string
// @Router		/tasks/{id} [get]
func GetTask(c *gin.Context) {
	id := c.Param("id")

	mutex.Lock()
	defer mutex.Unlock()

	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": notFound})
		return
	}

	c.JSON(http.StatusOK, task)
}
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate ID
	newTask.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	if newTask.Status == "" {
		newTask.Status = "Not Complete"
	}

	mutex.Lock()
	tasks[newTask.ID] = newTask
	mutex.Unlock()

	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask godoc
// @Summary     update task
// @Description update task by ID
// @Tags        tasks
// @Accept      json
// @Produce     json
// @Param       id    path      string       true    "Task ID"
// @Param       task  body      models.Task  true    "Task info"
// @Successes   200   {object}  models.Task  true
// @Failure     404   {object}  map[string]string
// @Failure     400   {object}  map[string]string
// @Router      /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updateData models.Task
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mutex.Lock()
	defer mutex.Unlock()

	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": notFound})
		return
	}

	// update fields
	task.Title = updateData.Title
	task.Description = updateData.Description
	if updateData.Status != "" {
		task.Status = updateData.Status
	}
	task.UpdatedAt = time.Now()

	// update task
	tasks[id] = task

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := tasks[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": notFound})
		return
	}

	delete(tasks, id)

	c.Status(http.StatusNoContent)
}

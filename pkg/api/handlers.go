package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"workspace/task_manager/pkg/task"
)

type TaskSpec struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
}

func (s *Server) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// handle get all tasks
		tasks := s.taskManager.GetAllTasks()
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		// handle create new task
		var t task.Task
		json.NewDecoder(r.Body).Decode(&t)
		s.taskManager.AddTask(t)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	switch r.Method {
	case http.MethodGet:
		// handle get a task
		t, err := s.taskManager.GetTask(id)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(t)
	case http.MethodPut:
		// handle update a task
		var t task.Task
		json.NewDecoder(r.Body).Decode(&t)
		s.taskManager.UpdateTask(id, t)
	case http.MethodDelete:
		// handle delete a task
		s.taskManager.DeleteTask(id)
	case http.MethodPost:
		// handle start/stop a task
		action := r.URL.Query().Get("action")
		if action == "start" {
			s.taskManager.StartTask(id)
		} else if action == "stop" {
			s.taskManager.StopTask(id)
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

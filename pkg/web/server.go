package web

import (
	"html/template"
	"net/http"
	"workspace/task_manager/pkg/task"
)

type Server struct {
	taskManager *task.Manager
}

func NewServer(taskManager *task.Manager) *Server {
	return &Server{
		taskManager: taskManager,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tasks := s.taskManager.GetAllTasks()
		tmpl, _ := template.ParseFiles("pkg/web/index.html")
		tmpl.Execute(w, tasks)
	case "/stop":
		id := r.URL.Query().Get("id")
		s.taskManager.StopTask(id)
	case "/delete":
		id := r.URL.Query().Get("id")
		s.taskManager.DeleteTask(id)
	case "/add":
		// handle add task
	}
}

package api

import (
	"net/http"
	"workspace/task_manager/pkg/task"
)

type Server struct {
	httpServer  *http.Server
	taskManager *task.Manager
}

func NewServer(port string, taskManager *task.Manager) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: ":" + port,
		},
		taskManager: taskManager,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/tasks", s.handleTasks)
	http.HandleFunc("/tasks/", s.handleTask)
	return s.httpServer.ListenAndServe()
}

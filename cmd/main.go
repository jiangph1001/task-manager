package main

import (
	"log"
	"net/http"
	"workspace/task_manager/pkg/task"
	"workspace/task_manager/pkg/web"
)

func main() {
	// Initialize the task manager
	taskManager, err := task.NewTaskManager("tasks.db")
	if err != nil {
		log.Fatalf("Failed to create task manager: %v", err)
	}
	defer taskManager.Close()

	// Initialize the web server
	server := web.NewServer(taskManager)

	// Start the web server
	log.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

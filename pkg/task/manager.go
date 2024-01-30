package task

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

// Manager is responsible for managing tasks
type Manager struct {
	tasks map[string]Task
	db    *leveldb.DB
	mu    sync.Mutex
}

// NewTaskManager creates a new Manager
func NewTaskManager(dbPath string) (*Manager, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}

	return &Manager{
		tasks: make(map[string]Task),
		db:    db,
	}, nil
}

// AddTask adds a new task to the Manager
func (tm *Manager) AddTask(t Task) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// uuid
	id := uuid.New().String()

	tm.tasks[id] = t

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return tm.db.Put([]byte(id), data, nil)
}

// DeleteTask deletes a task from the Manager
func (tm *Manager) DeleteTask(id string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	delete(tm.tasks, id)

	return tm.db.Delete([]byte(id), nil)
}

// StartTask starts a task in the Manager
func (tm *Manager) StartTask(id string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	task.SetStatus(Running)
	go func(t Task) {
		err := t.Execute()
		if err != nil {
			// handle error
		}
	}(task)

	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return tm.db.Put([]byte(id), data, nil)
}

// StopTask stops a task in the Manager
func (tm *Manager) StopTask(id string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	task.SetStatus(Stopped)

	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return tm.db.Put([]byte(id), data, nil)
}

// GetTask gets a task from the Manager
func (tm *Manager) GetTask(id string) (Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if exists {
		return task, nil
	}

	data, err := tm.db.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}

	var t Task
	err = json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}

	tm.tasks[id] = t

	return t, nil
}

func (tm *Manager) GetAllTasks() []Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tasks := make([]Task, 0, len(tm.tasks))
	for _, t := range tm.tasks {
		tasks = append(tasks, t)
	}

	return tasks
}

func (tm *Manager) UpdateTask(id string, t Task) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	_, exists := tm.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	tm.tasks[id] = t

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return tm.db.Put([]byte(id), data, nil)
}

func (tm *Manager) Close() {
	tm.db.Close()
}

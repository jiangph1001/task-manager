package taskimpl

import (
	"os/exec"

	"workspace/task_manager/pkg/task"
)

type TranscodeTask struct {
	id     string
	status task.TaskStatus
	cmd    *exec.Cmd
}

func NewTranscodeTask(id string, inputFile string, outputFile string) *TranscodeTask {
	cmd := exec.Command("ffmpeg", "-i", inputFile, outputFile)
	return &TranscodeTask{
		id:     id,
		status: task.Running,
		cmd:    cmd,
	}
}

func (t *TranscodeTask) ID() string {
	return t.id
}

func (t *TranscodeTask) Status() task.TaskStatus {
	return t.status
}

func (t *TranscodeTask) SetStatus(status task.TaskStatus) {
	t.status = status
}

func (t *TranscodeTask) Execute() error {
	t.status = task.Running
	err := t.cmd.Start()
	if err != nil {
		t.status = task.Failed
		return err
	}
	err = t.cmd.Wait()
	if err != nil {
		t.status = task.Failed
		return err
	}
	t.status = task.Completed
	return nil
}

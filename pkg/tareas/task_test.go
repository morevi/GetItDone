
package tareas

import (
  "testing"
  "time"
)

func TestNewTask(t *testing.T) {
  done := false
  desc := "test string"
  due  := time.Now()

  task := new(Task)
  task.New(done, desc, due)

  if task.Done != done || task.Content != desc || task.Due != due {
    t.Error("Task.New() did not pass")
  }
}

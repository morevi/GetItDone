
package tareas

import (
  "testing"
  "time"
)

func TestNewTask(t *testing.T) {
  t.Log("Test set task creation")
  done := false
  desc := "test string"
  due  := time.Now()

  task := new(Task)
  task.New(done, desc, due)

  if task.Done != done || task.Content != desc || task.Due != due {
    t.Error("Task.New() did not pass")
  }
}

func TestToogleCheck(t *testing.T) {
  t.Log("Test set task checking")
  task := new(Task)
  task.New(false, "", time.Now())

  task.ToogleCheck()

  if task.Done != true {
    t.Error("Task.ToogleCheck() did not pass")
  }
}

func TestSetContent(t *testing.T) {
  t.Log("Test set task content")
  task := new(Task)
  task.New(false, "", time.Now())

  task.SetContent("new content")

  if task.Content != "new content" {
    t.Error("Task.SetContent() did not pass")
  }
}

func TestSetDue(t*testing.T) {
  t.Log("Test set task due")
  task := new(Task)
  task.New(false, "", time.Now())

  testDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
  task.SetDue(testDate)

  if task.Due != testDate {
    t.Error("Task.SetDue() did not pass")
  }
}

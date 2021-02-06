
package tareas

import (
  "testing"
  "time"
  "reflect"
)

func TestNewTask(t *testing.T) {
  t.Log("Creates a new task")
  done := false
  desc := "test string"
  due  := time.Now()

  var want Task
  want.Done = done
  want.Content = desc
  want.Due = due

  var got Task
  got.New(done, desc, due)

  if !reflect.DeepEqual(got, want) {
    t.Errorf("\ngot : %+v\nwant: %+v\n", got, want)
  }
}

func TestToogleCheck(t *testing.T) {
  t.Log("Toggles completed field")
  task := new(Task)
  task.New(false, "", time.Now())

  task.ToogleCheck()

  if task.Done != true {
    t.Errorf("\ngot : %+v\nwant: %+v\n", task.Done, true)
  }
}

func TestSetContent(t *testing.T) {
  t.Log("Set content field in task")
  task := new(Task)
  task.New(false, "", time.Now())

  task.SetContent("new content")

  if task.Content != "new content" {
    t.Errorf("\ngot : %+v\nwant: %+v\n", task.Content, "new content")
  }
}

func TestSetDue(t*testing.T) {
  t.Log("Set due field in task")
  task := new(Task)
  task.New(false, "", time.Now())

  want := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
  task.SetDue(want)

  if task.Due != 3 {
    t.Errorf("\ngot : %+v\nwant: %+v\n", task.Due, want)
  }
}

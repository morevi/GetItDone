
package tareas

import (
  "testing"
  "time"
  "reflect"
)

func initializeProject() Project{
  var p Project
  tags := []string{"proyecto", "iv", "test"}
  desc := "test description"

  var t1 Task
  t1.New(false, "test task 1", time.Now().Add(24*time.Hour))
  var t2 Task
  t2.New(false, "test task 2", time.Now().Add(48*time.Hour))

  todos := []Task{t1, t2}

  p.Tags = tags
  p.Description = desc
  p.Items = todos

  return p
}

func TestProjectNew(t *testing.T) {
  t.Log("Test Project contructor")
  want := initializeProject()

  var t1 Task
  t1.New(false, "test task 1", time.Now().Add(24*time.Hour))
  var t2 Task
  t2.New(false, "test task 2", time.Now().Add(48*time.Hour))

  var p Project
  p.New([]string{"proyecto", "iv", "test"}, "test description", []Task{t1, t2})

  if reflect.DeepEqual(p,want) {
    t.Errorf("Test for Project.New() did not pass")
  }
}

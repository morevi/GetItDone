
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

func eq(a []string, b []string) bool {
    if (a == nil) != (b == nil) {
        return false
    }
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
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

func TestSetDescription(t *testing.T) {
  t.Log("Test project set description")
  want := "new project description"

  p := initializeProject()
  p.SetDescription(want)

  if p.Description != want {
    t.Errorf("received %s", p.Description)
  }
}

func TetSetTags(t *testing.T) {
  t.Log("Test project test tags")
  p := initializeProject()

  p.SetTags("a b c")
  want := []string{"a", "b", "c"}

  if eq(p.Tags, want) {
    t.Errorf(" setTags failed ")
  }
}

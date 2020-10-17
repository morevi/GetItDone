
package main

import (
  "fmt"
  "github.com/morevi/GetItDone/pkg/tareas"
)

func main() {
  t := tareas.Task{false, "Subir commits a github"}

  labels := []string{"Uni","IV","important"}
  items := []tareas.Task{t}
  l := tareas.TaskList{"Primer hito", labels, items}

  fmt.Println(l.Name, l.Label, l.Items)
}

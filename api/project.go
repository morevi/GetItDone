package handler

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "github.com/morevi/GetItDone/pkg/tareas"
)

func getFromDB() tareas.Dashboard {
  // El plan es que se cargue el dashboard de usuario desde una db,
  // pero por ahora no está implementado, asi que se generan datos
  // de ejemplo
  var d tareas.Dashboard
  d.New(91)

  var t tareas.Task
  t.New(true, "Hito 5", "2020-02-17")

  // Proyecto de ejemplo 1
  var p tareas.Project
  p.New([]string{"iv","serverless","despliegue"},
  "Hito 5 de la asignatura IV",
  []tareas.Task{t})
  d.Add(p)

  // Proyecto de ejemplo 2
  var p2 tareas.Project
  p2.New([]string{"iv"},
  "ejemplo",
  []tareas.Task{t})
  d.Add(p2)

  return d
}

func Handler(w http.ResponseWriter, r *http.Request) {
  d := getFromDB()
  var resultado []byte

  switch r.Method {
  // Get every or some projects
  case "GET":
    tags, ok := r.URL.Query()["tags"]
    if !ok {
      output, err := json.Marshal(d.GetAll())
      if err != nil {
        http.Error(w, err.Error(), 500)
      }
      resultado = output
    } else {
      output, err := json.Marshal(d.SearchByTags(tags))
      if err != nil {
        http.Error(w, err.Error(), 500)
      }
      resultado = output
    }

  case "POST":
    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    if err != nil {
	     http.Error(w, err.Error(), 500)
		    return
	   }

    var p tareas.Project
    err = json.Unmarshal(body, &p)
    if err != nil {
      http.Error(w, err.Error(), 500)
      return
    }
    d.Add(p)
    output, _ := json.Marshal(201)
    resultado = output
  }

  w.Header().Set("content-type","application/json")
  w.Write(resultado)
}


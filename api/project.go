package handler

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/morevi/GetItDone/pkg/tareas"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    var d tareas.Dashboard

    // get from the DB
    d.New(0)
    //
    res, _ := json.Marshal(d)
    fmt.Fprintf(w, string(res))
}


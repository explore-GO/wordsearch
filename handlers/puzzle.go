package handlers

import (
    "net/http"

    "github.com/gorilla/mux"
)

func (env *Environment) createPuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseForm()
    if err != nil {
        return StatusError{500, err}
    }

    body := r.PostFormValue("body")
    if len(body) == 0 {
        return StatusError{400, nil}
    }

    url, err := env.app.CreatePuzzle(body, 0)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}

func (env *Environment) updatePuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    url := vars["url"]

    err := r.ParseForm()
    if err != nil {
        return StatusError{500, err}
    }

    body := r.PostFormValue("body")
    if len(body) == 0 {
        return StatusError{400, nil}
    }

    err = env.app.UpdatePuzzle(url, body)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}

func (env *Environment) clonePuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    viewUrl := vars["view_url"]

    url, err := env.app.ClonePuzzle(viewUrl)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}

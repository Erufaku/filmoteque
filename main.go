package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

const port = ":8080"

type Store struct {
	Conn string
}

type ActorInfo struct {
	Name  string
	Male  string
	Birth time.Time
}

type FilmInfo struct {
	Name        string
	Description string
	Date        time.Time
	Rate        int
	Actors      []ActorInfo
}

// Tables
//1. Actors //uniq
//2. Films //uniq
//3. Actors_films //relat

// 1. act1 film1
// 2. act1 film2
// 3. act2 film3
//...

// sort films
//SELECT * FROM films ORDER BY rate

// find films
//SELECT * FROM films WHERE name LIKE '%sky%'

func add(a, b int) int {
	return a + b
}

func Test_add(t *testing.T) {
	result := add(1, 1)
	result = add(3, 1)
	result = result

	// want
	// actual
	// if want != actual

	//t.Errorf("")
}

func (s Store) Actors(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

	}

	var actor ActorInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &actor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// validate
	if actor.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate other fiels ...

	// logic
	err = addActor(actor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s Store) AdminActors(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

	}

	// auth for admin routes
	r.BasicAuth()

	// if OK {
	//...
	//} else {
	//w.WriteHeader(401)
	//return
	//}
}

func addActor(req ActorInfo) error {
	// SELECT COUNT(*) FROM actors WHERE name = $1 AND male...

	// INSERT INTO actors (name, male, birth) VALUES ($1, $2, $3)

	return nil
}

func updateActor(req ActorInfo) error {
	// SELECT COUNT(*) FROM actors WHERE name = $1 AND male...

	// INSERT INTO actors (name, male, birth) VALUES ($1, $2, $3)

	return nil
}

func main() {

	store := Store{}

	http.HandleFunc("/actor", store.Actors)
	http.HandleFunc("/actor", store.Actors)

	http.HandleFunc("/admin/actor", store.AdminActors)

	http.HandleFunc("/films", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(port, nil)
}

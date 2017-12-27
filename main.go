package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"text/template"
	"math/rand"
)

const viewSize = 16
var alive = true
var hero = "&Omega;"
var cross = "&#x2626;"
const water = "~"
const ground = "."
var terrains = [2]string {"~", "."}
var view [][]string
var world [][]string

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)

	if r.URL.Path != "/" {
		handlerError(w, r, http.StatusNotFound)
		return
	}

	t, _ := template.ParseFiles("index.html")
	x, _ := strconv.Atoi(r.Form.Get("x"))
	y, _ := strconv.Atoi(r.Form.Get("y"))

	if !alive {
		fmt.Println("redirect")
		http.Redirect(w, r, "/", 302)
		alive = true;
		return
	}

	alive = placeHero(x, y)
	fmt.Println("alive", alive)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, struct{View [][]string}{view})
}

func handlerFavicon(w http.ResponseWriter, r *http.Request) {
	handlerError(w, r, http.StatusNotFound)
}

func handlerError(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Not Found")
	}
}

func placeHero(x, y int) (bool) {
	for i := 0; i < viewSize; i++ {
		view[i] = make([]string, len(world[i]))
		copy(view[i], world[i]);
	}

	x = x%viewSize
	y = y%viewSize

	if x < 0 {
		x = viewSize + x
	}

	if y < 0 {
		y = viewSize + y
	}

	place := &view[y][x]

	if *place == water {
		*place = cross
		return false
	}

	*place = hero
	return true
}

func main() {
	for i := 0; i < viewSize; i++ {
		world = append(world, []string{})

		for j := 0; j < viewSize; j++ {
			world[i] = append(world[i], terrains[rand.Intn(len(terrains))])
		}
	}

	world[0][0] = ground

	view = make([][]string, len(world))
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/favicon.ico", handlerFavicon)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

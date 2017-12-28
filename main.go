package main

import (
	"fmt"
	"net/http"
// 	"strconv"
	"log"
	"text/template"
	"math/rand"
	"container/list"
)

const viewSize = 16
const water = "~"
const ground = "."

type Hero struct {
	Alive bool
	X int64
	Y int64
}

func NewHero() (Hero) {
	hero := Hero{}
	hero.Init();
	return hero
}

func (hero *Hero) Init() {
	hero.Alive = true
	hero.X = 0
	hero.Y = 0	
}

func (hero *Hero) Symbol() (string) {
	if hero.Alive {
		return "&Omega;"

	} else {
		return "&#x2626;"
	}
}

var hero = NewHero()

var terrains = [2]string {"~", "."}
var view [][]string
var viewLandscape [][]string

var world_latitude = list.New()
var world_longitude = list.New()

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Path != "/" {
		handlerError(w, r, http.StatusNotFound)
		return
	}

// 	hero.X, _ = strconv.ParseInt(r.Form.Get("x"), 10, 64)
// 	hero.Y, _ = strconv.ParseInt(r.Form.Get("y"), 10, 64)

	executeTemplate(w, r)
}

func handlerMove(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Path != "/move/" {
		handlerError(w, r, http.StatusNotFound)
		return
	}

	direction := r.Form.Get("direction")

	if direction == "left" {
		hero.X--
	} else if direction == "right" {
		hero.X++
	} else if direction == "up" {
		hero.Y--
	} else if direction == "down" {
		hero.Y++
	}

	executeTemplate(w, r)
}

func executeTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("was: hero.Alive", hero.Alive)
	if !hero.Alive {
		hero.Init();
		fmt.Println("redirect to /", hero)
		http.Redirect(w, r, "/", 302)
		return
	}

	placeHero()
	fmt.Println("now: hero.Alive", hero.Alive)
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	
	if err != nil {
		handlerError(w, r, http.StatusInternalServerError)
		return
	}
	
	t.Execute(w, struct{View [][]string}{view})	
}


func handlerFavicon(w http.ResponseWriter, r *http.Request) {
	handlerError(w, r, http.StatusNotFound)
}

func handlerError(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Not Found")
	} else {
		fmt.Fprint(w, "Error")
	}
}

func placeHero() {
	for i := 0; i < viewSize; i++ {
		view[i] = make([]string, len(viewLandscape[i]))
		copy(view[i], viewLandscape[i]);
	}

	x := hero.X % viewSize
	y := hero.Y % viewSize

	if x < 0 {
		x = viewSize + x
	}

	if y < 0 {
		y = viewSize + y
	}

	place := &view[y][x]
	fmt.Println("place", *place)

	if *place == water {
		hero.Alive = false
	} else {
		hero.Alive = true
	}

	*place = hero.Symbol()
}

func generateLandscape() ([][]string) {
	var landscape [][]string

	for i := 0; i < viewSize; i++ {
		landscape = append(landscape, []string{})
		
		for j := 0; j < viewSize; j++ {
			landscape[i] = append(landscape[i], terrains[rand.Intn(len(terrains))])
		}
	}
	
	return landscape
}

func main() {
	fmt.Println("at start: hero.Alive", hero.Alive)
	viewLandscape = generateLandscape()
	viewLandscape[0][0] = ground
	view = make([][]string, len(viewLandscape))

	world_latitude.PushBack(viewLandscape)
	world_longitude.PushBack(viewLandscape)

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/move/", handlerMove)
	http.HandleFunc("/favicon.ico", handlerFavicon)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

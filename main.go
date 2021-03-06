package main

import (
	"fmt"
	"net/http"
// 	"strconv"
	"log"
	"text/template"
	"./world"
)

const viewSize = 16

type Hero struct {
	world.Located
	Land *world.Land
	Alive bool
}

func NewHero() (Hero) {
	hero := Hero{}
	hero.Init();
	return hero
}

func (hero *Hero) Init() {
	hero.Alive = true
	hero.X = 6
	hero.Y = 6
}

func (hero *Hero) Symbol() (string) {
	if hero.Alive {
		return "&Omega;"

	} else {
		return "&#x2626;"
	}
}

var hero = NewHero()
var view [][]string
var screen [][]string


func handlerRoot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Path != "/" {
		handlerError(w, r, http.StatusNotFound)
		return
	}

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
	t, err := template.ParseFiles("web/index.html")

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
	screen = world.GetScreen(hero.Land.X, hero.Land.Y, viewSize)

	for i := 0; i < viewSize; i++ {
		view[i] = make([]string, viewSize)
		copy(view[i], screen[i]);
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

	if *place == world.WaterTerrain.String() {
		hero.Alive = false
	} else {
		hero.Alive = true
	}

	*place = hero.Symbol()
}

func main() {
	view = make([][]string, viewSize)
	fmt.Println("at start: hero.Alive", hero.Alive)

	hero.Land = world.GetLand(0, 0)
	r := []int{0, 1, 2, 3, 4}

	for _, i := range r {
		for _, j := range r {
			hero.Land.Update(i, j, world.GroundTerrain)
		}
	}

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/move/", handlerMove)
	http.HandleFunc("/favicon.ico", handlerFavicon)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

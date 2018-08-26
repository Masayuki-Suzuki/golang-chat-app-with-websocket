package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data) // Should check return value in actual project.
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load .env file.")
	}
}

func main() {
	EnvLoad()

	SECURITY_KEY := os.Getenv("SECURITY_KEY")
	GOOGLE_AUTH_CLIENT := os.Getenv("GOOGLE_AUTH_CLIENT")
	GOOGLE_AUTH_KEY := os.Getenv("GOOGLE_AUTH_KEY")
	GITHUB_AUTH_CLIENT := os.Getenv("GITHUB_AUTH_CLIENT")
	GITHUB_AUTH_KEY := os.Getenv("GITHUB_AUTH_KEY")

	var addr = flag.String("addr", ":3000", "Application Address")
	flag.Parse()
	// Setup Gomniauth
	gomniauth.SetSecurityKey(SECURITY_KEY)
	gomniauth.WithProviders(
		google.New(GOOGLE_AUTH_CLIENT, GOOGLE_AUTH_KEY, "http://localhost:3000/auth/callback/google"),
		github.New(GITHUB_AUTH_CLIENT, GITHUB_AUTH_KEY, "http://localhost:3000/auth/callback/github"))
	r := newRoom()
	//r.tracer = trace.New(os.Stdout)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run() // Start chat room.
	log.Println("Starting web server. Port: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

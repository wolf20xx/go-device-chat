package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
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
	log.Println("call templateHandler", t.filename)
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アドレス")
	flag.Parse()
	// gomniauth
	gomniauth.SetSecurityKey("DevChatApps共通セキュリティキー")
	gomniauth.WithProviders(
		google.New("****.apps.googleusercontent.com", "****", "http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	//r.tracer = trace.New(os.Stdout)
	//http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("start Web Server,port:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

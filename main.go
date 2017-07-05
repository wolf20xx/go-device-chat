package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
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

type googleAuthInfo struct {
	web gInfo `json:"web"`
}

type gInfo struct {
	client_id                   string   `json:"client_id"`
	project_id                  string   `json:"project_id"`
	auth_uri                    string   `json:"auth_uri"`
	token_uri                   string   `json:"token_uri"`
	auth_provider_x509_cirt_url string   `json:"token_uri"`
	client_secret               string   `json:"client_secret"`
	redirect_uris               []string `json:"redirect_uris"`
	javascript_origin           []string `json:"javascript_origins"`
}

func main() {
	var addr = flag.String("addr", ":8080", "アドレス")
	flag.Parse()
	// configの読み込み
	clientInfo, _ := ioutil.ReadFile("./testinfo.json")
	//var data interface{}
	var data googleAuthInfo
	if err := json.Unmarshal(clientInfo, &data); err != nil {
		log.Fatal(err)
	}
	log.Println(data.web.client_id)

	// gomniauth
	gomniauth.SetSecurityKey("DevChatApps共通セキュリティキー")
	gomniauth.WithProviders(
		google.New(data.web.client_id, data.web.client_secret, "http://localhost:8080/auth/callback/google"),
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

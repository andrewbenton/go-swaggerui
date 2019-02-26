package swaggerui

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rakyll/statik/fs"
)

type Config struct {
	AppName    string
	SwaggerURL string
	RootUIPath string
}

func buildTemplate(def string, cfg Config) []byte {
	tmpl := template.Must(template.New("swaggerui").Parse(def))
	buff := new(bytes.Buffer)
	tmpl.Execute(buff, cfg)
	return buff.Bytes()
}

func Handle(cfg Config) http.Handler {
	statikFs, err := fs.New()

	if err != nil {
		panic(err)
	}

	data, err := fs.ReadFile(statikFs, "/layout.html")

	if err != nil {
		panic(err)
	}

	index := buildTemplate(string(data), cfg)

	router := mux.NewRouter()
	router.HandleFunc("/swagger-ui", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write(index)
	})
	router.
		PathPrefix("/swagger-ui/").
		Handler(http.StripPrefix("/swagger-ui", http.FileServer(statikFs)))

	return router
}

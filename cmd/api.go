package cmd

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mantasgra/cloud-lab/internal/controller"
	"github.com/spf13/cobra"
)

var httpServerCmd = &cobra.Command{
	Use:   "server",
	Short: "serves an application by starting an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func server() {
	logger := log.New(os.Stdout, "[api] ", log.Ltime)
	r := mux.NewRouter()
	r.HandleFunc("/username", controller.GithubUsernameHandler).Methods("GET")

	// TODO. Add a handler here which will use controller.Sha256Handler function for HTTP "GET" method. Use curly brackets for path matching, like this: /{username}
	r.HandleFunc("/{username}", controller.Sha256Handler).Methods("GET")
	// serve static files - DO NOT CHANGE THIS!
	r.PathPrefix("/ui/css").Handler(http.StripPrefix("/ui/css", http.FileServer(http.Dir("static/css"))))
	r.PathPrefix("/ui/js").Handler(http.StripPrefix("/ui/js", http.FileServer(http.Dir("static/js"))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      r,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	logger.Printf("started http server on port %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

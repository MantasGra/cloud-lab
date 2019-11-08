package controller

import (
	"fmt"
	"net/http"
	"os"
	"https://github.com/MantasGra/cloud-lab/blob/master/internal/hashes"
)

func Sha256Handler(w http.ResponseWriter, r *http.Request) {
	// to get the username, use the following:
	username := mux.Vars(r)["username"]

	// to calculate sha256 hash of a string, use internal/hashes package and function GetHash
	hash, _ := GetHash("Sha256", username)
	// to send the response, use the following:
	fmt.Fprint(w, hash)
	w.WriteHeader(http.StatusOK)
}

func GithubUsernameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, os.Getenv("GITHUB_USERNAME"))
}

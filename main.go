package main

import (
	"github.com/drone/drone-go/drone"
	"log"
	"os"
	"strings"
)

type Params struct {
	Repos  string `json:"repository"`
	Server string `json:"server"`
	Token  string `json:"token"`
}

func parseRepo(repo string) (string, string) {
	var (
		owner string
		name  string
	)

	parts := strings.Split(repo, "/")
	if len(parts) == 2 {
		owner = parts[0]
		name = parts[1]
	}
	return owner, name
}

func main() {

	token := os.Getenv("TOKEN")
	host := os.Getenv("HOST")
	repo := os.Getenv("REPO")
	branch := os.Getenv("BRANCH")

	client := drone.NewClientToken(host, token)
	owner, name := parseRepo(repo)

	log.Println(token)
	log.Println(host)
	log.Println(repo)
	log.Println(branch)

	// get the latest build for the specified repository
	build, err := client.BuildLast(owner, name, branch)
	if err != nil {
		log.Printf("Error: unable to get latest build for %s.\n", repo)
		os.Exit(1)
	}
	// start a new  build
	_, err = client.BuildFork(owner, name, build.Number)
	if err != nil {
		log.Printf("Error: unable to trigger a new build for %s.\n", repo)
		log.Panic(err)
	}

	log.Printf("Starting new build %d for %s\n", build.Number, repo)
}

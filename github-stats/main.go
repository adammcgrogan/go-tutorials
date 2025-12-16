package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GithubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	CreatedAt   string `json:"created_at"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Could not fetch user %s\n", username)
		os.Exit(1)
	}

	var user GithubUser

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	fmt.Println("---------------------------------------")
	fmt.Printf("User:      %s (%s)\n", user.Login, user.Name)
	fmt.Printf("Bio:       %s\n", user.Bio)
	fmt.Printf("Location:  %s\n", user.Location)
	fmt.Printf("Repos:     %d\n", user.PublicRepos)
	fmt.Printf("Followers: %d\n", user.Followers)
	fmt.Printf("Joined:    %s\n", user.CreatedAt)
	fmt.Println("---------------------------------------")

}

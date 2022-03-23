/*
Copyright © 2022 Olumide I. Nwosu <olumidenwosu@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type Repository struct {
	Name        string `json:"name"`
	Author      User   `json:"owner"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	Watchers    int    `json:"watchers_count"`
	Year        string `json:"created_at"`
}

func parseRepositoryData(r Repository, full bool) {
	fmt.Println("Name: " + r.Name)
	fmt.Println("Description: " + r.Description)
	if full {
		fmt.Println("Author: " + r.Author.Username)
		fmt.Println("Language: " + r.Language)
		fmt.Printf("Stars: %d\n", r.Stars)
		fmt.Printf("Forks: %d\n", r.Forks)
		fmt.Printf("Watchers: %d\n", r.Watchers)
		fmt.Printf("Year: %s\n", strings.Split(r.Year, "-")[0])
	}
	println()
}

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Fetches info for a github repository.",
	Long:  `Fetches info for a github repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		var repository_name, user, URL string
		if len(args) >= 1 && args[0] != "" {
			user = args[0]
			URL = "https://api.github.com/users/" + user + "/repos"
			if len(args) >= 2 && args[1] != "" {
				repository_name = args[1]
				URL = fmt.Sprintf("https://api.github.com/repos/%s/%s", user, repository_name)
			}
		} else {
			fmt.Println("Error: Repository name and username must be provided!")
			return
		}
		response, err := http.Get(URL)
		fmt.Println(URL)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			fmt.Println("Fetching " + user + "'s repository data...")
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatalln(err)
			}
			if repository_name != "" {
				var result Repository
				if err := json.Unmarshal(body, &result); err != nil {
					fmt.Println("Can not unmarshal JSON")
				}
				parseRepositoryData(result, true)
			} else {
				var result []Repository
				if err := json.Unmarshal(body, &result); err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(result); i++ {
					parseRepositoryData(result[i], false)
				}
			}
		} else {
			fmt.Println("Error: User " + user + " does not exist! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}

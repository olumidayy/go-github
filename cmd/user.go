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

	"github.com/spf13/cobra"
)

type User struct {
    Name       string `json:"name"`
    Username    string `json:"login"`
    Location      string `json:"location"`
    Followers int `json:"followers"`
    Following int `json:"following"`
    PublicRepos int `json:"public_repos"`
    PublicGists int `json:"public_gists"`
}

func parseUser(u User) {
	fmt.Println("Name: " + u.Name)
	fmt.Println("Username: " + u.Username)
	fmt.Printf("Followers: %d\n", u.Followers)
	fmt.Printf("Following: %d\n", u.Following)
	fmt.Printf("Public Repositories: %d\n", u.PublicRepos)
	fmt.Printf("Public Gists: %d\n", u.PublicGists)
	fmt.Println("Location: " + u.Location)
}

// userCmd represents the get command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of a github user",
	Long: `This command fetches the data a particluar github user by querying
	the Github API with their username.`,
	Run: func(cmd *cobra.Command, args []string) {
		var user string
		if len(args) >= 1 && args[0] != "" {
            user = args[0]
        } else {
			fmt.Println("Error: Username must be provided!")
			return
		}
		response, err := http.Get("https://api.github.com/users/" + user)
        if err != nil {
            log.Fatal(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            fmt.Println("Just fetched " + user + "'s data.")
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatalln(err)
			}
			var result User
			if err := json.Unmarshal(body, &result); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}
			parseUser(result)
        } else {
            fmt.Println("Error: User " + user + " does not exist! :-(")
        }
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

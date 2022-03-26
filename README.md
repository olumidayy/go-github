# go-github

### A [Golang](https://go.dev/) CLI tool built with [Cobra (& cobra-cli)](https://github.com/spf13/cobra).

## Pre-requisites
- [Go SDK](https://go.dev/dl/)


## Getting started

You can decide to do one of two things:

- Clone the repository
```
$ git clone https://github.com/olumidayy/go-github.git
```
 After cloning, you can then build the tool by running:
 ```
 $ go build
 ```
- Or just install tool directly in your project.
```
$ go install github.com/olumidayy/go-github@latest
```



## Usage
There are two commands, `user` and `repo` (so far). The former fetches information about a particular user by passing their github username (compulsory), like so:

```
$ go-github user <USERNAME>
```
Below is an example:
```
$ go-github user olumidayy

Just fetched olumidayy's data.
Name: Olumide Nwosu
Username: olumidayy
Followers: 33
Following: 16
Public Repositories: 26
Public Gists: 33
Location: Lagos, Nigeria
```

The `repo` command can be used in two ways:
 - Passing just a username:
 ```
 $ go-github repo <USERNAME>
 ```
 This returns the name and description of all the user's public repositories.
 - Passing both a username and a repository name as arguments:
 ```
 $ go-github repo <USERNAME> <REPO_NAME>
 ```
 This is more specific and returns a summary of the repository in question.
 
These can be demonstrated as follows:
```
 $ go-github repo olumidayy

 Fetching olumidayy's repository data...
 Name: bckdrp
 Description: A GraphQL API for shortening URLs

 Name: birthday_reminder
 Description: Utility app to remind of loved ones' birthdays made with flutter.
 .....
```

```
 $ go-github repo olumidayy go-github

 Fetching olumidayy's repository data...
 Name: go-github
 Description: A CLI tool for communicating with github.
 Author: olumidayy
 Language: Go
 Stars: 13
 Forks: 0
 Watchers: 13
 Year: 2022
```

## Contributions
 Pull requests and contributions are very much welcome! 💙
 
 

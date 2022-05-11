package main

import (
	artifactory "artifactory/modules"
	"fmt"
)

func main() {
	url := "https://dortsh.jfrog.io/artifactory/api"
	apiKey := "KEY"
	apiUser := "USER"

	servicesManager := artifactory.CreateASM(url, apiKey, apiUser)
	repos := artifactory.GetAllRemoteRepos(servicesManager)

	fmt.Printf("Type: %T", servicesManager)
	fmt.Printf("Repos: %s", repos)

	// response := artifactory.GetAllRepos()
	// log.Print(response)
}

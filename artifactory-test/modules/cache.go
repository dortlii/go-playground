package artifactory

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"log"
)

func GetAllRemoteRepos(servicesManager artifactory.ArtifactoryServicesManager) (repos []string) {
	params := services.NewRepositoriesFilterParams()
	params.RepoType = "remote"

	output, err := servicesManager.GetAllRepositoriesFiltered(params)

	// Error Handling
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(output)

	return repos
}

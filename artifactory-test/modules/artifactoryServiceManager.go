package artifactory

import (
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/config"
	"log"
	"time"
)

func CreateASM(url string, apiKey string, apiUser string) (serviceManager artifactory.ArtifactoryServicesManager) {
	// Creating Artifactory Details
	rtDetails := auth.NewArtifactoryDetails()
	rtDetails.SetUrl(url)
	rtDetails.SetApiKey(apiKey)

	// Creating Artifactory Service Config
	serviceConfig, err := config.NewConfigBuilder().
		SetServiceDetails(rtDetails).
		SetDryRun(true).
		SetHttpTimeout(180 * time.Second).
		SetHttpRetries(8).
		Build()

	// Error Handling
	if err != nil {
		log.Fatalln(err)
	}

	// Creating New Artifactory Service Manager
	rtManager, err := artifactory.New(serviceConfig)

	// Error Handling
	if err != nil {
		log.Fatalln(err)
	}

	return rtManager
}

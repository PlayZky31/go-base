package cmd

import (
	"log"

	"gandiwa/internal/config"
	"gandiwa/internal/core"
	rest_api "gandiwa/internal/presentation/restApi"
)

func Init(appExitChan chan bool) {

	conf := config.LoadConfigFile("resources/config/config.yaml")
	coreContainer := core.NewCoreContainer(conf)
	server := rest_api.NewFiberServer(coreContainer.SetUpUseCase)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("failed to start the server caused by: %s", err.Error())
		}
	}()

	go func() {
		<-appExitChan
		if err := server.Shutdown(); err != nil {
			log.Printf("failed to stop the server caused by: %s \n", err.Error())
		}

		if err := coreContainer.ShutDownDrivers(); err != nil {
			log.Printf("failed to stop the server caused by: %s \n", err.Error())
		}

		log.Println("all adapters has been shutdown")
		appExitChan <- true
	}()
}

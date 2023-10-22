package core

import "gandiwa/internal/config"

type Container interface {
	SetUpUseCase()
	SetUpDrivers()
	ShutDownDrivers() error
}

type coreContainer struct {
	conf     *config.Config
	useCases UseCases
}

type UseCases interface {
	EmployeeUseCase
}

func NewCoreContainer(conf *config.Config) Container {
	return &coreContainer{
		conf: conf,
	}
}

func (c coreContainer) SetUpDrivers() {

}

func (c coreContainer) ShutDownDrivers() error {
	return nil
}

func (c coreContainer) CallUseCases() UseCases {
	return c.useCases
}

func (c coreContainer) SetUpUseCase() {

}

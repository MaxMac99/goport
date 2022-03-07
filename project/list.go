package project

import (
	"io/ioutil"

	"github.com/docker/compose/v2/pkg/api"
)

type Stack struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Status string `json:"Status,omitempty"`
}

func GetStacks() ([]Stack, error) {
	baseDir := getBaseDir()
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	var projects []Stack
	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}
		project, err := GetProject(dir.Name())
		if err != nil {
			continue
		}
		projects = append(projects, Stack{
			Id:   dir.Name(),
			Name: project.Name,
		})
	}
	return projects, nil
}

func (s *composeService) GetActiveStacks(opts api.ListOptions) ([]Stack, error) {
	service := getComposeService(s.apiClient)
	stacks, err := service.List(s.ctx, opts)
	if err != nil {
		return nil, err
	}
	remoteStacks := make([]Stack, 0)
	for _, stack := range stacks {
		remoteStacks = append(remoteStacks, Stack{
			Id:     stack.ID,
			Name:   stack.Name,
			Status: stack.Status,
		})
	}
	return remoteStacks, nil
}

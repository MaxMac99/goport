package project

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
)

const (
	PROJECTS_BASE_DIR_ENV_VAR = "GOPORT_PROJECTS_BASE_DIR"
	DEFAULT_PROJECTS_BASE_DIR = "/projects"
)

func getBaseDir() string {
	if value, ok := os.LookupEnv(PROJECTS_BASE_DIR_ENV_VAR); ok {
		return value
	}
	return DEFAULT_PROJECTS_BASE_DIR
}

func GetProject(id string) (*types.Project, error) {
	path := filepath.Join(getBaseDir(), id)
	filename, err := findComposeFilename(path)
	if err != nil {
		return nil, err
	}
	path = filepath.Join(path, filename)
	options, err := getProjectOptions(path)
	if err != nil {
		return nil, err
	}
	project, err := cli.ProjectFromOptions(options)
	if err != nil {
		return nil, err
	}

	var profiles []string
	if envProfiles, ok := options.Environment["COMPOSE_PROFILES"]; ok {
		profiles = append(profiles, strings.Split(envProfiles, ",")...)
	}

	project.ApplyProfiles(profiles)

	project.WithoutUnnecessaryResources()
	return project, nil
}

func AddProject(id string, body []byte, format string) error {
	projectDir := filepath.Join(getBaseDir(), id)
	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		return errors.New("The project with the name \"" + id + "\" already exists.")
	}
	if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
		return err
	}
	suffix := "yml"
	if format == "json" {
		suffix = format
	}
	path := filepath.Join(projectDir, "docker-compose."+suffix)
	if err := os.WriteFile(path, body, 0644); err != nil {
		os.RemoveAll(projectDir)
		return err
	}
	return nil
}

func RemoveProject(id string) error {
	projectDir := filepath.Join(getBaseDir(), id)
	return os.RemoveAll(projectDir)
}

func findComposeFilename(path string) (string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.Name() == "docker-compose.yml" || file.Name() == "docker-compose.yaml" || file.Name() == "docker-compose.json" {
			return file.Name(), err
		}
	}
	return "", os.ErrNotExist
}

func getProjectOptions(path string) (*cli.ProjectOptions, error) {
	return cli.NewProjectOptions([]string{path}, []cli.ProjectOptionsFn{
		cli.WithDotEnv,
		cli.WithOsEnv,
		cli.WithConfigFileEnv,
		cli.WithDefaultConfigPath,
		cli.WithResolvedPaths(true),
	}...)
}

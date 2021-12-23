package compose

import (
	"io/ioutil"
	"os"
)

func ListProjects() []string {
	baseFolder := os.Getenv("DOCKPORT_PROJECTS_BASE_DIR")
	folders, err := ioutil.ReadDir(baseFolder)
	if err != nil {
		return nil
	}
	var result []string
	for _, folder := range folders {
		if !folder.IsDir() {
			continue
		}
		projectFiles, err := ioutil.ReadDir(baseFolder + "/" + folder.Name())
		if err != nil {
			continue
		}
		for _, file := range projectFiles {
			if file.Name() == "docker-compose.yml" {
				result = append(result, baseFolder+"/"+folder.Name())
			}
		}
	}
	return result
}

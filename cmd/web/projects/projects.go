package projects

import (
	"encoding/json"
	"log"
	"net/http"
)

type Project struct {
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Techs       []string `json:"techs"`
	Description string   `json:"description"`
	NumPhoto    int      `json:"numPhoto"`
	Extension   string   `json:"extension"`
	Repo        string   `json:"repo"`
	Demo        string   `json:"demo"`
	Size        string   `json:"size"`
}

type Category struct {
	Name     string `json:"name"`
	Slug     string `json:"string"`
	Projects []Project
}

func GetProjects() []Project {
	resp, err := http.Get("http://localhost:8080/api/projects")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var projects []Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		log.Fatal(err)
	}

	return projects
}

// func GetProject(slug string) Project{

// }

func GetSomeProjects(slugs []string) []Project {
	// fmt.Println("HERE")
	allProjects := GetProjects()
	// for _, project := range allProjects {
	// 	fmt.Println(project.Slug)
	// }
	var projects []Project
	for _, slug := range slugs {
		for _, project := range allProjects {
			// fmt.Println("slug, " ", project.Slug)
			// fmt.Printf("|%s|%s|\n", slug, project.Slug)
			if slug == project.Slug {
				// fmt.Println("Matched ", slug)
				projects = append(projects, project)
			}
		}
	}
	// for _, project := range projects {
	// 	fmt.Println(project.Slug)
	// }
	return projects
}

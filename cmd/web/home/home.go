package home

import (
	"bytes"
	"dotcom-2025/cmd/web/projects"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type PhraseHelloWorld struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

func GetHelloWorld() []string {
	data, err := os.ReadFile("cmd/web/assets/helloWorld.json")
	if err != nil {
		log.Fatal(err)
	}
	var phrases []PhraseHelloWorld
	err = json.NewDecoder(bytes.NewBuffer(data)).Decode(&phrases)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(phrases)
	var words []string
	for _, phrase := range phrases {
		words = append(words, phrase.Text)
	}
	// fmt.Println(words)
	return words
}

func GetRecentProject() []projects.Project {
	var slugs = []string{"boombatag-2024", "box-of-curiosity", "minesweeper"}
	projects := projects.GetSomeProjects(slugs[:])
	// for _, project := range projects {
	// 	fmt.Println(project.Slug)
	// }
	return projects
}

func GetShots() []string {
	var paths = []string{"9509", "9525", "9652"}
	for i, path := range paths {
		paths[i] = "assets/gallery/IMG_" + path + ".JPG"
	}
	return paths
}

type Education struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

func GetEducation() []Education {
	resp, err := http.Get("http://localhost:8080/api/educations")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var educations []Education
	err = json.NewDecoder(resp.Body).Decode(&educations)
	if err != nil {
		log.Fatal(err)
	}
	// for i, e := range educations {
	// 	fmt.Printf("Education %d %s %s %s %s %s\n", i, e.Name, e.Slug, e.Start, e.End, e.Description)
	// }
	return educations
}

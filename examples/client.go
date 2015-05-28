package main

import (
	"fmt"
	"os"

	"github.com/PeoplePerHour/shippable-go"
)

func main() {
	token := os.Getenv("apiToken")
	c := shippable.NewClient(token)
	projects, _, _ := c.Projects.GetProjects()
	proj := &projects[0]
	fmt.Println(*proj.Name)

	proj, _, err := c.Projects.GetProject("54e2287a5ab6cc13528c39d2")
	fmt.Printf("%+v\n", proj)

	builds, _, err := c.Projects.GetRunningBuilds(*proj.ID)
	fmt.Println(builds, err)

}

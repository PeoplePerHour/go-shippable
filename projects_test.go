package shippable

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetProjects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects.json")
		fmt.Fprint(w, string(rawResponse))
	})
	projects, _, err := client.Projects.GetProjects()
	if err != nil {
		t.Errorf("Projects.GetProjects returned error %v", err)
	}

	wantedProjID := "5564844fedd7f2c052fbe82f"
	project := (*projects)[0]
	projID := *project.ID
	if projID != wantedProjID {
		t.Errorf("Projects.GetProjects returned %s, wanted %s", projID, wantedProjID)
	}
}

func TestGetProject(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + ".json")
		fmt.Fprint(w, string(rawResponse))
	})

	project, _, err := client.Projects.GetProject(projID)
	if err != nil {
		t.Errorf("Projects.GetProject returned error %v", err)
	}

	wantedProjectOwner := "PeoplePerHour"
	if projectOwner := *project.SourceRepoOwner.Login; projectOwner != wantedProjectOwner {
		t.Errorf("Projects.GetProject returned %s, wanted %s", projectOwner, wantedProjectOwner)
	}

	if fetchedProjID := *project.ID; projID != fetchedProjID {
		t.Errorf("Projects.GetProject returned %s, wanted %s", fetchedProjID, projID)
	}
}

func TestGetRunningBuildsLimit(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID+"/runningBuilds/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + "_runningBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	_, _, err := client.Projects.GetRunningBuildsLimit(projID, 666)
	if err == nil {
		t.Errorf("Projects.GetRunningBuildsLimit should have returned error")
	}

	_, _, err = client.Projects.GetRunningBuildsLimit(projID, 5)
	if err != nil {
		t.Errorf("Projects.GetRunningBuildsLimit should have returned error %v", err)
	}
}
func TestGetRunningBuilds(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID+"/runningBuilds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + "_runningBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	builds, _, err := client.Projects.GetRunningBuilds(projID)
	if err != nil {
		t.Errorf("Projects.GetRunningBuilds returned error %v", err)
	}

	for _, build := range *builds {
		if *build.ProjectID != projID {
			t.Errorf("Projects.GetRunningBuilds returned %s, wanted %s", build.ProjectID, projID)
		}
	}
}

func TestGetRecentBuilds(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID+"/recentBuilds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + "_recentBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	builds, _, err := client.Projects.GetRecentBuilds(projID)
	if err != nil {
		t.Errorf("Projects.GetRecentBuilds returned error %v", err)
	}

	for _, build := range *builds {
		if *build.ProjectID != projID {
			t.Errorf("Projects.GetRecentBuilds returned %s, wanted %s", build.ProjectID, projID)
		}
	}
}

func TestGetRecentBuildsLimit(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID+"/recentBuilds/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + "_recentBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	_, _, err := client.Projects.GetRecentBuildsLimit(projID, 666)
	if err == nil {
		t.Errorf("Projects.GetRecentBuildsLimit should have returned error")
	}

	_, _, err = client.Projects.GetRecentBuildsLimit(projID, 5)
	if err != nil {
		t.Errorf("Projects.GetRecentBuildsLimit should have returned error %v", err)
	}
}

func TestGetQueuedBuilds(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID+"/queuedBuilds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + "_queuedBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	builds, _, err := client.Projects.GetQueuedBuilds(projID)
	if err != nil {
		t.Errorf("Projects.GetQueuedBuilds returned error %v", err)
	}
	if len(*builds) > 0 {
		t.Errorf("Projects.GetQueuedBuilds returned %d builds, expected %d", len(*builds), 0)
	}
}

func TestGetQueuedBuildsLimit(t *testing.T) {
	setup()
	defer teardown()
	projID := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projID+"/queuedBuilds/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projID + "_queuedBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	_, _, err := client.Projects.GetQueuedBuildsLimit(projID, 666)
	if err == nil {
		t.Errorf("Projects.GetQueuedBuildsLimit should have returned error")
	}

	builds, _, err := client.Projects.GetQueuedBuildsLimit(projID, 5)
	if err != nil {
		t.Errorf("Projects.GetQueuedBuildsLimit should have returned error %v", err)
	}
	if len(*builds) > 0 {
		t.Errorf("Projects.GetQueuedBuilds returned %d builds, expected %d", len(*builds), 0)
	}
}

func TestProjectsURLParserErrors(t *testing.T) {
	_, _, err := client.Projects.GetProject("%")
	testURLParseError(t, err)
	_, _, err = client.Projects.GetRunningBuildsLimit("%", 0)
	testURLParseError(t, err)
	_, _, err = client.Projects.GetQueuedBuildsLimit("%", 0)
	testURLParseError(t, err)
	_, _, err = client.Projects.GetRecentBuildsLimit("%", 0)
	testURLParseError(t, err)
}

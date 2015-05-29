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
	projId := *project.ID
	if projId != wantedProjID {
		t.Errorf("Projects.GetProjects returned %s, wanted %s", projId, wantedProjID)
	}
}

func TestGetProject(t *testing.T) {
	setup()
	defer teardown()
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + ".json")
		fmt.Fprint(w, string(rawResponse))
	})

	project, _, err := client.Projects.GetProject(projId)
	if err != nil {
		t.Errorf("Projects.GetProject returned error %v", err)
	}

	wantedProjectOwner := "PeoplePerHour"
	if projectOwner := *project.SourceRepoOwner.Login; projectOwner != wantedProjectOwner {
		t.Errorf("Projects.GetProject returned %s, wanted %s", projectOwner, wantedProjectOwner)
	}

	if fetchedProjId := *project.ID; projId != fetchedProjId {
		t.Errorf("Projects.GetProject returned %s, wanted %s", fetchedProjId, projId)
	}
}

func TestGetRunningBuildsLimit(t *testing.T) {
	setup()
	defer teardown()
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId+"/runningBuilds/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + "_runningBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	_, _, err := client.Projects.GetRunningBuildsLimit(projId, 666)
	if err == nil {
		t.Errorf("Projects.GetRunningBuildsLimit should have returned error")
	}

	_, _, err = client.Projects.GetRunningBuildsLimit(projId, 5)
	if err != nil {
		t.Errorf("Projects.GetRunningBuildsLimit should have returned error %v", err)
	}
}
func TestGetRunningBuilds(t *testing.T) {
	setup()
	defer teardown()
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId+"/runningBuilds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + "_runningBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	builds, _, err := client.Projects.GetRunningBuilds(projId)
	if err != nil {
		t.Errorf("Projects.GetRunningBuilds returned error %v", err)
	}

	for _, build := range *builds {
		if build.ProjectID != projId {
			t.Errorf("Projects.GetRunningBuilds returned %s, wanted %s", build.ProjectID, projId)
		}
	}
}

func TestGetRecentBuilds(t *testing.T) {
	setup()
	defer teardown()
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId+"/recentBuilds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + "_recentBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	builds, _, err := client.Projects.GetRecentBuilds(projId)
	if err != nil {
		t.Errorf("Projects.GetRecentBuilds returned error %v", err)
	}

	for _, build := range *builds {
		if build.ProjectID != projId {
			t.Errorf("Projects.GetRecentBuilds returned %s, wanted %s", build.ProjectID, projId)
		}
	}
}

func TestGetRecentBuildsLimit(t *testing.T) {
	setup()
	defer teardown()
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId+"/recentBuilds/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + "_recentBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	_, _, err := client.Projects.GetRecentBuildsLimit(projId, 666)
	if err == nil {
		t.Errorf("Projects.GetRecentBuildsLimit should have returned error")
	}

	_, _, err = client.Projects.GetRecentBuildsLimit(projId, 5)
	if err != nil {
		t.Errorf("Projects.GetRecentBuildsLimit should have returned error %v", err)
	}
}

func TestGetQueuedBuilds(t *testing.T) {
	setup()
	defer teardown()
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId+"/queuedBuilds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + "_queuedBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	builds, _, err := client.Projects.GetQueuedBuilds(projId)
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
	projId := "556734d0edd7f2c052ff35b4"
	mux.HandleFunc("/projects/"+projId+"/queuedBuilds/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_projects_" + projId + "_queuedBuilds.json")
		fmt.Fprint(w, string(rawResponse))
	})
	_, _, err := client.Projects.GetQueuedBuildsLimit(projId, 666)
	if err == nil {
		t.Errorf("Projects.GetQueuedBuildsLimit should have returned error")
	}

	builds, _, err := client.Projects.GetQueuedBuildsLimit(projId, 5)
	if err != nil {
		t.Errorf("Projects.GetQueuedBuildsLimit should have returned error %v", err)
	}
	if len(*builds) > 0 {
		t.Errorf("Projects.GetQueuedBuilds returned %d builds, expected %d", len(*builds), 0)
	}
}

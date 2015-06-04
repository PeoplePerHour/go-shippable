package shippable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestEnableBuild(t *testing.T) {
	setup()
	defer teardown()

	wantedProjID := "556734d0edd7f2c052ff35b4"
	projectInput := &ProjectInput{
		ProjectID: &wantedProjID,
	}

	mux.HandleFunc("/workflow/enableRepoBuild", func(w http.ResponseWriter, r *http.Request) {
		req := new(ProjectInput)
		json.NewDecoder(r.Body).Decode(req)
		if !reflect.DeepEqual(req, projectInput) {
			t.Errorf("Request = %+v, wanted %+v", req, projectInput)
		}
		testMethod(t, r, "POST")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_post_workflow_enableRepoBuild.json")
		fmt.Fprint(w, string(rawResponse))
	})

	project, _, err := client.Workflow.EnableBuild(projectInput)
	if err != nil {
		t.Errorf("Workflow.EnableBuild returned error %v", err)
	}

	projId := *project.ID
	if projId != wantedProjID {
		t.Errorf("Workflow.EnableBuild returned %s, wanted %s", projId, wantedProjID)
	}
}

func TestDisableBuild(t *testing.T) {
	setup()
	defer teardown()

	wantedProjID := "556734d0edd7f2c052ff35b4"
	projectInput := &ProjectInput{
		ProjectID: &wantedProjID,
	}

	mux.HandleFunc("/workflow/disableBuild", func(w http.ResponseWriter, r *http.Request) {
		req := new(ProjectInput)
		json.NewDecoder(r.Body).Decode(req)
		if !reflect.DeepEqual(req, projectInput) {
			t.Errorf("Request = %+v, wanted %+v", req, projectInput)
		}
		testMethod(t, r, "POST")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_post_workflow_disableBuild.json")
		fmt.Fprint(w, string(rawResponse))
	})

	_, _, err := client.Workflow.DisableBuild(projectInput)
	// omg, ugly
	// TODO: Persuade Shippable guys to fix their API to return JSON...
	// see: https://github.com/Shippable/support/issues/1489#issuecomment-107070182
	if err != nil && err.Error() != "invalid character 'D' looking for beginning of value" {
		t.Errorf("Workflow.DisableBuild returned error %v", err)
	}
}

func TestDisableBuild_Proper(t *testing.T) {
	setup()
	defer teardown()

	wantedProjID := "556734d0edd7f2c052ff35b4"
	projectInput := &ProjectInput{
		ProjectID: &wantedProjID,
	}

	mux.HandleFunc("/workflow/disableBuild", func(w http.ResponseWriter, r *http.Request) {
		req := new(ProjectInput)
		json.NewDecoder(r.Body).Decode(req)
		if !reflect.DeepEqual(req, projectInput) {
			t.Errorf("Request = %+v, wanted %+v", req, projectInput)
		}
		testMethod(t, r, "POST")
		fmt.Fprint(w, string("[]"))
	})

	ok, _, err := client.Workflow.DisableBuild(projectInput)
	if err != nil {
		t.Errorf("Workflow.DisableBuild returned error %v", err)
	}
	if !ok {
		t.Errorf("Workflow.DisableBuild should have returned ok")
	}
}

func TestDisableBuild_FailedRequest(t *testing.T) {

	_, _, err := client.Workflow.DisableBuild(&ProjectInput{})
	if err == nil {
		t.Errorf("Workflow.DisableBuild should have returned error")
	}
}

func TestTriggerBuild(t *testing.T) {
	setup()
	defer teardown()

	wantedProjID := "556734d0edd7f2c052ff35b4"
	branch := "master"
	expectedBuildID := "f00b4r"
	triggerBuildInput := &TriggerBuildInput{
		ProjectID: &wantedProjID,
		Branch:    &branch,
	}

	mux.HandleFunc("/workflow/triggerBuild", func(w http.ResponseWriter, r *http.Request) {
		req := new(TriggerBuildInput)
		json.NewDecoder(r.Body).Decode(req)
		if !reflect.DeepEqual(req, triggerBuildInput) {
			t.Errorf("Request = %+v, wanted %+v", req, triggerBuildInput)
		}
		testMethod(t, r, "POST")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_post_workflow_triggerBuild.json")
		fmt.Fprint(w, string(rawResponse))
	})

	b, _, err := client.Workflow.TriggerBuild(triggerBuildInput)
	if err != nil {
		t.Errorf("Workflow.TriggerBuild returned error %v", err)
	}
	if *b.BuildID != expectedBuildID {
		t.Errorf("Expected ID '%s', got '%s'", expectedBuildID, *b.BuildID)
	}
}

func TestValidateDockerHubCredentials(t *testing.T) {
	setup()
	defer teardown()

	username := "username"
	password := "password"
	email := "username@example.com"
	dockerHubCredentials := &DockerHubCredentials{
		Username: &username,
		Password: &password,
		Email:    &email,
	}

	mux.HandleFunc("/workflow/validateDockerHubCredentials", func(w http.ResponseWriter, r *http.Request) {
		req := new(DockerHubCredentials)
		json.NewDecoder(r.Body).Decode(req)
		if !reflect.DeepEqual(req, dockerHubCredentials) {
			t.Errorf("Request = %+v, wanted %+v", req, dockerHubCredentials)
		}
		testMethod(t, r, "POST")
		// Shippable API currently returns 500. Never seen a proper response. We 'll just test return 200
		//rawResponse, _ := ioutil.ReadFile("./mocks/response_post_workflow_validateDockerHubCredentials.json")
		//fmt.Fprint(w, string(rawResponse))
		fmt.Fprint(w, "")
	})

	ok, _, err := client.Workflow.ValidateDockerHubCredentials(dockerHubCredentials)
	if err != nil {
		t.Errorf("Workflow.ValidateDockerHubCredentials returned error %v", err)
	}
	if !ok {
		t.Errorf("Workflow.ValidateDockerHubCredentials should succeed")
	}
}

func TestValidateDockerHubCredentials_FailedRequest(t *testing.T) {
	setup()
	defer teardown()

	ok, _, _ := client.Workflow.ValidateDockerHubCredentials(&DockerHubCredentials{})
	if ok {
		t.Errorf("Workflow.ValidateDockerHubCredentials should have failed")
	}
}

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

	ok, _, err := client.Workflow.DisableBuild(projectInput)
	// omg, ugly
	// TODO: Persuade Shippable guys to fix their API to return JSON...
	// see: https://github.com/Shippable/support/issues/1489#issuecomment-107070182
	if err != nil && err.Error() != "invalid character 'D' looking for beginning of value" {
		t.Errorf("Workflow.DisableBuild returned error %v", err)
	}

	if !ok {
		t.Errorf("Workflow.DisableBuild should have returned 'ok'")
	}

}

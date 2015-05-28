package shippable

import (
	"fmt"
	"net/http"
	//	"reflect"
	"testing"
	//	"time"
)

func TestGetProjects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse := `[{
						"repositoryProvider": "github",
						"sourceId": "36298993",
						"updatedDate": "2015-05-28T14:32:50.632Z",
						"sourceRepoOwner": {
						  "login": "pmoust",
						  "id": 2493339,
						  "avatar_url": "https://avatars.githubusercontent.com/u/2493339?v=3",
						  "gravatar_id": "",
						  "url": "https://api.github.com/users/pmoust",
						  "html_url": "https://github.com/pmoust",
						  "followers_url": "https://api.github.com/users/pmoust/followers",
						  "following_url": "https://api.github.com/users/pmoust/following{/other_user}",
						  "gists_url": "https://api.github.com/users/pmoust/gists{/gist_id}",
						  "starred_url": "https://api.github.com/users/pmoust/starred{/owner}{/repo}",
						  "subscriptions_url": "https://api.github.com/users/pmoust/subscriptions",
						  "organizations_url": "https://api.github.com/users/pmoust/orgs",
						  "repos_url": "https://api.github.com/users/pmoust/repos",
						  "events_url": "https://api.github.com/users/pmoust/events{/privacy}",
						  "received_events_url": null,
						  "type": "User",
						  "site_admin": false
						},
						"sourceUpdated": "2015-05-26T10:45:43.000Z",
						"sourceCreated": "2015-05-26T13:50:07.000Z",
						"sourcePushed": "2015-05-22T22:15:02.000Z",
						"repositorySshUrl": "git@github.com:pmoust/nginx-proxy.git",
						"isFork": true,
						"isPrivateRepository": false,
						"sourceDescription": "Automated nginx proxy for Docker containers using docker-gen",
						"repositoryUrl": "https://api.github.com/repos/pmoust/nginx-proxy",
						"fullName": "pmoust/nginx-proxy",
						"name": "nginx-proxy",
						"subscriptionId": "5480429dd46935d5fbbf5214",
						"cacheTag": 0,
						"projectAuthorizationLastSyncEndDate": "1970-01-01T00:00:00.000Z",
						"created": "2015-05-26T14:35:13.727Z",
						"settings": {
						  "environmentVariables": [],
						  "imageOptions": {
							"ports": [],
							"mounts": []
						  }
						},
						"autoBuild": false,
						"sourceSize": 134,
						"sourceWatchersCount": 0,
						"sourceStargazersCount": 0,
						"sourceForksCount": 0,
						"branches": [],
						"id": "5564844fedd7f2c052fbe82f"}]`
		fmt.Fprint(w, rawResponse)
	})
	projects, _, err := client.Projects.GetProjects()
	if err != nil {
		t.Errorf("Projects.GetProjects returned  error %v", err)
	}

	wantedProjID := "5564844fedd7f2c052fbe82f"
	if projId := *projects[0].ID; projId != wantedProjID {
		t.Errorf("Projects.GetProjects returned %s, wanted %s", projId, wantedProjID)
	}
}

func TestGetProject(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/556734d0edd7f2c052ff35b4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse := `{ "repositoryProvider": "github",
						  "sourceId": "36444563",
						  "updatedDate": "2015-05-28T15:32:54.515Z",
						  "sourceRepoOwner": {
							"site_admin": false,
							"type": "Organization",
							"received_events_url": null,
							"events_url": "https://api.github.com/users/PeoplePerHour/events{/privacy}",
							"repos_url": "https://api.github.com/users/PeoplePerHour/repos",
							"organizations_url": "https://api.github.com/users/PeoplePerHour/orgs",
							"subscriptions_url": "https://api.github.com/users/PeoplePerHour/subscriptions",
							"starred_url": "https://api.github.com/users/PeoplePerHour/starred{/owner}{/repo}",
							"gists_url": "https://api.github.com/users/PeoplePerHour/gists{/gist_id}",
							"following_url": "https://api.github.com/users/PeoplePerHour/following{/other_user}",
							"followers_url": "https://api.github.com/users/PeoplePerHour/followers",
							"html_url": "https://github.com/PeoplePerHour",
							"url": "https://api.github.com/users/PeoplePerHour",
							"gravatar_id": "",
							"avatar_url": "https://avatars.githubusercontent.com/u/675387?v=3",
							"id": 675387,
							"login": "PeoplePerHour"
						  },
						  "sourceUpdated": "2015-05-28T14:42:58.000Z",
						  "sourceCreated": "2015-05-28T14:37:24.000Z",
						  "sourcePushed": "2015-05-28T14:42:57.000Z",
						  "repositorySshUrl": "git@github.com:PeoplePerHour/shippable.git",
						  "isFork": false,
						  "isPrivateRepository": true,
						  "sourceDescription": "Shippable API client in Go",
						  "repositoryUrl": "https://api.github.com/repos/PeoplePerHour/shippable",
						  "fullName": "PeoplePerHour/shippable",
						  "name": "shippable",
						  "subscriptionId": "5480429dd46935d5fbbf5213",
						  "enabledDate": "2015-05-28T15:33:06.954Z",
						  "projectAuthorizationLastSyncStartDate": "2015-05-28T15:33:07.026Z",
						  "ownerTokenPresent": true,
						  "sourceDefaultBranch": "master",
						  "language": "go",
						  "cacheTag": 0,
						  "enabledByAccount": {
							"id": "5480425b0f168213009d1229",
							"identityUsedToEnable": {
							  "id": "5480425b0f168213009d122a",
							  "provider": "github",
							  "userName": "pmoust"
							}
						  },
						  "projectAuthorizationLastSyncEndDate": "2015-05-28T15:37:19.506Z",
						  "created": "2015-05-28T15:32:54.515Z",
						  "settings": {
							"environmentVariables": [],
							"imageOptions": {
							  "ports": [],
							  "mounts": []
							}
						  },
						  "deployKey": {
							"public": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDK18REmlDsfesCSZpa+7qYAC0Lq9AFguelWwFUEKvuj99V2LgoX96PyxgqHzKlbekG5ceHmCDXENyk6mp/0wHDwnQxEPppAcSdaX/CQf81Hbxaqb3z53sXFBk3E7r6rUycFuCNK9w9FLT4oGur6iMaiVD+nT193Rso2tvBXjCxX1JBaPpKRpcLD82UBBIl5/alc7ZjHPo3owqUxokxMtSbMpA4clwUS4/QMNtuxJrjFYiNSF/GmTDY3PLNRO9nVty01uyMONUHzwgWYZ19H60PacPwDNuLAU/JrTua2Kaqhw6Y46R1v9fRv9PNOg/iahXfk6ERb9NZLIYiB1slRo1j Shippable\n"
						  },
						  "autoBuild": true,
						  "sourceSize": null,
						  "sourceWatchersCount": 0,
						  "sourceStargazersCount": 0,
						  "sourceForksCount": 1,
						  "branches": [
							"master"
						  ],
						  "id": "556734d0edd7f2c052ff35b4"
						}`
		fmt.Fprint(w, rawResponse)
	})

	projId := "556734d0edd7f2c052ff35b4"
	project, _, err := client.Projects.GetProject(projId)
	if err != nil {
		t.Errorf("Projects.GetProject returned  error %v", err)
	}

	wantedProjectOwner := "PeoplePerHour"

	if projectOwner := *project.SourceRepoOwner.Login; projectOwner != wantedProjectOwner {
		t.Errorf("Projects.GetProject returned %s, wanted %s", projectOwner, wantedProjectOwner)
	}
}

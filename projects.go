package shippable

import (
	"fmt"
	"time"
)

// GetProjects returns a list of projects, and some info about them
func (p *ProjectService) GetProjects() (projects *[]Project, resp *Response, err error) {
	req, err := p.client.NewRequest("GET", "projects", nil)
	if err != nil {
		return nil, nil, err
	}
	projects = new([]Project)
	resp, err = p.client.Do(req, projects)
	return
}

// GetProject returns a more in-depth information about a specific
func (p *ProjectService) GetProject(id string) (project *Project, resp *Response, err error) {
	req, err := p.client.NewRequest("GET", "projects/"+id, nil)
	if err != nil {
		return nil, nil, err
	}
	project = new(Project)
	resp, err = p.client.Do(req, project)
	return
}

// GetRunningBuilds returns a list of a project's currently running builds
func (p *ProjectService) GetRunningBuilds(projectID string) (builds *[]Build, resp *Response, err error) {
	return p.GetRunningBuildsLimit(projectID, 0)
}

// GetRunningBuildsLimit returns a list of a project's currently running builds up to a limit
func (p *ProjectService) GetRunningBuildsLimit(projectID string, limit int) (builds *[]Build, resp *Response, err error) {
	url := fmt.Sprintf("projects/%s/runningBuilds", projectID)
	if limit > 0 {
		url = fmt.Sprintf("projects/%s/runningBuilds/%d", projectID, limit)
	}
	req, err := p.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	builds = new([]Build)
	resp, err = p.client.Do(req, builds)
	return
}

// GetQueuedBuilds returns a list of a project's currently queued builds
func (p *ProjectService) GetQueuedBuilds(projectID string) (builds *[]Build, resp *Response, err error) {
	return p.GetQueuedBuildsLimit(projectID, 0)
}

// GetQueuedBuildsLimit returns a list of a project's currently queued builds up to a limit
func (p *ProjectService) GetQueuedBuildsLimit(projectID string, limit int) (builds *[]Build, resp *Response, err error) {
	url := fmt.Sprintf("projects/%s/queuedBuilds", projectID)
	if limit > 0 {
		url = fmt.Sprintf("projects/%s/queuedBuilds/%d", projectID, limit)
	}
	req, err := p.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	builds = new([]Build)
	resp, err = p.client.Do(req, builds)
	return
}

// GetRecentBuilds returns a list of a project's currently recent builds
func (p *ProjectService) GetRecentBuilds(projectID string) (builds *[]Build, resp *Response, err error) {
	return p.GetRecentBuildsLimit(projectID, 0)
}

// GetRecentBuildsLimit returns a list of a project's currently recent builds up to a limit
func (p *ProjectService) GetRecentBuildsLimit(projectID string, limit int) (builds *[]Build, resp *Response, err error) {
	url := fmt.Sprintf("projects/%s/recentBuilds", projectID)
	if limit > 0 {
		url = fmt.Sprintf("projects/%s/recentBuilds/%d", projectID, limit)
	}
	req, err := p.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	builds = new([]Build)
	resp, err = p.client.Do(req, builds)
	return
}

// EnableBuilds enable builds for a given project
func (c *Client) EnableBuilds(p *Project) (project *Project, err error) {
	return &Project{}, nil
}

// DisableBuilds disables builds for a given project
func (c *Client) DisableBuilds(p *Project) (project *Project, err error) {
	return &Project{}, nil
}

// TriggerBuild triggers a build from a project's default branch
func (c *Client) TriggerBuild(p *Project) (build *Build, err error) {
	return &Build{}, nil
}

// TriggerBuildFrom triggers a build from a project's given branch
func (c *Client) TriggerBuildFrom(p *Project, branch string) (build *Build, err error) {
	return &Build{}, nil
}

type Project struct {
	AutoBuild                             *bool            `json:"autoBuild"`
	Branches                              *[]string        `json:"branches"`
	CacheTag                              *int             `json:"cacheTag"`
	Created                               *string          `json:"created"`
	DeployKey                             *DeployKey       `json:"deployKey"`
	EnabledDate                           *time.Time       `json:"enabledDate"`
	FullName                              *string          `json:"fullName"`
	ID                                    *string          `json:"id"`
	IsEnabled                             *bool            `json:"isEnabled"`
	IsFork                                *bool            `json:"isFork"`
	IsPrivateRepository                   *bool            `json:"isPrivateRepository"`
	Language                              *string          `json:"language"`
	Name                                  *string          `json:"name"`
	OwnerTokenPresent                     *bool            `json:"ownerTokenPresent"`
	ProjectAuthorizationLastSyncEndDate   *string          `json:"projectAuthorizationLastSyncEndDate"`
	ProjectAuthorizationLastSyncStartDate *string          `json:"projectAuthorizationLastSyncStartDate"`
	RepositoryProvider                    *string          `json:"repositoryProvider"`
	RepositorySSHURL                      *string          `json:"repositorySshUrl"`
	RepositoryURL                         *string          `json:"repositoryUrl"`
	Settings                              *ProjectSettings `json:"settings"`
	SourceCreated                         *time.Time       `json:"sourceCreated"`
	SourceDefaultBranch                   *string          `json:"sourceDefaultBranch"`
	SourceDescription                     *string          `json:"sourceDescription"`
	SourceForksCount                      *int             `json:"sourceForksCount"`
	SourceID                              *string          `json:"sourceId"`
	SourcePushed                          *string          `json:"sourcePushed"`
	SourceRepoOwner                       *GithubUser      `json:"sourceRepoOwner"`
	SourceSize                            *int             `json:"sourceSize"`
	SourceStargazersCount                 *int             `json:"sourceStargazersCount"`
	SourceUpdated                         *string          `json:"sourceUpdated"`
	SourceWatchersCount                   *int             `json:"sourceWatchersCount"`
	SubscriptionID                        *string          `json:"subscriptionId"`
	UpdatedDate                           *time.Time       `json:"updatedDate"`
}

type DeployKey struct {
	Public *string `json:"public"`
}

type ProjectSettings struct {
	EnvironmentVariables *[]string `json:"environmentVariables"`
	ImageOptions         *ImageOptions
}

type ImageOptions struct {
	Mounts *[]string `json:"mounts"`
	Ports  *[]string `json:"ports"`
}

type GithubUser struct {
	AvatarURL         *string `json:"avatar_url"`
	EventsURL         *string `json:"events_url"`
	FollowersURL      *string `json:"followers_url"`
	FollowingURL      *string `json:"following_url"`
	GistsURL          *string `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           *string `json:"html_url"`
	ID                *int    `json:"id"`
	Login             *string `json:"login"`
	OrganizationsURL  *string `json:"organizations_url"`
	ReceivedEventsURL *string `json:"received_events_url"`
	ReposURL          *string `json:"repos_url"`
	SiteAdmin         *bool   `json:"site_admin"`
	StarredURL        *string `json:"starred_url"`
	SubscriptionsURL  *string `json:"subscriptions_url"`
	Type              *string `json:"type"`
	URL               *string `json:"url"`
}

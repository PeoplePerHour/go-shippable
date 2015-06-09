package shippable

// ProjectInput is passwd to WorkflowService to perform operations like enabling/disabling
// automatic Builds
type ProjectInput struct {
	ProjectID *string
}

// TriggerBuildInput is passed to WorkflowService in order to trigger a build for a project.
// Omit branch to build the project's default branch.
type TriggerBuildInput struct {
	ProjectID *string
	Branch    *string
}

// TriggerBuildOutput is returned after successfully triggered a build via Shippable API.
type TriggerBuildOutput struct {
	BuildID *string
}

// DockerHubCredentials is used to authenticated against DockerHub registry
type DockerHubCredentials struct {
	Username *string
	Password *string
	Email    *string
}

// EnableBuild enable automatic builds for a given project.
func (w *WorkflowService) EnableBuild(p *ProjectInput) (project *Project, resp *Response, err error) {
	url := "workflow/enableRepoBuild"
	req, _ := w.client.NewRequest("POST", url, p)
	project = new(Project)
	resp, err = w.client.Do(req, project)
	return
}

// DisableBuild disables builds for a given project. This will delete project's history
// and its associated builds from Shippable.
func (w *WorkflowService) DisableBuild(p *ProjectInput) (ok bool, resp *Response, err error) {
	url := "workflow/disableBuild"
	req, _ := w.client.NewRequest("POST", url, p)
	resp, err = w.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}
	return true, resp, err
}

// TriggerBuild triggers a build from a given project and branch. If no branch is
// specified the default branch configured for the project in Github or Bitbucket
// will be built.
func (w *WorkflowService) TriggerBuild(t *TriggerBuildInput) (build *TriggerBuildOutput, resp *Response, err error) {
	url := "workflow/triggerBuild"
	req, err := w.client.NewRequest("POST", url, t)
	build = new(TriggerBuildOutput)
	resp, err = w.client.Do(req, build)
	return
}

// ValidateDockerHubCredentials verifies a DockerHub account for the authenticated Shippable API user.
func (w *WorkflowService) ValidateDockerHubCredentials(c *DockerHubCredentials) (ok bool, resp *Response, err error) {
	url := "workflow/validateDockerHubCredentials"
	ok = false
	req, _ := w.client.NewRequest("POST", url, c)
	resp, err = w.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}
	if resp.StatusCode < 300 {
		ok = true
	}
	return
}

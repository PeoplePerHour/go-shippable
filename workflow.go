package shippable

// ProjectInput
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

// EnableBuild enable automatic builds for a given project.
func (w *WorkflowService) EnableBuild(p *ProjectInput) (project *Project, resp *Response, err error) {
	url := "workflow/enableRepoBuild"
	req, err := w.client.NewRequest("POST", url, p)
	if err != nil {
		return nil, nil, err
	}
	project = new(Project)
	resp, err = w.client.Do(req, project)
	return
}

// DisableBuild disables builds for a given project. This will delete project's history
// and its associated builds from Shippable.
func (w *WorkflowService) DisableBuild(p *ProjectInput) (ok bool, resp *Response, err error) {
	url := "workflow/disableBuild"
	req, err := w.client.NewRequest("POST", url, p)
	if err != nil {
		return false, nil, err
	}
	resp, err = w.client.Do(req, ok)
	if resp.StatusCode > 299 {
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
	if err != nil {
		return nil, nil, err
	}
	build = new(TriggerBuildOutput)
	resp, err = w.client.Do(req, build)
	return
}

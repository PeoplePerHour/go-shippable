package shippable

// ProjectInput
type ProjectInput struct {
	ProjectID *string
}

// TriggerBuildInput
type TriggerBuildInput struct {
	ProjectID *string
	Branch    *string
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
func (w *WorkflowService) TriggerBuild(t *TriggerBuildInput) (build *Build, resp *Response, err error) {
	url := "workflow/triggerBuild"
	req, err := w.client.NewRequest("POST", url, t)
	if err != nil {
		return nil, nil, err
	}
	build = new(Build)
	resp, err = w.client.Do(req, build)
	return
}

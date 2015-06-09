package shippable

// ProjectService interacts with the /projects endpoint of Shippable API.
type ProjectService struct {
	client *Client
}

// AccountService interacts with the /acounts endpoint of Shippable API.
type AccountService struct {
	client *Client
}

// WorkflowService interacts with the /workflow endpoint of Shippable API.
type WorkflowService struct {
	client *Client
}

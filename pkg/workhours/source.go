package workhours

type Source struct {
	Toggl Toggl `json:"toggl"`
}

type Toggl struct {
	APIKey      string `json:"api_key"`
	WorkspaceID string `json:"workspace_id"`
}

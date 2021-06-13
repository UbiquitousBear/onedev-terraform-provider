package provider

type Project struct {
	Id                     int    `json:"id,omitempty"`
	ForkedFromId           int    `json:"forkedFromId,omitempty"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	IssueManagementEnabled bool   `json:"issueManagementEnabled"`
}

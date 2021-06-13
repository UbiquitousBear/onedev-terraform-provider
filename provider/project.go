package provider

type Project struct {
	Id                     int    `json:"id,omitempty"`
	ForkedFromId           int    `json:"forkedFromId"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	IssueManagementEnabled bool   `json:"issueManagementEnabled"`
}

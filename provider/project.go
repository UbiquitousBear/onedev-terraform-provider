package provider

type Project struct {
	Id							int      `json:"id"`
	ForkedFromId				int      `json:"forkedFromId"`
	Name        				string   `json:"name"`
	Description 				string   `json:"description"`
	IssueManagementEnabled		bool     `json:"issueManagementEnabled"`
}

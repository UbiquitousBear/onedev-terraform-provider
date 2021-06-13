provider "onedev" {}

resource "onedev_project" "test" {
  name                   = "test-project"
  description            = "test description"
  issueManagementEnabled = false
}

provider "onedev" {
  address = "http://devportal.sh"
}

resource "onedev_project" "test" {
  name                   = "test-project"
  description            = "test description"
  issueManagementEnabled = false
}

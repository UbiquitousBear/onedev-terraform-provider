terraform {
  required_providers {
    onedev = {
      source  = "hashicorp/onedev"
    }
  }
}

provider "onedev" {
  address = "http://devportal.sh"
}

resource "onedev_test" "test" {
  name                   = "test-project"
  description            = "test description"
  issuemanagementenabled = false
}

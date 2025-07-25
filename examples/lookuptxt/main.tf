terraform {
  required_providers {
    iputils = {
      source = "registry.terraform.io/krisiasty/iputils"
    }
  }
  required_version = ">= 1.8.0"
}

provider "iputils" {
}


output "example" {
  value = provider::iputils::lookuptxt("example.com")
}

output "google" {
  value = provider::iputils::lookuptxt("google.com")
}

# non-existing domain
# output "nonexisting" {
#   value = provider::iputils::lookuptxt("nonexistingdomain.xyz")
# }


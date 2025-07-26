terraform {
  required_providers {
    iputils = {
      source = "krisiasty/iputils"
    }
  }
  required_version = ">= 1.8.0"
}

provider "iputils" {
}

output "google" {
  value = provider::iputils::lookupaddr("google.com")
}

output "cloudflare" {
  value = provider::iputils::lookupaddr("cloudflare.com")
}

output "example" {
  value = provider::iputils::lookupaddr("example.com")
}

# non-existing domain
# output "nonexisting" {
#   value = provider::iputils::lookupaddr("nonexistingdomain.xyz")
# }

output "localhost" {
  value = provider::iputils::lookupaddr("localhost")
}

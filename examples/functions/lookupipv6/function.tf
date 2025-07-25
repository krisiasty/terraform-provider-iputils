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

output "google" {
  value = provider::iputils::lookupipv6("google.com")
}

output "cloudflare" {
  value = provider::iputils::lookupipv6("cloudflare.com")
}

output "example" {
  value = provider::iputils::lookupipv6("example.com")
}

# non-existing domain
# output "nonexisting" {
#   value = provider::iputils::lookupipv6("nonexistingdomain.xyz")
# }

output "localhost" {
  value = provider::iputils::lookupipv6("localhost")
}

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
  value = provider::iputils::lookupipv4("google.com")
}

output "cloudflare" {
  value = provider::iputils::lookupipv4("cloudflare.com")
}

output "example" {
  value = provider::iputils::lookupipv4("example.com")
}

# non-existing domain
# output "nonexisting" {
#   value = provider::iputils::lookupipv4("nonexistingdomain.xyz")
# }

output "localhost" {
  value = provider::iputils::lookupipv4("localhost")
}

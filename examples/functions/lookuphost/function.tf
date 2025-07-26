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

output "googlev4" {
  value = provider::iputils::lookuphost("142.250.203.206")
}

output "googlev6" {
  value = provider::iputils::lookuphost("2a00:1450:401b:810::200e")
}

output "one" {
  value = provider::iputils::lookuphost("1.1.1.1")
}

output "onev6" {
  value = provider::iputils::lookuphost("2606:4700:4700::1111")
}


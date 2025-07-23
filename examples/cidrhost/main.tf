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

output "host1" {
  value = provider::iputils::cidrhost("192.168.128.1/24")
}

output "host2" {
  value = provider::iputils::cidrhost("10.0.0.0/16")
}

# output "invalid_address" {
#   value = provider::iputils::cidrhost("10.0.0.256/16")
# }

# output "invalid_mask" {
#   value = provider::iputils::cidrhost("10.0.0.1/33")
# }

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

output "addr1" {
  value = provider::iputils::fromcidr("192.168.128.42/24")
}

output "addr2" {
  value = provider::iputils::fromcidr("10.0.0.0/16")
}

# output "ipv6" {
#   value = provider::iputils::fromcidr("2001:db8::1/64")
# }

# output "invalid_address" {
#   value = provider::iputils::fromcidr("10.0.0.256/16")
# }

# output "invalid_mask" {
#   value = provider::iputils::fromcidr("10.0.0.1/33")
# }

# output "missing_mask" {
#   value = provider::iputils::fromcidr("10.0.0.1")
# }

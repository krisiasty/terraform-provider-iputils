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

output "bits1" {
  value = provider::iputils::cidrbits("192.168.128.1/24")
}

output "bits2" {
  value = provider::iputils::cidrbits("10.0.0.0/16")
}

# output "invalid_address" {
#    value = provider::iputils::cidrbits("10.0.0.256/16")
# }

# output "invalid_mask" {
#   value = provider::iputils::cidrbits("10.0.0.1/33")
# }

# output "missing_mask" {
#   value = provider::iputils::cidrbits("10.0.0.1")
# }

output "ipv6" {
  value = provider::iputils::cidrbits("2001:db8::1/128")
}
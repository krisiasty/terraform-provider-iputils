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

output "mask1" {
  value = provider::iputils::cidrmask("192.168.128.1/24")
}

output "mask2" {
  value = provider::iputils::cidrmask("10.0.0.0/16")
}

# output "invalid_address" {
#    value = provider::iputils::cidrmask("10.0.0.256/16")
# }

# output "invalid_mask" {
#   value = provider::iputils::cidrmask("10.0.0.1/33")
# }

# output "missing_mask" {
#   value = provider::iputils::cidrmask("10.0.0.1")
# }

output "ipv6" {
  value = provider::iputils::cidrmask("2001:db8::1/64")
}

# output "ipv6_invalid" {
#   value = provider::iputils::cidrmask("2001:db8::1/130")
# }

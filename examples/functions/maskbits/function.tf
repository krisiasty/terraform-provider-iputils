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
  value = provider::iputils::maskbits("255.255.255.0")
}

output "mask2" {
  value = provider::iputils::maskbits("255.255.0.0")
}

output "mask3" {
  value = provider::iputils::maskbits("255.255.255.255")
}

output "mask4" {
  value = provider::iputils::maskbits("0.0.0.0")
}

# output "non_canonical_mask" {
#   value = provider::iputils::maskbits("255.0.255.0")
# }


# output "invalid_mask" {
#   value = provider::iputils::maskbits("256.0.0.0")
# }


# output "invalid_format" {
#   value = provider::iputils::maskbits("ffffff00")
# }

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tocidr function - terraform-provider-iputils"
subcategory: ""
description: |-
  
---

# function: tocidr

Convert IP address and subnet mask to CIDR notation

## Example Usage

```terraform
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

output "cidr1" {
  value = provider::iputils::tocidr("192.168.128.42","255.255.255.0")
}

output "cidr2" {
  value = provider::iputils::tocidr("10.0.0.0","255.255.0.0")
}

# output "invalid_address" {
#   value = provider::iputils::tocidr("10.0.0.256","255.255.0.0")
# }

# output "non_canonical_mask" {
#   value = provider::iputils::tocidr("10.0.0.1","255.0.255.0")
# }


# output "invalid_mask" {
#   value = provider::iputils::tocidr("10.0.0.1","256.0.0.0")
# }


# output "invalid_address2" {
#   value = provider::iputils::tocidr("10.0.0","255.255.0.0")
# }

# output "invalid_address2" {
#   value = provider::iputils::tocidr("::1","255.255.0.0")
# }
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
tocidr(address string, mask string) string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `address` (String) IP address (ipv4)
1. `mask` (String) Subnet mask (e.g. 255.255.255.0)

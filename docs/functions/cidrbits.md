---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cidrbits function - terraform-provider-iputils"
subcategory: ""
description: |-
  
---

# function: cidrbits

Return number of bits in the mask from the address in CIDR notation

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
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
cidrbits(cidr string) number
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `cidr` (String) IP address and subnet mask in CIDR notation (ipv4 or ipv6)

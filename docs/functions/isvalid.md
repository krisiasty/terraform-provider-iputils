---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "isvalid function - terraform-provider-iputils"
subcategory: ""
description: |-
  
---

# function: isvalid

Returns true if provided address is valid

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

output "addr1" {
  value = provider::iputils::isvalid("192.168.128.42")
}

output "addr2" {
  value = provider::iputils::isvalid("10.0.0.0")
}

output "addr_cidr" {
  value = provider::iputils::isvalid("10.0.0.0/16")
}

output "invalid_address" {
  value = provider::iputils::isvalid("10.0.0.256")
}


output "invalid_address2" {
  value = provider::iputils::isvalid("10.0.0")
}

output "ipv6" {
  value = provider::iputils::isvalid("::1")
}

output "ipv6_cidr" {
  value = provider::iputils::isvalid("2001:0db8:0::1/64")
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
isvalid(address string) bool
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `address` (String) IP address (ipv4 or ipv6)

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "lookupipv6 function - terraform-provider-iputils"
subcategory: ""
description: |-
  
---

# function: lookupipv6

Returns a list of IPv6 addresses returned by DNS

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

output "google" {
  value = provider::iputils::lookupipv6("google.com")
}

output "cloudflare" {
  value = provider::iputils::lookupipv6("cloudflare.com")
}

output "example" {
  value = provider::iputils::lookupipv6("example.com")
}

# non-existing domain
# output "nonexisting" {
#   value = provider::iputils::lookupipv6("nonexistingdomain.xyz")
# }

output "localhost" {
  value = provider::iputils::lookupipv6("localhost")
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
lookupipv6(host string) list of string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `host` (String) hostname to look up

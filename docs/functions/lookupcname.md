---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "lookupcname function - terraform-provider-iputils"
subcategory: ""
description: |-
  
---

# function: lookupcname

Returns canonical name (CNAME) record for specified host

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

output "ibm" {
  value = provider::iputils::lookupcname("www.ibm.com")
}

output "microsoft" {
  value = provider::iputils::lookupcname("www.microsoft.com")
}

### no cname record, should return the same name
output "example" {
  value = provider::iputils::lookupcname("example.com")
}

# non-existing domain
# output "nonexisting" {
#   value = provider::iputils::lookupcname("nonexistingdomain.xyz")
# }

output "localhost" {
  value = provider::iputils::lookupcname("localhost")
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
lookupcname(host string) string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `host` (String) hostname to look up

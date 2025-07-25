---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "urlport function - terraform-provider-iputils"
subcategory: ""
description: |-
  
---

# function: urlport

Returns the port part of the URL

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

# should return "80"
output "port1" {
  value = provider::iputils::urlport("http://192.168.128.42")
}

# should return "443"
output "port2" {
  value = provider::iputils::urlport("https://www.ibm.com/path/to/resource?query=param#fragment")
}

# should return "8080"
output "port3" {
  value = provider::iputils::urlport("http://localhost:8080")
}

# should return "8443"
output "port4" {
  value = provider::iputils::urlport("https://[2001:0db8:0::1]:8443/?param1=value1&param2=value2")
}   

# this is not a valid URL (missing scheme), will fail
# output "not_valid_url" {
#   value = provider::iputils::urlport("www.example.com/path/to/resource")
# }

# should return "21" as the default port for ftp
output "ftp" {
  value = provider::iputils::urlport("ftp://example.com/directory/file.txt")
}

# should return "2222"
output "ssh" {
  value = provider::iputils::urlport("ssh://example.com:2222")
}

# should return "22" as the default port for ssh
output "ssh_no_port" {
  value = provider::iputils::urlport("ssh://example.com")
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
urlport(url string) string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `url` (String) Valid URL to extract port number from (e.g. https://example.com:8443/path?query=1#fragment)

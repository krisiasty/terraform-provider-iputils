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

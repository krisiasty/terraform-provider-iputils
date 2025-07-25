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

output "prefix1" {
  value = provider::iputils::prefixsize("0.0.0.0/0")
}

output "prefix2" {
  value = provider::iputils::prefixsize("10.0.0.0/8")
}

output "prefix3" {
  value = provider::iputils::prefixsize("100.64.0.0/10")
}


output "prefix4" {
  value = provider::iputils::prefixsize("192.168.0.0/16")
}

output "prefix5" {
  value = provider::iputils::prefixsize("192.168.0.0/24")
}

output "prefix6" {
  value = provider::iputils::prefixsize("192.168.0.0/25")
}

output "prefix7" {
  value = provider::iputils::prefixsize("192.168.0.0/31")
}


output "prefix8" {
  value = provider::iputils::prefixsize("192.168.0.0/32")
}

### invalid prefix
# output "prefix9" {
#   value = provider::iputils::prefixsize("192.168.0.0/33")
# }



### prefix too large
# output "ipv6_prefix1" {
#   value = provider::iputils::prefixsize("2001:0db8::/32")
# }

### prefix too large
# output "ipv6_prefix2" {
#   value = provider::iputils::prefixsize("2001:0db8::/64")
# }

output "ipv6_prefix3" {
  value = provider::iputils::prefixsize("2001:0db8::/127")
}

output "ipv6_prefix4" {
  value = provider::iputils::prefixsize("2001:0db8::/128")
}

### maximum prefix size
output "ipv6_prefix5" {
  value = provider::iputils::prefixsize("2001:0db8::/66")
}

output "ipv6_prefix6" {
  value = provider::iputils::prefixsize("2001:0db8::/96")
}

### invalid prefix, not canonical

# output "ipv6_prefix7" {
#   value = provider::iputils::prefixsize("2001:0db8::1/96")
# }

output "ipv6_prefix8" {
  value = provider::iputils::prefixsize("2001:0db8::0/96")
}

### invalid prefix
# output "ipv6_prefix9" {
#    value = provider::iputils::prefixsize("2001:0db8::/129")
# }

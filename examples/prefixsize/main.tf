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
  value = provider::iputils::prefixsize("192.168.0.0/0")
}

output "prefix2" {
  value = provider::iputils::prefixsize("192.168.0.0/8")
}

output "prefix3" {
  value = provider::iputils::prefixsize("192.168.0.0/16")
}

output "prefix4" {
  value = provider::iputils::prefixsize("192.168.0.0/24")
}

output "prefix5" {
  value = provider::iputils::prefixsize("192.168.0.0/25")
}

output "prefix6" {
  value = provider::iputils::prefixsize("192.168.0.0/31")
}


output "prefix7" {
  value = provider::iputils::prefixsize("192.168.0.0/32")
}

### invalid prefix
# output "prefix8" {
#   value = provider::iputils::prefixsize("192.168.0.0/33")
# }


# output "invalid_prefix1" {
#   value = provider::iputils::prefixsize("192.168.128.42-192.168.128.41")
# }

# output "invalid_prefix2" {
#   value = provider::iputils::prefixsize("192.168.128.42-10.10.10.10")
# }

# output "invalid_prefix3" {
#   value = provider::iputils::prefixsize("192.168.128.42-192.168.128.256")
# }

# output "prefix_with_spaces" {
#   value = provider::iputils::prefixsize("192.168.128.42 - 192.168.128.253")
# }

# output "short_prefix" {
#   value = provider::iputils::prefixsize("192.168.128.42-44")
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

### invalid prefix
# output "ipv6_prefix5" {
#    value = provider::iputils::prefixsize("2001:0db8::/129")
# }

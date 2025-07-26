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

output "range1" {
  value = provider::iputils::isrange("192.168.128.42-192.168.128.42")
}

output "range2" {
  value = provider::iputils::isrange("192.168.128.42-192.168.128.254")
}

output "range3" {
  value = provider::iputils::isrange("192.168.128.0-192.168.128.42")
}

output "range4" {
  value = provider::iputils::isrange("192.168.128.1-192.168.128.255")
}

output "range5" {
  value = provider::iputils::isrange("192.168.128.0-192.168.128.255")
}

output "range6" {
  value = provider::iputils::isrange("192.168.128.0-192.168.129.255")
}

output "invalid_cidr" {
  value = provider::iputils::isrange("10.0.0.0/33")
}

output "invalid_range1" {
  value = provider::iputils::isrange("192.168.128.42-192.168.128.41")
}

output "invalid_range2" {
  value = provider::iputils::isrange("192.168.128.42-10.10.10.10")
}

output "invalid_range3" {
  value = provider::iputils::isrange("192.168.128.42-192.168.128.256")
}

output "range_with_spaces" {
  value = provider::iputils::isrange("192.168.128.42 - 192.168.128.253")
}

output "short_range" {
  value = provider::iputils::isrange("192.168.128.42-44")
}

output "invalid_address" {
  value = provider::iputils::isrange("10.0.0.256")
}


output "invalid_address2" {
  value = provider::iputils::isrange("10.0.0")
}

output "ipv6" {
  value = provider::iputils::isrange("::1")
}

output "ipv6_cidr" {
  value = provider::iputils::isrange("2001:0db8:0::1/64")
}

output "invalid_ipv6_cidr" {
  value = provider::iputils::isrange("2001:0db8:0::1/132")
}

output "invalid_ipv6_cidr2" {
  value = provider::iputils::isrange("2001:invalid::1/12")
}

output "ipv6_range1" {
  value = provider::iputils::isrange("::1-::2")
}


output "ipv6_range2" {
  value = provider::iputils::isrange("2001:0db8::1:1-2001:0db8::2:fffe")
}

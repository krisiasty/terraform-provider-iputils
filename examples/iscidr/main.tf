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

output "addr1" {
  value = provider::iputils::iscidr("192.168.128.42")
}

output "addr2" {
  value = provider::iputils::iscidr("10.0.0.0")
}

output "addr_cidr" {
  value = provider::iputils::iscidr("10.0.0.0/16")
}

output "invalid_cidr" {
  value = provider::iputils::iscidr("10.0.0.0/33")
}

output "invalid_address" {
  value = provider::iputils::iscidr("10.0.0.256")
}


output "invalid_address2" {
  value = provider::iputils::iscidr("10.0.0")
}

output "ipv6" {
  value = provider::iputils::iscidr("::1")
}

output "ipv6_cidr" {
  value = provider::iputils::iscidr("2001:0db8:0::1/64")
}

output "invalid_ipv6_cidr" {
  value = provider::iputils::iscidr("2001:0db8:0::1/132")
}

output "invalid_ipv6_cidr2" {
  value = provider::iputils::iscidr("2001:invalid::1/12")
}
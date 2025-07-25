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

output "ipv4_0" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", 0)
}


output "ipv4_1" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", 1)
}


output "ipv4_2" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", 2)
}

output "ipv4_254" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", 254)
}


output "ipv4_255" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", 255)
}

### invalid, out of bounds
# output "ipv4_256" {
#   value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", 256)
# }

output "ipv4_neg1" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", -1)
}

output "ipv4_neg2" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", -2)
}

output "ipv4_neg254" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", -254)
}

output "ipv4_neg255" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", -255)
}

output "ipv4_neg256" {
  value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", -256)
}

### invalid, out of bounds
# output "ipv4_neg257" {
#     value = provider::iputils::rangehost("192.168.128.0-192.168.128.255", -257)
# }

output "ipv6_0" {
  value = provider::iputils::rangehost("2001:db8::0-2001:db8::ffff", 0)
}

output "ipv6_1" {
  value = provider::iputils::rangehost("2001:db8::0-2001:db8::ffff", 1)
}

output "ipv6_2" {
  value = provider::iputils::rangehost("2001:db8::0-2001:db8::ffff", 2)
}  

output "ipv6_neg1" {
  value = provider::iputils::rangehost("2001:db8::0-2001:db8::1:0", -1)
}

output "ipv6_neg2" {
  value = provider::iputils::rangehost("2001:db8::0-2001:db8::1:0", -2)
}

output "ipv6_neg2_2" {
  value = provider::iputils::rangehost("2001:db8::0-2001:db8::2:0", -2)
}
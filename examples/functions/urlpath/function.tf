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

# this url does not have any path, will return an empty string
output "path1" {
  value = provider::iputils::urlpath("http://192.168.128.42")
}

# returns "/path/to/resource"; query and fragment are not considered part of the path
output "path2" {
  value = provider::iputils::urlpath("https://www.ibm.com:443/path/to/resource?query=param#fragment")
}

# returns "/dir1/dir2/dir3/"; trailing slash is included
output "path3" {
  value = provider::iputils::urlpath("http://localhost:8080/dir1/dir2/dir3/")
}

# returns "/"; root path
output "path4" {
  value = provider::iputils::urlpath("https://[2001:0db8:0::1]/?param1=value1&param2=value2")
}   

# invalid URL - missing scheme, will fail
# output "not_valid_url" {
#   value = provider::iputils::urlpath("www.example.com/path/to/resource")
# }

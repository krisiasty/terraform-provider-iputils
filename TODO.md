# List of planned functions

The list below is my initial proposals for useful functions that may be implemented by **iputils** provider.
Keep in mind that the functions may be implemented in any order, their name and arguments may change, new functions may be added or existing one removed from the list at any time.

- [X] given ip address and subnet mask, return address in cidr notation -> **tocidr**
- [X] given ip address with mask in cidr notation, return an object with separate address and netmask -> **fromcidr**
- [X] given ip address with mask in cidr notation, return ip address only -> **cidrhost** (eg. "192.168.128.10")
- [X] given ip address with mask in cidr notation, return subnet mask only -> **cidrmask** (eg. "/24")
- [X] given ip address in cidr format, return number of subnet mask bits -> **cidrbits** (e.g. 22)
- [X] given subnet mask, return cidr subnet mask -> **maskcidr** (e.g. "/24")
- [X] given subnet mask, return number of subnet mask bits -> **maskbits** (e.g. 22)
- [X] given ip address, return true if valid -> **isvalid**
- [X] given ip address in cidr format, return true if valid -> **iscidr**
- [ ] given ip address (ip only or cidr), return true if unicast -> **isunicast**
- [ ] given ip address (ip only or cidr), return true if multicast -> **ismulticast**
- [ ] given ip address (ip only or cidr), return true if private -> **isprivate** (rfc1918 + cgnat / rfc4193 for ipv6)
- [X] given ip range, return true if valid -> **isrange**
- [X] given ip range, return number of ip addresses -> **rangesize**
- [X] given ip range, return nth address from beginning or end (if negative) -> **rangehost**
- [X] given ip prefix, return number of ip addresses -> **prefixsize**
- [ ] ~~given ip prefix, return nth address from beginning or end (if negative) -> **prefixhost**~~ (already implemented by built-in **cidrhost** function)
- [ ] given ip address and netmask, return number of ip addresses -> **subnetsize**
- [ ] given ip address and netmask, return nth address from beginning or end (if negative) -> **subnethost**
- [X] given host name, lookup ip addresses -> **lookupaddr**, **lookupipv4**, **lookupipv6**
- [X] given IP address, lookup hostnames -> **lookuphost**
- [X] given hostname, lookup canonical name -> **lookupcname**
- [X] given hostname, lookup TXT records -> **lookuptxt**
- [ ] given hostname or ip address with port, return host portion only -> **hostonly**
- [ ] given hostname or ip address with port, return port number only -> **portonly**
- [X] given url, return host portion only -> **urlhost**
- [X] given url, return port only -> **urlport**
- [ ] given url, return path only -> **urlpath**
- [ ] given url, return query params only -> **urlquery**
- [ ] given url, return escaped url -> **urlescape**
- [ ] given ip prefix, offset and count, return count ip addresses starting from offset - **prefixhostlist**
- [ ] given ip address with subnet mask, offset and count, return count ip addresses starting from offset - **subnethostlist**
- [ ] given ip range, offset and count, return count ip addresses starting from offset - **rangehostlist**
- [ ] given ip prefix, return IP range -> **prefix2range**
- [ ] given ip address and subnet, return IP range -> **subnet2range**
- [ ] given ip prefix and ip address, return true if address is from the prefix -> **prefixcontains**
- [ ] given ip address, netmask and address, return true if address is from specified subnet -> **subnetcontains**
- [ ] given ip range and address, return true if address is from specified range -> **rangecontains**

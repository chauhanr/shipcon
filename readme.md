# Shipcon Application

#### mdns (multicast DNS)
mdns is a protocol resolves hostname to IP address within a small network that does not include a local name server.
It is a zero configuration service, using the same programming interface, packet format and operating semantics as
the unicast domain system

The reason for using mdns while testing locally is to avoid the use of Consul or etcd which is used for discovery
in production.


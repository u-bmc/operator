// SPDX-License-Identifier: BSD-3-Clause

package dhcp

import (
	"github.com/insomniacslk/dhcp/dhcpv4/client4"
	"github.com/insomniacslk/dhcp/dhcpv6/client6"
)

// IPVersion is a type used to distinguish between IPv4 and IPv6.
type IPVersion uint

const (
	// IPv4 represents the IP version 4.
	IPv4 IPVersion = iota
	// IPv6 represents the IP version 6.
	IPv6
)

// GetLease is a function that retrieves the lease for the specified IP version and interface.
// It returns a list of string representations of the packets exchanged during the DHCP transaction and an error if any.
func GetLease(version IPVersion, iface string) ([]string, error) {
	// If the IP version is IPv4, get the lease for IPv4.
	if version == IPv4 {
		return getLeasev4(iface)
	}

	// Otherwise, get the lease for IPv6.
	return getLeasev6(iface)
}

func getLeasev4(iface string) ([]string, error) {
	// NewClient sets up a new DHCPv4 client with default values
	// for read and write timeouts, for destination address and listening
	// address
	client := client4.NewClient()

	// Exchange runs a Solicit-Advertise-Request-Reply transaction on the
	// specified network interface, and returns a list of DHCPv4 packets
	// (a "conversation") and an error if any. Notice that Exchange may
	// return a non-empty packet list even if there is an error. This is
	// intended, because the transaction may fail at any point, and we
	// still want to know what packets were exchanged until then.
	// A default Solicit packet will be used during the "conversation",
	// which can be manipulated by using modifiers.
	conversation, err := client.Exchange(iface)

	packets := make([]string, 0, len(conversation))
	// Summary() prints a verbose representation of the exchanged packets.
	for i, packet := range conversation {
		packets[i] = packet.Summary()
	}

	return packets, err
}

func getLeasev6(iface string) ([]string, error) {
	// NewClient sets up a new DHCPv6 client with default values
	// for read and write timeouts, for destination address and listening
	// address
	client := client6.NewClient()

	// Exchange runs a Solicit-Advertise-Request-Reply transaction on the
	// specified network interface, and returns a list of DHCPv6 packets
	// (a "conversation") and an error if any. Notice that Exchange may
	// return a non-empty packet list even if there is an error. This is
	// intended, because the transaction may fail at any point, and we
	// still want to know what packets were exchanged until then.
	// A default Solicit packet will be used during the "conversation",
	// which can be manipulated by using modifiers.
	conversation, err := client.Exchange(iface)

	packets := make([]string, 0, len(conversation))
	// Summary() prints a verbose representation of the exchanged packets.
	for i, packet := range conversation {
		packets[i] = packet.Summary()
	}

	return packets, err
}

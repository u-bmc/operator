// SPDX-License-Identifier: BSD-3-Clause

package iface

import (
	"bytes"
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

type Mode int

const (
	ModeDHCP Mode = iota
	ModeStatic
)

type Iface struct {
	Name    string
	MAC     net.HardwareAddr
	Mode    Mode
	IPv4    []net.IP
	IPv6    []net.IP
	Gateway []net.IP
	DNS     []net.IP
}

func (i Iface) AddIP(cidr string) error {
	l, err := netlink.LinkByName(i.Name)
	if err != nil {
		return fmt.Errorf("unable to get interface %s: %w", i.Name, err)
	}

	addr, err := netlink.ParseAddr(cidr)
	if err != nil {
		return fmt.Errorf("unable to parse %q: %w", cidr, err)
	}

	h, err := netlink.NewHandle(unix.NETLINK_ROUTE)
	if err != nil {
		return fmt.Errorf("can't get handle: %w", err)
	}
	defer h.Delete()

	if err := h.AddrReplace(l, addr); err != nil {
		return fmt.Errorf("unable to replace addr with %v: %w", addr, err)
	}

	return nil
}

func (i Iface) SetLinkUp() error {
	l, err := netlink.LinkByName(i.Name)
	if err != nil {
		return fmt.Errorf("unable to get interface %s: %w", i.Name, err)
	}

	h, err := netlink.NewHandle(unix.NETLINK_ROUTE)
	if err != nil {
		return fmt.Errorf("can't get handle: %w", err)
	}
	defer h.Delete()

	if err := h.LinkSetUp(l); err != nil {
		return fmt.Errorf("failed to set link %q up: %w", i.Name, err)
	}

	return nil
}

func (i Iface) SetLinkDown() error {
	l, err := netlink.LinkByName(i.Name)
	if err != nil {
		return fmt.Errorf("unable to get interface %s: %w", i.Name, err)
	}

	h, err := netlink.NewHandle(unix.NETLINK_ROUTE)
	if err != nil {
		return fmt.Errorf("can't get handle: %w", err)
	}
	defer h.Delete()

	if err := h.LinkSetDown(l); err != nil {
		return fmt.Errorf("failed to set link %q down: %w", i.Name, err)
	}

	return nil
}

func IsLinkLocalMAC(addr []byte, hw []byte) bool {
	chw := append([]byte{0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, hw[0], hw[1], hw[2], 0xff, 0xfe, hw[3], hw[4], hw[5])
	chw[8] ^= 0x2
	return bytes.Equal(chw, addr)
}

// func ipv6LinkFixer(iface string) {
// 	// Verify that the link local address based on the MAC address is present
// 	// every 1 second. If it's not, reset the interface (and in the future
// 	// of course reset DHCP etc.).
// 	h, err := netlink.NewHandle(unix.NETLINK_ROUTE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer h.Delete()

// 	sleep := 1 * time.Second
// 	for {
// 		time.Sleep(sleep)
// 		sleep = 1 * time.Second
// 		l, err := netlink.LinkByName(iface)
// 		if err != nil {
// 			panic(err)
// 		}
// 		a := l.Attrs()
// 		// OperState does not work with all interfaces, so use flags
// 		// NC-SI reports 'unknown' for all states (and so does loopback interfaces, FWIW)
// 		if a.Flags&net.FlagUp == 0 {
// 			continue
// 		}
// 		addrs, err := h.AddrList(l, netlink.FAMILY_V6)
// 		if err != nil {
// 			panic(err)
// 		}
// 		found := false
// 		for _, addr := range addrs {
// 			if addr.IP.IsLinkLocalUnicast() && isLinkLocalForMAC(addr.IP.To16(), a.HardwareAddr) {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			err = setLinkDown(iface)
// 			if err != nil {
// 				panic(err)
// 			}
// 			err = setLinkUp(iface)
// 			if err != nil {
// 				panic(err)
// 			}
// 			// Back off 10 seconds before trying again to avoid flapping too much
// 			sleep = 10 * time.Second
// 		}
// 	}
// }

// func StartNetwork(config *proto.Network) (*network, error) {
// 	if config == nil {
// 		panic("StartNetwork called with nil config")
// 	}

// 	// Fun story: if you don't have both IPv4 and IPv6 loopback configured
// 	// golang binaries will not bind to :: but to 0.0.0.0 instead.
// 	// Isn't that surprising?
// 	if err := addIp("127.0.0.1/8", "lo"); err != nil {
// 		return nil, err
// 	}
// 	if err := addIp("::1/32", "lo"); err != nil {
// 		return nil, err
// 	}
// 	if err := setLinkUp("lo"); err != nil {
// 		return nil, err
// 	}

// 	iface := "eth0"

// 	// TODO(bluecmd): Set ipv4/ipv6 objects to remember the host addresses
// 	if config.Vlan != 0 {
// 		log.Infof("TODO: Interface was configured to use VLAN but that's not implemented yet")
// 	}

// 	// TODO use insomniacslk/dhcp instead of external dhclient
// 	// dhclient := exec.Command("dhclient")
// 	// err := dhclient.Run()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// If the MAC address changes on the interface the interface needs to be
// 	// taken down and up again in order for all IPv6 addresses and things to be
// 	// refreshed. MAC address changes happens when NC-SI reads the correct
// 	// MAC address from the adapter, or a controller hotswap potentially.
// 	go ipv6LinkFixer(iface)

// 	if config.Ipv4Address != "" {
// 		if err := addIp(config.Ipv4Address, iface); err != nil {
// 			log.Errorf("Error adding IPv4 %s to interface %s: %v", config.Ipv4Address, iface, err)
// 		}
// 	}
// 	if config.Ipv6Address != "" {
// 		if err := addIp(config.Ipv6Address, iface); err != nil {
// 			log.Errorf("Error adding IPv6 %s to interface %s: %v", config.Ipv6Address, iface, err)
// 		}
// 	}

// 	if len(config.Ipv4Route)+len(config.Ipv6Route) > 0 {
// 		log.Infof("TODO: IP routes are configured but not supported yet")
// 	}

// 	go func() {
// 		c := make(chan *dns.RDNSSOption)
// 		go dns.RDNSS(c)
// 		for r := range c {
// 			log.Infof("TODO: got RDNSS %v", r)
// 		}
// 	}()

// 	// When we exit this function we must have received a hostname or otherwise
// 	// had one configured. The rest of the startup flow depends on it.

// 	// TODO(bluecmd): Read hostname from config file or DHCP, don't have any default
// 	fqdn := "ubmc.local"
// 	if config.Hostname != "" {
// 		fqdn = config.Hostname
// 	}
// 	err := unix.Sethostname([]byte(fqdn))
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	return &network{fqdn: fqdn}, nil
// }

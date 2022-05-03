package utils

import (
	"errors"
	"net"
)

type IPMatcher struct {
	IP     net.IP
	SubNet *net.IPNet
}
type IPMatchers []*IPMatcher

func NewIPMatcher(ipStr string) (*IPMatcher, error) {
	ip, subNet, err := net.ParseCIDR(ipStr)
	if err != nil {
		ip = net.ParseIP(ipStr)
		if ip == nil {
			return nil, errors.New("invalid IP: " + ipStr)
		}
	}
	return &IPMatcher{ip, subNet}, nil
}

func (m IPMatcher) Match(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return m.IP.Equal(ip) || m.SubNet != nil && m.SubNet.Contains(ip)
}

func NewIPMatchers(ips []string) (list IPMatchers, err error) {
	for _, ipStr := range ips {
		var m *IPMatcher
		m, err = NewIPMatcher(ipStr)
		if err != nil {
			return
		}
		list = append(list, m)
	}
	return
}

func IPContains(ipMatchers []*IPMatcher, ip string) bool {
	for _, m := range ipMatchers {
		if m.Match(ip) {
			return true
		}
	}
	return false
}

func (ms IPMatchers) Match(ip string) bool {
	return IPContains(ms, ip)
}

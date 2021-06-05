package main

import (
	"github.com/wangtuanjie/ip17mon"
)

var datPath = "17monipdb.dat"

type ipipObj struct {
	Locator ip17mon.Locator
	path    string
}

type ImpIpip interface {
	// NewIpipObj() (ip17mon.Locator, error)
	NewIpipObj() error
	QueryIP(string) ([]string, error)
}

func LoadIpip(path string) ImpIpip {
	return &ipipObj{
		path: path,
	}
}

func (i *ipipObj) NewIpipObj() error {
	loc, err := ip17mon.New(i.path)
	if err != nil {
		return err
	}

	i.Locator = loc
	return nil
}

func (i *ipipObj) QueryIP(ip string) ([]string, error) {
	loc, err := i.Locator.Find(ip)
	if err != nil {
		return nil, err
	}
	return []string{
		loc.Country,
		loc.Region,
		loc.City,
		loc.Isp}, nil
}

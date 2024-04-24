package main

import "fmt"

type UserDetails interface {
	GetName() string
	GetAddress() string
}

type MacOsUsers struct {
	name, address string
}

func (macOsUsers MacOsUsers) GetName() string {
	return macOsUsers.name
}

func (macOsUsers MacOsUsers) GetAddress() string {
	return macOsUsers.address
}

type LinuxUsers struct {
	name, address string
}

func (linuxUser LinuxUsers) GetName() string {
	return linuxUser.name
}

func (linuxUser LinuxUsers) GetAddress() string {
	return linuxUser.address
}

func printUserDetails(userDetails UserDetails) {
	fmt.Println(userDetails.GetName())
	fmt.Println(userDetails.GetAddress())
}

func main() {
	macUser := MacOsUsers{
		name:    "Melody Lhaksani",
		address: "Jakarta, Indonesia",
	}

	linuxUser := LinuxUsers{
		name:    "Aditya Prabowo",
		address: "Surabaya, Indonesia",
	}

	printUserDetails(macUser)
	printUserDetails(linuxUser)
}

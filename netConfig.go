package main

import (
	"fmt"
	"net"
)

func main(){
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n",i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("By name:",byName)
		addr, err := byName.Addrs()
		fmt.Println("byName.Addrs():",addr)
		for k, v:= range addr{
			fmt.Printf("Interface addr #%v: %v\n",k,v.String())
		}
		fmt.Println()
	}

}

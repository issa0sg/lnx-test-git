package main

import (
        "fmt"
        "net"
	"os/exec"
)

func main(){
	cmd := exec.Command("touch","test")
	stdout, _ := cmd.Output()
	fmt.Print(string(stdout))
	//cmd.Run()
}

func cmd_func(){
	cmd := exec.Command("netstat","-nr")
	stdout, _ := cmd.Output()
	fmt.Print(string(stdout))
}

func netCapabilities(){
        interfaces, err := net.Interfaces()
        if err != nil {
                fmt.Print(err)
                return
        }
        for _,i := range interfaces{
                fmt.Printf("Name: %v\n",i.Name)
                fmt.Println("Interface Flags:",i.Flags.String())
                fmt.Println("Interface MTU:",i.MTU)
                fmt.Println("Interface Hardware Address:",i.HardwareAddr)
                fmt.Println()
        }
}

func netConfig(){
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

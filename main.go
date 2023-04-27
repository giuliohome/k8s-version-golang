package main

import (
	"fmt"
	"strconv"

	"k8s.io/client-go/discovery"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	cfg := config.GetConfigOrDie()
	discClient, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		panic(err)
	}
	serverVersion, err := discClient.ServerVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println("kubernetes " + serverVersion.Major + "." + serverVersion.Minor)
	major, _ := strconv.Atoi(serverVersion.Major)
	minor, _ := strconv.Atoi(serverVersion.Minor)
	if major == 1 && minor >= 25 {
		fmt.Println("Yeah!")
	} else {
		fmt.Println("Not upgraded!")
	}
}

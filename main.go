package main

import "fmt"

var (
	MAJOR = 0
	MINOR = 1
	REVISION = 2201031052
)

func version() string {
	return fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, REVISION)
}

func main() {
	fmt.Println("dhcpdadmd", version())
}

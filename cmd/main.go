package main

import (
	"fmt"
	"os"
	_ "project/common/init"
)

func main() {
	os.Args = os.Args[1:]
	fmt.Println(os.Args)

}

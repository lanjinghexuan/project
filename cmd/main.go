package main

import (
	"fmt"
	_ "github.com/lanjinghexuan/project/common/init"
	"os"
)

func main() {
	os.Args = os.Args[1:]
	fmt.Println(os.Args)

}

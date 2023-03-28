package main

import (
	"flag"
	"fmt"
)

var (
	cliName   = flag.String("name", "nick", "Input Your Name")
	cliAge    = flag.Int("age", 28, "Input Your Age")
	CliGender = flag.String("gender", "maile", "Input Your Gender")
	cliFlag   int
)

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
}
func main() {
	Init()
	flag.Parse()

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *CliGender)
	fmt.Println("flagname=", cliFlag)
}

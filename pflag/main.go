// package main

// import (
// 	"fmt"

// 	"github.com/spf13/pflag"
// )

// type host struct {
// 	value string
// }

// func (h *host) String() string {
// 	return h.value
// }

// func (h *host) Set(v string) error {
// 	h.value = v
// 	return nil
// }

// func (h *host) Type() string {
// 	return "host"
// }

// func main() {
// 	// name / defalut value / usage string
// 	var ip *int = pflag.Int("ip", 1234, "help message for ip")

// 	var port int
// 	pflag.IntVar(&port, "port", 8080, "help message for port")

// 	var h host
// 	pflag.Var(&h, "host", "help message for host")

// 	// 解析命令行参数
// 	pflag.Parse()

// 	fmt.Printf("ip: %d\n", *ip)
// 	fmt.Printf("port: %d\n", port)
// 	fmt.Printf("host: %+v\n", h)

//		fmt.Printf("NFlag: %v\n", pflag.NFlag()) // 返回已设置的命令行标志个数
//		fmt.Printf("NArg: %v\n", pflag.NArg())   // 返回处理完标志后剩余的参数个数
//		fmt.Printf("Args: %v\n", pflag.Args())   // 返回处理完标志后剩余的参数列表
//		fmt.Printf("Arg(1): %v\n", pflag.Arg(1)) // 返回处理完标志后剩余的参数列表中第 i 项
//	}
package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type host struct {
	value string
}

func (h *host) String() string {
	return h.value
}

func (h *host) Set(v string) error {
	h.value = v
	return nil
}

func (h *host) Type() string {
	return "host"
}

func main() {
	flagset := pflag.NewFlagSet("test", pflag.ExitOnError)

	var ip = flagset.IntP("ip", "i", 1234, "help message for ip")

	var boolVar bool
	flagset.BoolVarP(&boolVar, "boolVar", "b", true, "help message for boolVar")

	var h host
	flagset.VarP(&h, "host", "H", "help message for host")

	flagset.SortFlags = false

	flagset.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("boolVar: %t\n", boolVar)
	fmt.Printf("host: %+v\n", h)

	i, err := flagset.GetInt("ip")
	fmt.Printf("i: %d, err: %v\n", i, err)
}

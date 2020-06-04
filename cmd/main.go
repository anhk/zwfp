package main

import (
	"flag"
	"fmt"
	"github.com/anhk/zwfp"
)

var (
	embed = flag.String("d", "", "需要解码的数据")

	data = flag.String("i", "", "需要合成的数据")
	key  = flag.String("k", "", "需要注入的隐私")
)

func main() {
	flag.Parse()
	fmt.Println(*embed)

	if *embed != "" {
		data, key, err := zwfp.Extract(*embed)
		if err != nil {
			panic(err)
		}
		fmt.Println("你的数据：", data)
		fmt.Println("你的隐私：", key)
		return
	}

	if *data == "" || *key == "" {
		fmt.Println("错误的输入")
		flag.Usage()
		return
	}

	result, err := zwfp.Embed(*data, *key)
	if err != nil {
		panic(err)
	}
	fmt.Println("合成数据：", result)
}

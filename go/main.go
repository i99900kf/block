package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"example.com/m/BLC"
)

func main() {
	// 5.检测pow
	//1.创建一个big对象 0000000.....00001
	target := big.NewInt(1)
	fmt.Printf("0x%x\n", target) //0x1

	//2.左移256-bits位
	target = target.Lsh(target, 256-BLC.TargetBit)

	fmt.Printf("0x%x\n", target) //61
	//61位：0x1000000000000000000000000000000000000000000000000000000000000
	//64位：0x0001000000000000000000000000000000000000000000000000000000000000

	s1 := "HelloWorld"
	hash := sha256.Sum256([]byte(s1))
	fmt.Printf("0x%x\n", hash)

}

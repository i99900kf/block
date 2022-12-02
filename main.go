package main

import (
	"block/blockchain2"
	"fmt"
	"strconv"
)

/*
 * @Author: YJH
 * @Date: 2022-11-29 18:35:50
 * @LastEditors: YJH
 * @LastEditTime: 2022-12-02 10:34:28
 * @FilePath: \go-crmf:\docker\block\main.go
 * @Description:
 *
 * Copyright (c) 2022 by 不见不散, All Rights Reserved.
 */
func main() {
	//bc:=blockchain1.NewBlockchain()/第一版：通过包名称来调用
	bc := blockchain2.NewBlockchain() //第二版：通过包名称来调用
	bc.AddBlock("Send 1 BTC to Ivan.")
	bc.AddBlock("Send 2 more BTC to Ivan.")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("Nonce:%d八n", block.Nonce) //第二版
		//验证pow:第二版
		pow := blockchain2.NewProofOfWork(block)
		fmt.Printf("POW:%s\n", strconv.FormatBool(pow.Validate()))
		//打印空行：自动换行(Printf不能自动换行)
		fmt.Println()
	}
}

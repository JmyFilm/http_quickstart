package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"math/rand"
)

var sno *snowflake.Node
var no *snowflake.Node

func init() {
	sno, _ = snowflake.NewNode(int64(rand.Intn(1024))) // Single deployment
}

func InitNo(node int64) {
	var err error
	no, err = snowflake.NewNode(node)
	if err != nil {
		log.Fatal("snowflake ERROR ", err)
	}
}

func SnowFlakeSId() string {
	return sno.Generate().Base58() // The length of Base58 is 11
}

func SnowFlakeId() string {
	return no.Generate().Base58() // The length of Base58 is 11
}

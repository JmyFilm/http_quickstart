package utils

import (
	crand "crypto/rand"
	"edit-your-project-name/config"
	"encoding/binary"
	"github.com/bwmarrin/snowflake"
	"log"
	mrand "math/rand"
)

var No *snowflake.Node

func InitNo() {
	var err error
	No, err = snowflake.NewNode(config.App.NodeId)
	log.Fatal("snowflake ERROR", err)
}

func GetRandNum() uint32 {
	var seed int64
	_ = binary.Read(crand.Reader, binary.BigEndian, &seed)
	r := mrand.New(mrand.NewSource(seed))
	return r.Uint32()
}

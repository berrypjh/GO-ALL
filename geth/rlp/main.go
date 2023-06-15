package main

import (
	"bytes"
	"log"

	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	// 인코딩할 2차원 배열 문자열
	val := [][]string{
		{"asdf", "qwer", "zxcv"},
		{"asdf", "qwer", "zxcv"},
		{"asdf", "qwer", "zxcv"},
		{"asdf", "qwer", "zxcv"},
		{"asdf", "qwer", "zxcv"},
		{"asdf", "qwer", "zxcv"},
		{"asdf", "qwer", "zxcv"},
	}

	b := new(bytes.Buffer)
	if err := rlp.Encode(b, val); err != nil {
		log.Fatalf("Failed to encode %v to rlp.", val)
	}

	// 다시 2차원 배열 문자열로 디코딩 합니다.
	var decoded [][]string
	// rlp.Decode(b.Bytes, &decoded)
	rlp.Decode(bytes.NewReader(b.Bytes()), &decoded)
	log.Printf("%v", decoded) // [[asdf qwer zxcv] [asdf qwer zxcv] [asdf qwer zxcv]]
}

package main

import (
	"fmt"
	"log"
	"encoding/gob"
	"bytes"
)

type P struct {
	X, Y, Z int
	Name string
	Tag []string
	Attr map[string]string
}

type Q struct {
	X, Y *int32
	Name string
	Tag []string
	Attr map[string]string
}

func main()  {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(P{3, 4, 5, "MeatTaro", []string{"AA", "BB", "CC"}, map[string]string{"webiste": "http://xueyuanjun.com"}})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}, tag: %v, Attr: %v\n", q.Name, *q.X, *q.Y, q.Tag, q.Attr)
}
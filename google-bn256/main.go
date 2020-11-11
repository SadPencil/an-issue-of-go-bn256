package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bn256"
)

func main() {
	var err error
	var success bool

	b, err := hex.DecodeString("3977594841811ccef002cc6e764766e78e61d55c669e645f5f5d8a31d974c6d921cc1cafe7861c902b7c0eede9a19f15622ea81dfc699d4136b231bb5fa17fd1")
	if err != nil {
		panic(err)
	}

	p := &bn256.G1{}
	_, success = p.Unmarshal(b)
	if !success {
		panic(fmt.Errorf("unmarshal failed (point not on curve)"))
	}

	fmt.Println("p", p.String())
	fmt.Println("p", hex.EncodeToString(p.Marshal()))

	q := &bn256.G1{}
	_, success = q.Unmarshal(b)
	if !success {
		panic(fmt.Errorf("unmarshal failed (point not on curve)"))
	}

	fmt.Println("q", q.String())
	fmt.Println("q", hex.EncodeToString(q.Marshal()))

	// note: the address of p and q are NOT equal.
	// Thus, in the following code, invoking Add() method doesn't break the rule
	if p == q {
		panic(fmt.Errorf("impossible"))
	}

	// let r = p + q
	r := &bn256.G1{}
	r.Add(p, q)
	fmt.Println("r=p+q", r.String())
	fmt.Println("r=p+q", hex.EncodeToString(r.Marshal()))

	{
		var t bn256.G1
		_, success = t.Unmarshal(r.Marshal())
		if !success {
			panic(fmt.Errorf("unmarshal failed (point not on curve)"))
		}
		fmt.Println("OK")
	}

	// let p = p + q
	p.Add(p, q)
	fmt.Println("p=p+q", p.String())
	fmt.Println("p=p+q", hex.EncodeToString(p.Marshal()))

	{
		var t bn256.G1
		_, success = t.Unmarshal(p.Marshal())
		if !success {
			panic(fmt.Errorf("unmarshal failed (point not on curve)"))
		}
		fmt.Println("OK")
	}

}

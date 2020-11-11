package main

import (
	"encoding/hex"
	"fmt"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/google"
)

func main() {
	b, err := hex.DecodeString("04b312e56e4d8de6ff244f0efd84192c76dc29de68d2014c030a785a34345798148afcd94afc147b069c9a6b68dcc6e9bb2d3b88985ba9d50f59220ef30cc41d")
	if err != nil {
		panic(err)
	}

	p := &bn256.G1{}
	_, err = p.Unmarshal(b)
	if err != nil {
		panic(err)
	}

	fmt.Println("p", p.String())
	fmt.Println("p", hex.EncodeToString(p.Marshal()))

	q := &bn256.G1{}
	_, err = q.Unmarshal(b)
	if err != nil {
		panic(err)
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
		_, err = t.Unmarshal(r.Marshal())
		if err != nil {
			panic(err)
		}
		fmt.Println("OK")
	}

	// let p = p + q
	p.Add(p, q)
	fmt.Println("p=p+q", p.String())
	fmt.Println("p=p+q", hex.EncodeToString(p.Marshal()))

	{
		var t bn256.G1
		_, err = t.Unmarshal(p.Marshal())
		if err != nil {
			panic(err)
		}
		fmt.Println("OK")
	}

}

package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String()string {
	return fmt.Sprintf("Hash of %s is %s",p.Path,p.Hash)
}

type PairWithLength struct {
	Pair   // we have a field without a name but has a type. and it has embedded Pair fields, so now PairWithLenght has promoted fields of pair
	Length int
}


func main() {
	p := Pair{"/user", "0xsfhdf"}
	pl := PairWithLength{p, 5}
	fmt.Println(p)
	fmt.Println(pl) // function String is not associated with the type PairWithLenght but still we are getting the formatted string as in Pair, 
	// this is due to the fact tht not only fields gets promoted but the methods associated also get promoted.
	//NOTE : if the type PairWithLenght has its own function String then its going to use that otherwise it will use the promoted one from pair
}
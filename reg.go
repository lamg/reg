package main

import (
	. "regexp"
	. "fmt"
	. "os"
)

func main(){
	if len(Args) != 3{
		Fprintf(Stderr, "usage: %s regexp text\n",Args[0])
		Exit(1)
	}
	if r, e := Compile(Args[1]); e != nil{
		Fprintf(Stderr, "Bad regexp %s\n",Args[1])
		Exit(1)
	}else{
		if ind := r.FindStringSubmatchIndex(Args[2]); ind != nil {
			res := make([]string, len(ind)>>1)
			for i := 0; i < len(ind)>>1; i++ {
				res[i] = Args[2][ind[i<<1]:ind[i<<1+1]]
			}
			Printf("%v\n",res)
			Exit(0)
		}
		Exit(1)
	}
}

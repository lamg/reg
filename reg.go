// Copyright 2013 Luis Ángel Méndez Gort. All rigths reserved.
// Use of this code is governed by GPLv3 license that can be 
// found in the LICENSE file.

package main

import (
	. "regexp"
	. "fmt"
	. "os"
)

var (
	help = `
Description:

	Reg matches regexp against text
if there is no match exit status is 1 and if
regexp fails to compile an error message is 
shown matches exit status is 0 and matched 
text and matches are shown.
`
	usage = "usage: %s regexp text\n"
)

func main(){
	if len(Args) != 3{
		Fprintf(Stderr, usage,Args[0])
		Fprintf(Stderr, help)
		Exit(1)
	}
	if r, e := Compile(Args[1]); e != nil{
		Fprintf(Stderr, "Bad regexp %s\n",Args[1])
		Exit(1)
	}else{
		if ind := r.FindStringSubmatchIndex(Args[2]); ind != nil {
			res := make([]string, len(ind)>>1)
			var s int
			var f string
			for i := 0; i < len(ind)>>1; i++ {
				s = i<<1
				if i == 0 {
					f = "all:"
				}else{
					f = Sprintf("sub%d:",i)
				}
				if ind[s] == -1 {
					res[i] = f+"<nil>"
				}else{
					res[i] = f+Args[2][ind[s]:ind[s+1]]
				}
			}
			Printf("%v\n",res)
			Exit(0)
		}
		Exit(1)
	}
}

Reg
===

Description
-----------
	Reg is a simple program that receives to arguments 
(regexp, text) where regexp is a regular expression as 
understood by the Go regexp library default syntax and text 
is a string. Reg if regexp is recognized in text prints the list
of submatches and exits with 0, otherwise prints nothing and 
exits with 1. If regexp is malformed prints a message to 
stderr and exits with 1.

Install instructions
--------------------
- It was developed using Go 1.0.1. You should have a compatible 
  version.
- Fetch: go get github.com/lamg/reg.
- Go install: go install reg.

Terms of use
------------
Copyright 2013 Luis Ángel Méndez Gort. All rights reserved.
Use of this program is governed by GPLv3 license found at
the LICENSE file.

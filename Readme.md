						Reg
						===

Description
-----------
	Reg is a simple program that receives to arguments 
(regexp, text) where regexp is a regular expression as 
understood by the Go regexp library standard syntax and text 
is a string. Reg if regexp is recognized in text prints the list
of submatches and exits with 0, otherwise prints nothing and 
exits with 1. If regexp is malformed prints a message to 
stderr and exits with 1.

Install
-------
- It was developed using Go 1.0. You should have a compatible 
  version.
- Fetch: go get github.com/lamg/reg.
- Go install: go install reg.

License
-------
GPLv3.

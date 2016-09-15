This directory is to house the code for the go server.



To run the web server:
export GOPATH="/Path/To/Project/cse_cit_480/goserver"

Then, go get the PostgreSQL connecter we use lib/pq:
go get github.com/lib/pq

Get the mux router:
go get github.com/gorilla/mux

Then run the main program (assuming it isn't already compiled)
go run src/treview.com/bloom/main.go

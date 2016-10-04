# Go Server
This directory is to house the code for the go server.

## Running the project
To run the web server(this is already configured on droplet):
export GOPATH="/Path/To/Project/cse_cit_480/goserver"

Then, go get the PostgreSQL connecter we use lib/pq:
go get github.com/lib/pq

Get the mux router:
go get github.com/gorilla/mux

Then run the main program (assuming it isn't already compiled)
cd goserver/src
go run treview.com/bloom/main.go treview.com/bloom/routes.go

## Compiling the project
To compile the project, follow the rules to run the project
the only change is that the last line you should run:
go build treview.com/bloom/main.go treview.com/bloom/routes.go

Then to run the compiled exe, type:
./main
## Editing the project
Before committing your code, you must run the 'go fmt' command on your go files. 

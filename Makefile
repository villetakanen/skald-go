default : build

update_references:
	go get -v ./...

clean:
	rm -rf cmd

test: 
	go test src

build: clean
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o cmd/skald src/skald.go

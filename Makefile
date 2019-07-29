default : build

update_references:
	go get -v ./...

clean:
	rm -rf cmd

test: 
	go test github.com/vitku/skald/... 

build: clean
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o cmd/skald service/*

serve:
	cd cmd && ./skald &

restart: 
	pkill skald && cd cmd && ./skald &

kill:
	pkill skald

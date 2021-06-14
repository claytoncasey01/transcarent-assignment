export APP=transcarent-assignment
export LDFLAGS="-w -s"

build:
	go build .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run .

test: 
	go test -v ./...

container:
	docker build -t transcarent-assignment .

run-container:
	docker run --rm -p 8080:8080 -it transcarent-assignment

.PHONY: build run build-static test container
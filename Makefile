BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

vendor:
	@dep ensure -v

engine: vendor
	go build -o ${BINARY} main.go
	
start: vendor
	go run main.go

proto:
	protoc --go_out=plugins=grpc:. *.proto

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t afandy/mdm_dtschema .

run:
	docker-compose up -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
LDFLAGS=-ldflags "-X github.com/marioharper/commuter/directions.MAPS_API_KEY=$(MAPS_API_KEY) -X github.com/marioharper/commuter/weather.DARK_SKY_API_KEY=$(DARK_SKY_API_KEY)"

build: 
	@gox ${LDFLAGS} -os="linux darwin windows openbsd" -output="./dist/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...
.PHONY: build 

install:
	go install ${LDFLAGS} ./...
.PHONY: install

run:
	go run ${LDFLAGS} main.go
.PHONY: run

clean:
	go clean
.PHONY: clean
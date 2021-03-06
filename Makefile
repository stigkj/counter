sources = main.go postgrescounter.go

dist/counter-osx: $(sources)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/counter-osx

dist/counter-linux: $(sources)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/counter-linux

.PHONY: build release clean
build: dist/counter-osx dist/counter-linux

.PHONY: run
run:
	go run $(sources)


release: build
	release.sh counter $(VERSION) dist/*

packages = \
	github.com/gorilla/context \
	github.com/gorilla/handlers \
	github.com/gorilla/mux \
	github.com/lib/pq \
	gopkg.in/bsm/ratelimit.v1 \
	gopkg.in/bufio.v1 \
	gopkg.in/mgo.v2 \
	gopkg.in/redis.v3 \
	gopkg.in/unrolled/render.v1

GIT_HASH = $(shell git show-ref --hash=8 HEAD)

install:
	go get $(packages)

docker-build: dist/counter-linux
	docker build -t stigkj/counter .

docker-run:
	docker run -p 3000:80 -it --rm --name counter stigkj/counter

docker-push:
	docker tag stigkj/counter stigkj/counter:$(GIT_HASH)
	docker push stigkj/counter:$(GIT_HASH)
	docker push stigkj/counter:latest

clean :
	-rm -r dist

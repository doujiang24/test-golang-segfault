

build-so:
	sudo docker run --rm -v $(PWD):/go/src/go-filter -w /go/src/go-filter \
                -e GOPROXY=https://goproxy.cn \
                golang:1.22.9-bullseye \
                make local-build

local-build:
	go build -v -o libgolang.so -buildmode=c-shared -buildvcs=false .

run:
	sudo docker run --rm -v $(PWD)/envoy.yaml:/etc/envoy/envoy.yaml \
                -v $(PWD)/libgolang.so:/etc/envoy/libgolang.so \
                -p 8089:8089 \
                envoyproxy/envoy:contrib-v1.32.1 \
                envoy -c /etc/envoy/envoy.yaml


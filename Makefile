GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod

export GO111MODULE = on
export GOPROXY=https://goproxy.cn

all:build

build: go-mod-cache
	$(GOBUILD) -o build/disk-info ./cmd

clean:
	rm -rf build/

go-mod-cache:
	$(GOMOD) download

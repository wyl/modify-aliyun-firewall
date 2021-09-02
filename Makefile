
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./go-modify-aliyun-firewall-linux src/main.go

build: build-linux
	go build -o ./go-modify-aliyun-firewall src/main.go

desktop: build
	cp go-modify-aliyun-firewall  ~/go-modify-aliyun-firewall

sync: build
	scp go-modify-aliyun-firewall-linux d01:~/go-modify-aliyun-firewall-linux
releaser:
	goreleaser --snapshot --skip-publish --rm-dist;

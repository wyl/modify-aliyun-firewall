
build:
	go build -o ./go-modify-aliyun-firewall modify-aliyun-firewall.go

releaser:
	goreleaser --snapshot --skip-publish --rm-dist;

# 变量
NAME=golang-gin-template
BINDIR=bin
GOBUILD=go build

all: linux-amd64 darwin-amd64
# clean
clean:
	rm -rf bin
# linux
linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@
	chmod +x $(BINDIR)/$(NAME)-$@
# macos
darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@
	chmod +x $(BINDIR)/$(NAME)-$@


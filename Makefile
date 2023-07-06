phony:develop
export GO111MODULE = on
SERVER=expezgoserver
# alpha,release,final,auto
MAJOR?="0"
MINOR?="0"
PATCH?="1"
TAG_TYPE?="alpha"
TYPE_VERSION?="8"
DATETIME=`date +%Y%m%d%H%M`
GIT_TAG=v$(MAJOR).$(MINOR).$(PATCH)-$(TAG_TYPE).$(TYPE_VERSION)
MESSAGE?="新增加 swagger 文档"


version:Makefile
	@echo "package version" > version/ver.go
	@echo "var AppVersion=\"$(GIT_TAG)\"" >> version/ver.go
release:version
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(SERVER) main.go
init:
	@rm -f go.mod go.sum
	@go mod init tplgo
	@go mod tidy
develop:version
	@go generate main.go
	@go build -o $(SERVER) -gcflags="all=-N -l" main.go
publish:version
#linux系统 build
	git add .
	git commit -m $(MESSAGE)
	git push
	git tag -a $(GIT_TAG) -m $(MESSAGE)
	git push origin --tags
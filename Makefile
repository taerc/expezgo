phony:develop
export GO111MODULE = on
SERVER=expezgoserver
# alpha,release,final,auto
MAJOR?="0"
MINOR?="0"
PATCH?="1"
TAG_TYPE?="alpha"
TYPE_VERSION?="5"
MESSAGE?="ezgo-gitlab版本测试完成,并通过测试验证."
DATETIME=`date +%Y%m%d%H%M`
GIT_TAG=v$(MAJOR).$(MINOR).$(PATCH)-$(TAG_TYPE).$(TYPE_VERSION)


release:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(SERVER) main.go
init:
	@rm -f go.mod go.sum
	@go mod init tplgo
	@go mod tidy
develop:
	go build -o $(SERVER) -gcflags="all=-N -l" main.go
publish:
#linux系统 build
	git add .
	git commit -m $(MESSAGE)
	git push
	git tag -a $(GIT_TAG) -m $(MESSAGE)
	git push origin --tags
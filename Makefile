phony:develop
export GO111MODULE = on
SERVER=expezgoserver
# alpha,release,final,auto
MAJOR?="0"
MINOR?="0"
PATCH?="1"
TAG_TYPE?="alpha"
TYPE_VERSION?="16"
DATETIME=`date +%Y%m%d%H%M`
GIT_TAG=v$(MAJOR).$(MINOR).$(PATCH)-$(TAG_TYPE).$(TYPE_VERSION)
MESSAGE?="gorm Preload 测试"
BUILD?=build


version:Makefile
	@echo "package version" > version/ver.go
	@echo "var AppVersion=\"$(GIT_TAG)\"" >> version/ver.go
release:version
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(SERVER) cmd/expezgoserver.go
init:
	@rm -f go.mod go.sum
	@go mod init expezgo
	@go mod tidy
develop:version
	#@go generate main.go
	@go build -o build/$(SERVER) -gcflags="all=-N -l" cmd/expezgoserver.go
publish:version
#linux系统 build
	git add .
	git commit -m $(MESSAGE)
	git push
	git tag -a $(GIT_TAG) -m $(MESSAGE)
	git push origin --tags


expent:research/exp_ent_sql.go
	@go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery,sql/upsert,sql/modifier ./modules/ent/schema
	@go build -o build/expent -gcflags="all=-N -l" research/exp_ent_sql.go

geo:research/exp_ent_geo_data.go
	go build -o build/geo research/exp_ent_geo_data.go

context:research/exp_14_context.go
	go build -o build/context research/exp_14_context.go

validate:research/exp_validate.go
	go build -o $(BUILD)/validate research/exp_validate.go

kafka:research/exp_15_kafka.go
	go build -o $(BUILD)/kafka research/exp_15_kafka.go

cli-kafka:research/exp_16_cli_kafka.go
	go build -o $(BUILD)/cli-kafka research/exp_16_cli_kafka.go
embed-queue:research/exp_17_embed_queue.go
	go build -o $(BUILD)/embed-queue research/exp_17_embed_queue.go
conv-pinyin:research/exp_18_conv_pinyin.go
	go build -o $(BUILD)/conv-pinyin research/exp_18_conv_pinyin.go
server-ws:research/exp_07_server_ws.go
	go build -o $(BUILD)/server-ws research/exp_07_server_ws.go
client-ws:research/exp_07_client_ws.go
	go build -o $(BUILD)/client-ws research/exp_07_client_ws.go




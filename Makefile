phony:run
export GO111MODULE = on
SERVER=expezgoserver
# alpha,release,final,auto
MAJOR?="0"
MINOR?="0"
PATCH?="1"
TAG_TYPE?="alpha"
TYPE_VERSION?="17"
DATETIME=`date +%Y%m%d%H%M`
GIT_TAG=v$(MAJOR).$(MINOR).$(PATCH)-$(TAG_TYPE).$(TYPE_VERSION)
MESSAGE?="gorm Preload 测试"
BUILD?=build


version:Makefile
	@echo "package version" > version/ver.go
	@echo "var AppVersion=\"$(GIT_TAG)\"" >> version/ver.go

release:version
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(SERVER) cmd/expezgoserver.go

sqlmonitor:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/sqlmonitor cmd/sqlmonitor/main.go


init:
	@go mod download
	@go generate ./modules/ent
	@go mod tidy

develop:version
	@go build -o build/$(SERVER) -gcflags="all=-N -l" cmd/expezgoserver.go
publish:version
#linux系统 build
	git add .
	git commit -m $(MESSAGE)
	git push
	git tag -a $(GIT_TAG) -m $(MESSAGE)
	git push origin --tags

ent:
	@go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery,sql/upsert,sql/modifier,intercept ./modules/ent/schema

expent:research/exp_ent_sql.go
	@go build -o build/expent -gcflags="all=-N -l" research/exp_ent_sql.go

geo:research/exp_ent_geo_data.go
	go build -o build/geo research/exp_ent_geo_data.go

context:research/exp_14_context.go
	go build -o build/context research/exp_14_context.go

validate:research/exp_validate.go
	go build -o $(BUILD)/validate research/exp_validate.go

kafka:research/exp_15_kafka.go
	go build -o $(BUILD)/kafka research/exp_15_kafka.go

go-consts:research/exp_04_consts.go
	go build -o $(BUILD)/go-consts research/exp_04_consts.go

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
db-driver:research/exp_21_db_driver.go
	go build -o $(BUILD)/db-driver research/exp_21_db_driver.go

ent-hook:research/exp_19_ent_hook.go
	go build -o $(BUILD)/ent-hook research/exp_19_ent_hook.go
json:research/exp_03_json.go
	go build -o $(BUILD)/json research/exp_03_json.go
ent-sql:research/exp_27_ent_sql.go
	go build -o $(BUILD)/ent-sql research/exp_27_ent_sql.go

xls:research/exp_20_xls.go
	go build -o $(BUILD)/xls research/exp_20_xls.go
fsevent:research/exp_22_fsevent.go
	go build -o $(BUILD)/fsevent research/exp_22_fsevent.go
videoindex:research/exp_23_videoindex.go
	go build -o $(BUILD)/videoindex research/exp_23_videoindex.go
tcpserver:research/exp_24_tcpserver.go
	go build -o $(BUILD)/tcpserver research/exp_24_tcpserver.go
byteorder:research/exp_25_byteorder.go
	go build -o $(BUILD)/byteorder research/exp_25_byteorder.go
pdf:research/exp_29_pdf.go
	go build -o $(BUILD)/pdf research/exp_29_pdf.go

word:research/exp_28_word.go
	go build -o $(BUILD)/word research/exp_28_word.go



graph:
	@entviz ./modules/drone
	@mv schema-viz.html docs



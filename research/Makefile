phony:

export GO111MODULE = on
image: exp_11_template.go
	go build -o image exp_11_template.go
gen:
	go generate
	go build -o gen exp_13_generate.go

strings:
	go generate "research/errorcode"
	go build -o gen exp_13_generate.go

context:exp_14_context.go
	go build -o context exp_14_context.go

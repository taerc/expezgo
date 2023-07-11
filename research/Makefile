phony:expent
export GO111MODULE="off"

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

expent:exp_ent_sql.go
	go build -o expent exp_ent_sql.go

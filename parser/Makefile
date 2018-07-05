all : parser.go 

parser.go : parser.go.y ../ast/*.go
	goyacc -o $@ parser.go.y
	gofmt -s -w .

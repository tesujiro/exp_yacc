all : interpretor

interpretor : interpretor.go
	go build -o interpretor .

test: *_test.go
	go test -v .

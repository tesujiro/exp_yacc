all : interpretor

interpretor : interpretor.go
	go build .

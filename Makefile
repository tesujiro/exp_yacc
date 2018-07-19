all : interpretor

interpretor : interpretor.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go
	go build -o interpretor .

test: *_test.go
	go test -v . ./vm ./lib

prof:
	# make prof ARG=[PPROF PATH]
	go tool pprof --pdf ./interpretor ${ARG} > ./prof.pdf

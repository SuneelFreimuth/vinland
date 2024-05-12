test: test/parse_test.go
	go test test/*.go

gen: src/parser/Vinland.g4
	go generate ./...
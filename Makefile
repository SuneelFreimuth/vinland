test: .FORCE
	go test test/*.go

gen: src/parser/Vinland.g4
	go generate ./...

format: .FORCE
	go fmt ./src ./src/parser ./src/ast

.FORCE:
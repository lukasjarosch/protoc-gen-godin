
.PHONY: example
example:
	@go install && cd ./example && protoc -I . --go_out=. --godin_out=. example.proto
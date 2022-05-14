generate-proto:
	buf generate

install:
	go install ./cmd/protoc-gen-govalidate

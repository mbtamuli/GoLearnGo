build:
	go build -o protobuf main.go

clean:
	rm -rf out
	rm -f protobuf

regenerate-protobuf-classes: output-dir
	protoc \
		--proto_path=proto \
		--go_out=pb \
		--go_opt=paths=source_relative \
		person.proto

output-dir:
	mkdir -p pb

run: build generate-protobuf-classes
	./protobuf
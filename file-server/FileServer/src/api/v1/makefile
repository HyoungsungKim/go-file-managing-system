# cat -e -t -v makefile 
# https://stackoverflow.com/a/38440100
# If you are using vscode, chage "Space:4" to tab 
compile:
	protoc *.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
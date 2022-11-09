gen:
	rm pb/* || true
	protoc -I=proto/read_only --go_out=pb/ proto/read_only/*.proto
	protoc --go-grpc_out=pb/ proto/read_only/*.proto -I=proto/read_only
	protoc -I=proto/writable --go_out=pb/ proto/writable/*.proto
	protoc --go-grpc_out=pb/ proto/writable/*.proto -I=proto/writable
git:
	$(MAKE) gen
	git add .
	git commit -a -m "$m"
	git push -u origin main
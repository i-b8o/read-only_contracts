gen:
	rm -rf pb/* || true
	protoc -I=proto/read_only --go_out=pb/ proto/read_only/*.proto
	protoc --go-grpc_out=pb/ proto/read_only/*.proto -I=proto/read_only
	protoc -I=proto/writable --go_out=pb/ proto/writable/*.proto
	protoc --go-grpc_out=pb/ proto/writable/*.proto -I=proto/writable
	protoc -I=proto/supreme --go_out=pb/ proto/supreme/*.proto
	protoc --go-grpc_out=pb/ proto/supreme/*.proto -I=proto/supreme
	
update:
	$(MAKE) gen
	git add .
	git commit -a -m '$m'
	git push -u origin main
	./update.sh
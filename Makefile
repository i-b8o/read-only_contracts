gen:
	protoc -I=proto/read_only --go_out=pb/read_only proto/read_only/*.proto
	protoc --go-grpc_out=pb/read_only proto/read_only/*.proto -I=proto/read_only
	protoc -I=proto/writable --go_out=pb/writable proto/writable/*.proto
	protoc --go-grpc_out=pb/writable proto/writable/*.proto -I=proto/writable
	protoc -I=proto/supreme --go_out=pb/supreme proto/supreme/*.proto
	protoc --go-grpc_out=pb/supreme proto/supreme/*.proto -I=proto/supreme
	
git:
	$(MAKE) gen
	git add .
	git commit -a -m "$m"
	git push -u origin main
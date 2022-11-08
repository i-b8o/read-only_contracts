gen:
	rm pb/* || true
	protoc -I=proto/ --go_out=pb/ proto/*.proto
	protoc --go-grpc_out=pb/ proto/*.proto -I=proto/

git:
	git add .
	git commit -a -m "$m"
	git push -u origin main
gen:
	rm -rf pb/* || true
	protoc -I=proto/read_only --go_out=pb/ proto/read_only/*.proto
	protoc --go-grpc_out=pb/ proto/read_only/*.proto -I=proto/read_only
	protoc -I=proto/writable --go_out=pb/ proto/writable/*.proto
	protoc --go-grpc_out=pb/ proto/writable/*.proto -I=proto/writable
	protoc -I=proto/supreme --go_out=pb/ proto/supreme/*.proto
	protoc --go-grpc_out=pb/ proto/supreme/*.proto -I=proto/supreme
	
git:
	$(MAKE) gen
	git add .
	git commit -a -m "$m"
	git push -u origin main

pull:
	git pull origin main

push_all:
	$(MAKE) git
	../regulations_writable_service/git_push.sh "$m" || true
	../regulations_supreme_service/git_push.sh "$m" || true
	../regulations_read_only_service/git_push.sh "$m" || true
	../regulations_parser/git_push.sh "$m" || true
	../regulations_mobile_service/git_push.sh "$m" || true

pull_all:
	$(MAKE) pull
	../regulations_writable_service/git_pull.sh "$m" || true
	../regulations_supreme_service/git_pull.sh "$m" || true
	../regulations_read_only_service/git_pull.sh "$m" || true
	../regulations_parser/git_pull.sh "$m" || true
	../regulations_mobile_service/git_pull.sh "$m" || true

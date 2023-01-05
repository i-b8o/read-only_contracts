gen:
	rm -rf pb/* || true
	protoc -I=proto/reader --go_out=pb/ proto/reader/*.proto
	protoc --go-grpc_out=pb/ proto/reader/*.proto -I=proto/reader
	protoc -I=proto/writer --go_out=pb/ proto/writer/*.proto
	protoc --go-grpc_out=pb/ proto/writer/*.proto -I=proto/writer
	protoc -I=proto/master --go_out=pb/ proto/master/*.proto
	protoc --go-grpc_out=pb/ proto/master/*.proto -I=proto/master
	protoc -I=proto/searcher --go_out=pb/ proto/searcher/*.proto
	protoc --go-grpc_out=pb/ proto/searcher/*.proto -I=proto/searcher
	protoc --dart_out=grpc:../read_only_dashboard/lib/pb/ -Iproto proto/master/*.proto

git:
	git add .
	git commit -a -m '$m' || true
	git push -u origin main || true
	

update_contracts:
	go get -u github.com/i-b8o/read-only_contracts@$m

test_db:
	sudo docker compose up --build test

# TODO Search service
# TODO for registered (or not only) users feauter to save chain of paragraphs/chapters
GCP_PROJECT=xxxx
CHILD_PORT=8888

build-single:
	go build -o exe ./singleserver

single-dockerbuild:
	docker build . -f Dockerfile -t opencensus-grpc-go-sample:vtest

build-child:
	go build -o exe ./childserver

run-parent:
	go run ./parentserver -project $(GCP_PROJECT)

run-child:
	go run ./childserver -project $(GCP_PROJECT) -p $(CHILD_PORT)

load:
	./load.sh
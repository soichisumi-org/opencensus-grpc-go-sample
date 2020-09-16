GCP_PROJECT=xxxx
CHILD_PORT=8888

run-parent:
	go run ./parentserver -project $(GCP_PROJECT)
run-child:
	go run ./childserver -project $(GCP_PROJECT) -p $(CHILD_PORT)
load:
	./load.sh
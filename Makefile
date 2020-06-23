run-parent:
	go run ./parentserver -project "gcp-project"
run-child:
	go run ./childserver -project "gcp-project" -p 8888
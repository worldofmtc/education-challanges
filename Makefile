run-challanges:
	go test -v ./...
run-challange $(challange):
	go test -v ./... -run $(challange)
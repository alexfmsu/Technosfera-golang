run:
	go run main.go
	
test:
	go test -v

cover:
	go test -coverprofile=cover.out ./

fmt:
	gofmt -w *.go

html:
	go tool cover -html=cover.out -o cover.html

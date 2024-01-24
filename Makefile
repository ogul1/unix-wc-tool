build:
	go build -o go-wc main.go service.go

clean:
	go clean

test:
	go test
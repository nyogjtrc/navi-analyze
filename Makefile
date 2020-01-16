build: *
	GOOS=linux GOARCH=amd64 go build -o build/navi-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o build/navi-mac main.go

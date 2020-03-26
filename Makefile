default:
	go get pack.ag/tftp
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tftp

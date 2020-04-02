release :
	go build -o yaml2json main.go

clean :
	rm -f yaml2json

proxy :
	go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct

default:

build:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go build -o a.out

test:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go test

tidy:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go mod tidy

version:
	goversion -m a.out

clean:

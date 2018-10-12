default:

mod:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go test

tidy:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go mod tidy

clean:

default:

build:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go build -o a.out

update:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go get -u

update-patch:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go get -u=patch

test:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go test

tidy:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go mod tidy

graph:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go mod graph

graph-dot:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go mod graph | awk -f graph.awk > graph.dot
	dot -Tpng -o graph.dot.png graph.dot

version:
	goversion -m a.out

clean:
	HTTPS_PROXY=socks5://127.0.0.1:2080 go clean

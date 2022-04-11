1. Module name == Github repo URL
2. go mod tidy
3. go test
4. git commit
5. git tag v0.1.0
6. git push origin v0.1.0
7. GOPROXY=proxy.golang.org go list -m github.com/ryan730824/greetings@v0.1.0
8. go get github.com/ryan730824/greetings@v0.1.0
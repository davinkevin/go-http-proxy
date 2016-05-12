**Building for Raspberry PI :**

 ```
 docker run --rm -v /Users/kdavin/Workspace/Go/src/github.com/davinkevin/go-http-proxy:/usr/src/go-http-proxy-arm -w /usr/src/go-http-proxy-arm -e GOOS=linux -e GOARCH=arm golang:1.5.4-alpine go build -v
 ```
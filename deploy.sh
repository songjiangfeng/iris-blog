#/bin/sh
echo start ..
cp -r html assets logs uploads config.json.default deploy/
go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go build -o deploy/iris_blog_linux

go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64
go build -o deploy/iris_blog_windows.exe

go env -w CGO_ENABLED=0 GOOS=darwin GOARCH=amd64

echo finished ..





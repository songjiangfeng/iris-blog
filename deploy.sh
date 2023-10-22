#/bin/sh

#works only for  mac 

echo start ..

deploy_path=../deploy
mkdir -p $deploy_path
cp -r app  iris_blog.sql  $deploy_path


go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go build -o $deploy_path/iris_blog_linux

go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64
go build -o $deploy_path/iris_blog_windows.exe

go env -w CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
go build -o $deploy_path/iris_blog_darwin

zip -q -r iris_blog`date +%Y%m%d%H%M`.zip $deploy_path

echo finished ..





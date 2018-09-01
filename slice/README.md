# slice

##### 运行需要安装
- glide

```shell
glide install
go test -cover=true -coverprofile=/tmp/test-cover.out
go tool cover -html=/tmp/test-cover.out -o=/tmp/test-cover.html
```

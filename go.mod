module github.com/linuxxiaoyu/go-gin-example

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200824131525-c12d262b63d8 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.60.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/linuxxiaoyu/go-gin-example/pkg/setting	=> ./pkg/setting
)
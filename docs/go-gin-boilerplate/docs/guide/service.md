---
sidebar_position: 3
---

# 服务

服务层是业务逻辑的具体实现，主要负责处理业务逻辑，包括数据的读取、写入、修改、删除等操作。

在 `internal/service` 目录下，用于存放服务层相关代码。

## 服务示例

在 `internal/service` 目录下，我们规定一个文件声明一个服务，例如 `example.go` 文件，就是我们的示例服务，用于存放示例服务的定义。

```go

type ExampleService struct{}

func (exampleService *ExampleService) CreateExample(data map[string]interface{}) *models.Example {

	example := models.Example{
		Name: data["name"].(string),
	}

	res := db.Create(&example)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	}

	return &example

}
```

上面的代码中，我们定义了 `ExampleService` 结构体，用于存放示例服务的方法。

在 `CreateExample` 方法中，我们接收一个 `map[string]interface{}` 类型的参数，用于接收请求参数，然后创建 `Example` 模型，并将其写入数据库。

确保 在 服务层，我们只处理业务逻辑，不要处理请求参数的解析、响应的返回等操作。

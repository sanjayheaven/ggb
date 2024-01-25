---
sidebar_position: 3.1
---

# 模型

模型层负责数据结构的定义，Go Gin Boilerplate 使用 GORM 作为 ORM 工具，支持 MySQL、PostgreSQL、SQLite、SQL Server 等多种数据库。

所以在 模型 定义的时候，我们可以使用 gin 框架提供的 gorm 标签，来定义模型的字段类型、长度、索引、唯一性等。

在 `internal/model` 目录下，用于存放模型层相关代码。

## 基础模型

`internal/model.go` 作为模型包的入口文件，用于定义模型的基础模型等内容。

在 `internal/model.go` 中，我们定义 `BasicModel` 结构体，作为所有模型的基础模型。

```go
// BasicModel 基础模型

type BasicModel struct {
    ID        uint       `gorm:"primarykey" json:"id"` // ID
    CreatedAt time.Time  `json:"created_at"`           // 创建时间
    UpdatedAt time.Time  `json:"updated_at"`           // 更新时间
    DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
```

### 嵌入基础模型

在定义模型的时候，我们可以嵌入 `BasicModel` 结构体，来继承基础模型的字段。

```go

type Example struct {
    BasicModel

    // other code field ...
}
```

## 模型定义

在 `internal/model` 目录下，我们规定一个文件声明一个模型，例如 `example.go` 文件，就是我们的示例模型，用于存放示例模型的定义。

```go
package models

type Example struct {
    BasicModel

    Name   string `json:"name"`                         // Name
    Status string `json:"status" gorm:"default:active"` // Status, active or inactive
}
```

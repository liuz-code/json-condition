# 背景
因为项目上要求，在做数据推送的时候需要根据json的内容将不同的数据推送到不同的地址中，所以我开发了这个rule条件的配置插件，根据rule的配置可以搜索出满足条件的json数据，然后根据项目中配置的推送地址将数据推送出去。

# Json条件检索

go语言，通过rule条件筛选出json是否可用。可以在批量的json中快速筛选出满足条件的json。

数据筛选支持多层数据搜索，可以通过.进行多层惊悚node分割，也可以对数组进行搜索。代码可能存在很多不足，请大家多多指点共同进步，希望大家能一起参与进来。

example中有试用方式请参考试用。

## Rule协议

| 条件运算符:          | 说明:          |
| ----------------- | ----------------- |
| and | 等于，操作数据等于对象中的数据 |
| or | 或，操作数据或，等于其中一个即可 |
| not | 不等于，操作数据不等于对象中的数据 |
| gt | 大于，操作数据大于对象中的数据，必须为int或float类型的数据 |
| gte | 大于等于，操作数据大于或等于对象中的数据，必须为int或float类型的数据 |
| lt | 小于，操作数据小于对象中的数据，必须为int或float类型的数据 |
| lte | 小于等于，操作数据小于或等于对象中的数据，必须为int或float类型的数据 |
| like | 模糊，操作数据会模糊查询对象中定义的数据前、后、中间出现过大于一次的命中，必须为string类型 |

#### 协议示例

```json
{
  "and": {
    "projectId": "10001",
    "code": "test-push"
  },
  "or": {
    "projectId": "10001",
    "code": "test-push"
  },
  "not": {
    "projectId": "10001",
    "code": "test-push"
  },
  "gt": {
    "age": 10,
    "number": 5
  },
  "gte": {
    "age": 10,
    "number": 5
  },
  "lt": {
    "age": 10,
    "number": 5
  },
  "lte": {
    "age": 10,
    "number": 5
  },
  "like": {
    "name": "你好"
  }
}
```

## Data数据

data为rule筛选的数据源。
```json
{
  "projectId": "10001",
  "code": "test-push",
  "name": "xiaoming",
  "days": [
    "one",
    "two",
    "three"
  ]
}
```

## 运行测试

执行 `go run .\example\condition\example.go 或 .\example\find\main.go`
# 08 API契约与Schema注册中心

纯Go契约注册中心，提供OpenAPI、AsyncAPI和JSON Schema资产的版本化登记、引用解析、兼容性判断、样例校验、消费者影响分析和游标分发接口。

## 启动

```bash
go run ./cmd/registry -config configs/config.yaml
```

默认监听`:8088`。健康检查：`GET /healthz`、`GET /readyz`、`GET /metrics`。

## 主要API

- `POST /v1/namespaces` 创建命名空间
- `POST /v1/assets` 登记资产（kind=openapi|asyncapi|jsonschema）
- `POST /v1/assets/{id}/versions` 创建版本，状态默认为draft
- `POST /v1/assets/{id}/versions/{version}/publish` 发布版本并执行兼容性门禁
- `POST /v1/compatibility` 比较两个JSON契约
- `POST /v1/validate` 校验JSON样例
- `POST /v1/consumers` 注册消费者
- `GET /v1/impact?asset_id=...&version=...` 查询影响消费者
- `GET /v1/changes?cursor=...` 按游标读取变更
- `POST /v1/import` 原子导入契约包

所有写入操作支持`Idempotency-Key`。服务使用内存仓储便于开发，接口可替换为PostgreSQL/对象存储；迁移文件位于`migrations/`。

## 验证

```bash
go test ./...
go vet ./...
curl -s localhost:8088/healthz
```

示例契约位于`examples/`，包括引用循环拒绝、兼容性差异、样例错误路径和消费者影响查询流程。

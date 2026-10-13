# OwlMail 0.11.0

发布日期：2026-10-13

本功能版本加入确定性的邮件测试流程、邮件断言、投递记录与重试，以及感知配置的诊断和健康检查。以下变更以原始 `v0.10.1` 发布版本为基线。

## 新增

- **投递游标与测试运行。** 触发发信前取得 cursor，再通过原生 REST 等待接口或 MCP `wait_for_email` 查询游标之后已保留的投递。精确 SMTP envelope 收件人与 `X-OwlMail-Run-ID` 筛选器组织测试邮件；按测试运行清理的 REST 接口避免删除无关数据。
- **邮件断言。** 只读 `/check` 接口验证主题、纯文本、附件数量与元数据、允许的链接主机，并通过有界正则表达式提取内容。新增 **Verify** 页签提供正文与链接检查，不访问 URL，也不修改已读状态。
- **投递记录与重试。** Relay 任务列表和持久化 Webhook 历史展示尝试次数、失败类别与持久化健康状态。**Deliveries** 页签可查看记录，并确认手动重试。Webhook 仅重放失败目标并复用投递 ID；原邮件与适用的目标配置必须仍然存在。
- **本地诊断和 SMTP-only 模式。** `owlmail doctor` 离线检查配置与 Relay 恢复记录。`-disable-web` 仅运行 SMTP，不提供 Web、REST、WebSocket 或 HTTP MCP。`healthcheck` 命令与容器启动快照支持自定义监听配置、子路径、HTTPS 和 SMTP-only 部署。
- **命令面板与回归验证。** 邮箱快捷操作提供可搜索命令面板。真实浏览器与容器检查覆盖预览、主题、快捷键、子路径和打包后的健康检查。暗色模式与七种预览宽度在 0.10.1 已经存在，不属于本版新增。
- **分发许可文本。** 发布文件 `owlmail-licenses.tar.gz`、主分支构建快照及镜像中的 `/usr/share/licenses/owlmail/` 保存项目、实际链接模块和 Go 工具链的原始许可证与 notice，并提供哈希清单。许可归档包含在签名校验和及来源证明中。

## 更新与修正

- readiness 现在检查 SMTP 监听状态、本地邮箱写入与同步的缓存探针、已启用的附件存储、Relay worker 与任务持久化状态，以及已配置下游 SMTP/TLS/认证的探针。原 0.10.1 只检查已启用的 S3 附件存储。配置了出站主机时，即使未启用自动 Relay，也会参与就绪检查。
- 容器探针读取服务启动时保存的私有配置快照，不再固定假设 HTTP 1080。此模式要求 OwlMail 为 PID 1；使用 init wrapper 的部署需要提供明确的探针配置。
- 中英文入口文档改为安装、浏览器、CI 与 Agent 工作流，并修正 MCP / REST 等待结果、临时 EML 存储、Origin 检查和历史 Relay 接口说明。既有 SMTP 认证模式、HTML 隔离与需显式开启的只读 MCP 保留原有边界，本版不会默认强制 SMTP AUTH 或开启 MCP。

## 升级

1. 停止上游 SMTP 流量并关闭旧实例。备份**完整**邮件目录或卷，包括 EML、附件、sidecar 与排队投递记录。确认不再需要回滚前保留备份。
2. 部署 `ghcr.io/soulteary/owlmail:0.11.0` 或已验证的 manifest digest。重要归档先在副本上验证；每个目录由一个可写实例持有。
3. 分别检查 liveness 和 readiness。使用 `/healthz` 决定是否重启，使用 `/readyz` 判断依赖是否就绪。存储或下游故障可能使 readiness 返回 503，而 liveness 仍为 200；检查此前把 readiness 当作重启探针的部署。
4. 核对恢复后的邮件、已读状态、原始源码和附件，再发送一封唯一测试邮件，并检查已配置的 Webhook 或 Relay 恢复路径。

本地 EML 与 sidecar 继续作为邮箱权威数据源；这不提供共享数据库，也不承诺完整 MailDev 协议兼容。更早版本的升级还应阅读 [0.10.0 迁移说明](Release-0.10.0.md)。

## 工作流边界

- cursor 属于单个进程，重启后失效；`run_id` 是测试筛选器，不是租户鉴权。REST 等待接受 `count` 并返回 `emails`；MCP 等待返回单个可选 `email`、`matched` 与 `timed_out`，没有 `count` 参数。
- 按运行清理处理有界快照，并保护未完成 Relay 的源邮件；需要检查失败 ID 和 `more_possible`。不要全局清空共享测试邮箱。
- Webhook 和恢复后的 Relay 是**至少一次**投递，可能重复。接收器应按 Webhook 投递 ID 去重。没有 Redis 时，Webhook 任务从本地 outbox 进入内存队列后不再保证跨重启恢复。
- doctor 是离线检查，不证明服务可用或存储可写。Relay 接受任务不表示最终投递成功。邮件断言检查已存储内容并报告截断，不认证客户端渲染，不访问链接，也不验证提取出的验证码是否可用。
- OwlMail 仍面向开发、CI 和可信测试网络，SMTP 默认 NO AUTH。共享部署应配置认证、HTTPS 与网络访问范围。MCP 继续默认关闭且只读；原始 MIME 可能包含附件载荷，但没有独立的解码附件下载工具。

限制与合约见[投递等待](../en/Delivery-Waits-and-Test-Runs.md)、[API 参考](API-Reference.md)、[MCP](MCP-Reference.md)与[运维指南](Operations.md)。

## 安装

```bash
docker pull ghcr.io/soulteary/owlmail:0.11.0
docker run --rm \
  -p 127.0.0.1:1025:1025 \
  -p 127.0.0.1:1080:1080 \
  -v owlmail-data:/app/mail \
  ghcr.io/soulteary/owlmail:0.11.0
```

为了可重复部署，请记录已发布的 manifest digest 并使用
`ghcr.io/soulteary/owlmail@sha256:<digest>`。

## 发布文件

- `checksums.txt` 与 `checksums.txt.sigstore.json`
- `owlmail-linux-amd64` 与 `owlmail-linux-amd64.spdx.json`
- `owlmail-linux-arm64` 与 `owlmail-linux-arm64.spdx.json`
- `owlmail-darwin-amd64` 与 `owlmail-darwin-amd64.spdx.json`
- `owlmail-darwin-arm64` 与 `owlmail-darwin-arm64.spdx.json`
- `owlmail-windows-amd64.exe` 与 `owlmail-windows-amd64.exe.spdx.json`
- `owlmail-licenses.tar.gz`

Linux amd64 下载、校验与启动示例：

```bash
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.0/owlmail-linux-amd64
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.0/checksums.txt
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.0/checksums.txt.sigstore.json
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.0/owlmail-licenses.tar.gz
grep ' owlmail-linux-amd64$' checksums.txt | sha256sum -c -
grep ' owlmail-licenses.tar.gz$' checksums.txt | sha256sum -c -
gh attestation verify owlmail-linux-amd64 --repo soulteary/owlmail
cosign verify-blob \
  --bundle checksums.txt.sigstore.json \
  --certificate-identity-regexp '^https://github.com/soulteary/owlmail/.github/workflows/release.yml@refs/tags/v' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  checksums.txt
tar -xzf owlmail-licenses.tar.gz
chmod +x owlmail-linux-amd64
./owlmail-linux-amd64
```

重新分发二进制时保留解压得到的 `owlmail-licenses/` 目录。归档覆盖 OwlMail 与其 Go 依赖；镜像操作系统软件包继续适用各自条款。

```bash
cosign verify \
  --certificate-identity-regexp '^https://github.com/soulteary/owlmail/.github/workflows/release.yml@refs/tags/v' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  ghcr.io/soulteary/owlmail:0.11.0
```

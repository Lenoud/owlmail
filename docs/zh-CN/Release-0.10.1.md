# OwlMail 0.10.1

发布日期：2026-10-11

此维护版本修复普通 SMTP listener 的关闭竞态，并更新编译器与依赖。

## 修复

在 `smtp.Server.Serve` 接管 listener 前，`Close` 直接关闭普通 SMTP listener。快速启动后立即停止不再留下占用端口的 listener 和阻塞的服务协程。

## 更新

- 使用 Go 1.27.2 构建。
- 升级 soulteary kit 和第三方依赖，包括 SQLite/modernc 与 AWS SDK。
- 从同一不可变 tag 发布二进制和多架构容器镜像。

## 升级

使用 `ghcr.io/soulteary/owlmail:0.10.1` 或已验证的镜像 digest。运行时 API、配置和存储格式兼容 0.10.0，并保留其安全保护。更早版本的部署还应阅读 [0.10.0 迁移说明](Release-0.10.0.md)。

## 安装

```bash
docker pull ghcr.io/soulteary/owlmail:0.10.1
docker run --rm \
  -p 127.0.0.1:1025:1025 \
  -p 127.0.0.1:1080:1080 \
  -v owlmail-data:/app/mail \
  ghcr.io/soulteary/owlmail:0.10.1
```

为了可重复部署，请记录已发布的 manifest digest 并使用
`ghcr.io/soulteary/owlmail@sha256:<digest>`。

## 发布文件

- `checksums.txt`
- `checksums.txt.sigstore.json`
- `owlmail-linux-amd64` 与 `owlmail-linux-amd64.spdx.json`
- `owlmail-linux-arm64` 与 `owlmail-linux-arm64.spdx.json`
- `owlmail-darwin-amd64` 与 `owlmail-darwin-amd64.spdx.json`
- `owlmail-darwin-arm64` 与 `owlmail-darwin-arm64.spdx.json`
- `owlmail-windows-amd64.exe` 与
  `owlmail-windows-amd64.exe.spdx.json`

Linux amd64 下载、校验与启动示例：

```bash
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.10.1/owlmail-linux-amd64
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.10.1/checksums.txt
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.10.1/checksums.txt.sigstore.json
grep ' owlmail-linux-amd64$' checksums.txt | sha256sum -c -
gh attestation verify owlmail-linux-amd64 --repo soulteary/owlmail
cosign verify-blob \
  --bundle checksums.txt.sigstore.json \
  --certificate-identity-regexp '^https://github.com/soulteary/owlmail/.github/workflows/release.yml@refs/tags/v' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  checksums.txt
chmod +x owlmail-linux-amd64
./owlmail-linux-amd64
```

```bash
cosign verify \
  --certificate-identity-regexp '^https://github.com/soulteary/owlmail/.github/workflows/release.yml@refs/tags/v' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  ghcr.io/soulteary/owlmail:0.10.1
```


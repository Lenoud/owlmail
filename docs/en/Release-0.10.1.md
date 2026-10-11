# OwlMail 0.10.1

Released: 2026-10-11

This maintenance release fixes a plain SMTP listener shutdown race and refreshes the build toolchain and dependencies.

## Fixed

`Close` now closes the plain SMTP listener directly during the startup window before `smtp.Server.Serve` owns it. Rapid startup/shutdown no longer leaves the listener bound and the serving goroutine blocked.

## Changed

- Build with Go 1.27.2.
- Upgrade soulteary kits and third-party dependencies, including SQLite/modernc and AWS SDK packages.
- Publish binaries and a multi-architecture container image from the same immutable tag.

## Upgrade

Use `ghcr.io/soulteary/owlmail:0.10.1` or a verified image digest. Runtime APIs, configuration, and storage formats remain compatible with 0.10.0. The security protections introduced in 0.10.0 remain enabled; deployments upgrading from older releases should also read [the 0.10.0 migration notes](Release-0.10.0.md).

## Install

```bash
docker pull ghcr.io/soulteary/owlmail:0.10.1
docker run --rm \
  -p 127.0.0.1:1025:1025 \
  -p 127.0.0.1:1080:1080 \
  -v owlmail-data:/app/mail \
  ghcr.io/soulteary/owlmail:0.10.1
```

For repeatable deployment, record the published manifest digest and use
`ghcr.io/soulteary/owlmail@sha256:<digest>`.

## Release artifacts

- `checksums.txt`
- `checksums.txt.sigstore.json`
- `owlmail-linux-amd64` and `owlmail-linux-amd64.spdx.json`
- `owlmail-linux-arm64` and `owlmail-linux-arm64.spdx.json`
- `owlmail-darwin-amd64` and `owlmail-darwin-amd64.spdx.json`
- `owlmail-darwin-arm64` and `owlmail-darwin-arm64.spdx.json`
- `owlmail-windows-amd64.exe` and
  `owlmail-windows-amd64.exe.spdx.json`

Linux amd64 download, verification, and launch example:

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


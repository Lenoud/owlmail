# OwlMail 0.11.1

Released: 2026-10-13

This release completes publication of the 0.11 feature update. The `v0.11.0`
source tag and container image remain available, but that workflow was
interrupted before image signing and binary release creation. Use `v0.11.1`
for the complete release. Application behavior is unchanged from the
`v0.11.0` source.

The feature update adds deterministic email-test workflows, message checks,
delivery history and retries, and configuration-aware diagnostics and probes.
The changes below are relative to the original `v0.10.1` release.

## Added

- **Delivery cursors and test runs.** Capture a cursor before triggering mail,
  then use the native REST wait endpoint or MCP `wait_for_email` to include
  retained deliveries after it. Exact SMTP envelope-recipient and
  `X-OwlMail-Run-ID` selectors group test mail; a run-scoped REST cleanup
  endpoint avoids deleting unrelated messages.
- **Message assertions.** The read-only `/check` endpoint verifies subject and
  plain-text expectations, attachment count and metadata, and allowed link
  hosts. It also extracts values with bounded regular expressions. The new
  **Verify** tab exposes content and link checks without fetching URLs or
  changing read state.
- **Delivery inspection and replay.** Relay job lists and persisted Webhook
  history show attempts, categorized failures, and persistence health. The
  **Deliveries** tab can inspect them and confirm a manual retry. Webhook
  retries replay only failed targets and reuse the delivery ID; the source
  message and applicable target configuration must still exist.
- **Local diagnostics and SMTP-only mode.** `owlmail doctor` inspects
  configuration and Relay recovery records offline. `-disable-web` runs SMTP
  without Web, REST, WebSocket, or HTTP MCP. The `healthcheck` command and
  container startup snapshot support listener overrides, base paths, HTTPS,
  and SMTP-only deployments.
- **Mailbox command palette and regression coverage.** Keyboard commands now
  have a searchable palette. Live browser and container checks cover preview,
  theme, shortcut, base-path, and packaged healthcheck behavior. Dark mode and
  seven preview widths already existed in 0.10.1 and are not new in this release.
- **Distributed license texts.** `owlmail-licenses.tar.gz`, main-build snapshots,
  and `/usr/share/licenses/owlmail/` in the image preserve original project,
  linked-module, and Go toolchain license and notice files, with a hashed
  manifest. The archive is included in signed release checksums and provenance.

## Changed and corrected

- Readiness now checks bound SMTP listeners, cached local mailbox write/sync
  health, enabled attachment storage, Relay worker and persistence state, and
  configured downstream SMTP/TLS/authentication probes. Original 0.10.1 checked
  enabled S3 attachment storage only. A configured outgoing host participates
  even when automatic Relay is disabled.
- The container probe reads the server's private startup configuration snapshot
  instead of assuming HTTP port 1080. It requires OwlMail as PID 1; deployments
  using an init wrapper must supply an explicit probe configuration.
- English and Chinese entry documentation now follows setup, browser, CI, and
  Agent workflows. Documentation distinguishes MCP and REST wait shapes,
  temporary EML storage, Origin checks, and historical Relay behavior. Existing
  SMTP authentication modes, HTML isolation, and opt-in read-only MCP retain
  their existing boundaries; this release does not enable SMTP AUTH or MCP by default.

## Upgrade

1. Stop new SMTP traffic and shut down the old instance. Back up the **complete**
   mail directory or volume, including EML, attachments, sidecars, and queued
   delivery records. Keep the backup until rollback is no longer needed.
2. Deploy `ghcr.io/soulteary/owlmail:0.11.1` or a verified manifest digest. Test
   an important archive on a copy first; one writable instance owns a directory.
3. Check liveness separately from readiness. Use `/healthz` for restart
   decisions and `/readyz` for dependency readiness. A downstream or storage
   outage can now make readiness return 503 while liveness remains 200. Review
   deployments that previously used readiness as their restart probe.
4. Verify restored messages, read state, raw source, and attachments, then send
   one unique test email. Check any configured Webhook or Relay recovery path.

Local EML and sidecar storage remain the mailbox's source of truth; this is not
a shared database or an assurance of full MailDev protocol compatibility.
For upgrades from older releases, also read [0.10.0's migration notes](Release-0.10.0.md).

## Workflow boundaries

- A delivery cursor belongs to one process and expires on restart. `run_id` is
  a test selector, not tenant authentication. REST waits accept `count` and
  return `emails`; MCP waits return one optional `email` with `matched` and
  `timed_out`, and have no `count` argument.
- Run cleanup handles a bounded snapshot and protects pending Relay sources;
  inspect failed IDs and `more_possible`. Do not clear a shared inbox globally.
- Webhook and recovered Relay delivery are **at least once** and may repeat.
  Deduplicate Webhook delivery IDs. Without Redis, a Webhook job is not
  restart-safe after leaving the local outbox for the in-memory queue.
- Doctor is offline and does not prove runtime availability or storage write
  access. An accepted Relay job is not a successful final delivery. Verification
  checks stored content and reports truncation; it does not certify rendering,
  fetch links, or validate that an extracted code is usable.
- OwlMail remains for development, CI, and trusted test networks. SMTP defaults
  to NO AUTH. Configure authentication, HTTPS, and network access for shared
  deployments. MCP remains opt-in and read-only; raw MIME may contain attachment
  payloads, but there is no dedicated decoded-attachment download tool.

See [delivery waits](Delivery-Waits-and-Test-Runs.md), [the API reference](API-Reference.md),
[MCP](MCP-Reference.md), and [operations](Operations.md) for limits and contracts.

## Install

```bash
docker pull ghcr.io/soulteary/owlmail:0.11.1
docker run --rm \
  -p 127.0.0.1:1025:1025 \
  -p 127.0.0.1:1080:1080 \
  -v owlmail-data:/app/mail \
  ghcr.io/soulteary/owlmail:0.11.1
```

For repeatable deployment, record the published manifest digest and use
`ghcr.io/soulteary/owlmail@sha256:<digest>`.

## Release artifacts

- `checksums.txt` and `checksums.txt.sigstore.json`
- `owlmail-linux-amd64` and `owlmail-linux-amd64.spdx.json`
- `owlmail-linux-arm64` and `owlmail-linux-arm64.spdx.json`
- `owlmail-darwin-amd64` and `owlmail-darwin-amd64.spdx.json`
- `owlmail-darwin-arm64` and `owlmail-darwin-arm64.spdx.json`
- `owlmail-windows-amd64.exe` and `owlmail-windows-amd64.exe.spdx.json`
- `owlmail-licenses.tar.gz`

Linux amd64 download, verification, and launch example:

```bash
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.1/owlmail-linux-amd64
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.1/checksums.txt
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.1/checksums.txt.sigstore.json
curl -fLO https://github.com/soulteary/owlmail/releases/download/v0.11.1/owlmail-licenses.tar.gz
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

Keep the extracted `owlmail-licenses/` directory with redistributed binaries.
The bundle covers OwlMail and its Go dependencies; image operating-system
packages retain their own terms.

```bash
cosign verify \
  --certificate-identity-regexp '^https://github.com/soulteary/owlmail/.github/workflows/release.yml@refs/tags/v' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  ghcr.io/soulteary/owlmail:0.11.1
```

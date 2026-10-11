# Go Report Card

**Grade: A+** (95.7%)

| Metric | Value |
| ------ | ----- |
| Files | 134 |
| Issues | 70 |

## Checks

| Check | Score |
| ----- | ----- |
| gofmt | 100% |
| go_vet | 100% |
| ineffassign | 100% |
| gocyclo | 63% |
| license | 100% |
| misspell | 84% |

## Issues

### gocyclo

- `internal/sendmail/sendmail.go`
  - Line 168: cyclomatic complexity 41 for function parseArgs
  - Line 413: cyclomatic complexity 26 for function prepareMessage
  - Line 535: cyclomatic complexity 17 for function submit
- `internal/mailserver/store.go`
  - Line 907: cyclomatic complexity 40 for function (*MailServer).LoadMailsFromDirectory
  - Line 752: cyclomatic complexity 32 for function (*MailServer).parseEmailMessage
  - Line 395: cyclomatic complexity 28 for function (*MailServer).DeleteAllEmail
  - Line 46: cyclomatic complexity 26 for function (*MailServer).saveEmailToStore
- `internal/mcpserver/server_test.go`
  - Line 25: cyclomatic complexity 35 for function TestReadOnlyToolsUseMailboxSnapshots
- `internal/outgoing/smtp_security_test.go`
  - Line 384: cyclomatic complexity 24 for function startTestSMTPServer
  - Line 240: cyclomatic complexity 18 for function TestSendMailEnvelopeDeadlineCoversAllRecipients
- `internal/mailserver/mailserver_store_test.go`
  - Line 315: cyclomatic complexity 18 for function TestMailServerDeleteAllEmail
- `internal/api/api_mailcatcher_compat.go`
  - Line 73: cyclomatic complexity 17 for function mailCatcherMessageDTO
- `internal/api/relay_jobs_test.go`
  - Line 131: cyclomatic complexity 16 for function TestNativeRelayReturnsQueryableJob
- `internal/api/api.go`
  - Line 205: cyclomatic complexity 32 for function (*API).setupRoutes
- `internal/api/api_emails.go`
  - Line 473: cyclomatic complexity 31 for function applyEmailFilters
  - Line 350: cyclomatic complexity 27 for function (*API).exportEmails
  - Line 195: cyclomatic complexity 16 for function parseEmailQuery
- `internal/mailserver/attachment_migration_test.go`
  - Line 562: cyclomatic complexity 24 for function TestAttachmentMigrationDryRunVerifiesRemoteAttachments
  - Line 194: cyclomatic complexity 19 for function TestAttachmentMigrationValidatesAttachmentFreeMessages
  - Line 480: cyclomatic complexity 16 for function TestAttachmentMigrationVerifiesRemoteOnlySourcesBeforeWrites
- `internal/mailserver/mailserver_utils_test.go`
  - Line 289: cyclomatic complexity 17 for function TestSanitizeHTML
- `internal/mailserver/mailserver_events_test.go`
  - Line 79: cyclomatic complexity 17 for function TestOnWithConcurrencyBoundsHandlersBeforeStartingGoroutines
- `internal/mailserver/query_test.go`
  - Line 399: cyclomatic complexity 16 for function legacyEmailMatches
  - Line 65: cyclomatic complexity 16 for function TestMailboxQueryReturnsDetachedResults
- `internal/api/relay_jobs_persistence_test.go`
  - Line 253: cyclomatic complexity 16 for function TestRetryRelayJobPreservesQueuedWorkDuringShutdown
- `internal/mailserver/config.go`
  - Line 57: cyclomatic complexity 35 for function NewMailServerWithOptions
- `internal/api/api_mailcatcher_compat_test.go`
  - Line 35: cyclomatic complexity 38 for function TestMailCatcherRESTFacadeContract
- `internal/config/validation_test.go`
  - Line 119: cyclomatic complexity 37 for function TestValidateConfig
- `internal/config/config_test.go`
  - Line 364: cyclomatic complexity 31 for function TestDefaultConfig
  - Line 567: cyclomatic complexity 29 for function TestDefineAndResolveConfig
  - Line 512: cyclomatic complexity 29 for function TestSMTPAndS3ConfigResolution
- `internal/mailserver/mailserver_fuzz_test.go`
  - Line 214: cyclomatic complexity 27 for function FuzzParseEmailMessage
  - Line 362: cyclomatic complexity 17 for function FuzzParseEmailDate
- `internal/mcpserver/workflows_test.go`
  - Line 18: cyclomatic complexity 26 for function TestLatestAndEventDrivenWaitReturnBoundedSummaries
  - Line 313: cyclomatic complexity 25 for function TestReadOnlyResourcesAndPrompts
- `internal/api/api_emails_test.go`
  - Line 1073: cyclomatic complexity 20 for function TestAPIGetEmailPreviewsWithFilters
  - Line 2837: cyclomatic complexity 19 for function TestApplyEmailFilters
  - Line 692: cyclomatic complexity 16 for function TestAPIGetAllEmailsWithFilters
- `cmd/owlmail/main.go`
  - Line 501: cyclomatic complexity 34 for function startAPIServer
  - Line 785: cyclomatic complexity 33 for function createMailServer
  - Line 1010: cyclomatic complexity 23 for function runMCPStdio
  - Line 1143: cyclomatic complexity 17 for function main
- `internal/api/api_test.go`
  - Line 331: cyclomatic complexity 28 for function TestBasePathRouting
- `internal/attachmentstore/s3.go`
  - Line 70: cyclomatic complexity 17 for function NewS3
- `internal/webhook/payload_fuzz_test.go`
  - Line 33: cyclomatic complexity 16 for function FuzzMatchTextPattern
- `internal/mailserver/read_only.go`
  - Line 26: cyclomatic complexity 62 for function (*MailServer).RefreshReadOnlyMailbox
- `internal/mailserver/attachment_storage.go`
  - Line 139: cyclomatic complexity 23 for function (*MailServer).restoreLegacyLocalAttachmentMetadataContext
- `internal/api/relay_jobs.go`
  - Line 401: cyclomatic complexity 23 for function (*API).enqueueRelayJob
  - Line 331: cyclomatic complexity 17 for function relayFailureCategory
- `internal/sendmail/sendmail_test.go`
  - Line 26: cyclomatic complexity 21 for function TestParseArgs
- `internal/webhook/service.go`
  - Line 265: cyclomatic complexity 18 for function (*Service).runOutbox
- `internal/config/validation.go`
  - Line 28: cyclomatic complexity 85 for function ValidateConfig
  - Line 315: cyclomatic complexity 18 for function NormalizeBasePathname
- `internal/outgoing/smtp.go`
  - Line 141: cyclomatic complexity 40 for function sendMailStreamWithConfig
- `internal/webhook/config.go`
  - Line 142: cyclomatic complexity 32 for function compileTarget
- `internal/api/api_maildev_compat_test.go`
  - Line 63: cyclomatic complexity 29 for function TestMailDevRESTFacadeContractAndReadSideEffect
- `internal/webhook/payload.go`
  - Line 241: cyclomatic complexity 23 for function matchTextChunk
  - Line 174: cyclomatic complexity 16 for function matchTextPattern
- `cmd/owlmail/main_test.go`
  - Line 79: cyclomatic complexity 23 for function TestCompleteWebAuthConfig
  - Line 927: cyclomatic complexity 20 for function TestCreateMailServer
  - Line 371: cyclomatic complexity 18 for function TestSetupOutgoingConfig
  - Line 220: cyclomatic complexity 16 for function TestLoadAutoRelayRules
- `internal/mailserver/read_only_test.go`
  - Line 470: cyclomatic complexity 21 for function TestRefreshReadOnlyMailboxReappliesRepairedAttachmentMetadata
  - Line 107: cyclomatic complexity 16 for function TestRefreshReadOnlyMailboxDefersNewEventUntilEnvelopeMetadataIsReadable
- `internal/mailserver/sqlite_index_test.go`
  - Line 60: cyclomatic complexity 44 for function TestSQLiteMailboxIndexQueryAndRebuild
  - Line 250: cyclomatic complexity 19 for function TestMailServerUsesSQLiteIndexAndSynchronizesMutations
  - Line 179: cyclomatic complexity 17 for function TestSQLiteMailboxIndexPreservesWideMessageTimeRange
- `internal/api/api_config.go`
  - Line 107: cyclomatic complexity 31 for function (*API).patchOutgoingConfig
- `internal/mailserver/attachment_storage_test.go`
  - Line 755: cyclomatic complexity 19 for function TestS3ModeQuarantineRetriesRemoteCleanupAcrossRestart
  - Line 168: cyclomatic complexity 19 for function TestRemoteAttachmentDeleteFailureRecoversAcrossRestarts
  - Line 105: cyclomatic complexity 19 for function TestRemoteAttachmentLifecycleSurvivesRestart
  - Line 601: cyclomatic complexity 18 for function TestS3ModeRestoresSameSizeLegacyAttachmentsByContent
  - Line 457: cyclomatic complexity 18 for function TestS3ModeDeleteAllRetriesOnlyFailedEmails
  - Line 656: cyclomatic complexity 16 for function TestS3ModeDoesNotUpgradeAmbiguousLegacyAttachmentMetadata
- `internal/api/api_mcp_test.go`
  - Line 70: cyclomatic complexity 16 for function TestAuthenticatedMCPProtocolUsesPrefixedRoute
- `internal/mailserver/storage_transaction.go`
  - Line 30: cyclomatic complexity 27 for function (*MailServer).storeIncomingEmail
  - Line 314: cyclomatic complexity 18 for function (*MailServer).recoverStorageArtifacts
- `internal/maildev/compat.go`
  - Line 416: cyclomatic complexity 22 for function nestedValues
- `internal/mailserver/query.go`
  - Line 301: cyclomatic complexity 21 for function (compiledEmailQuery).matches
  - Line 345: cyclomatic complexity 17 for function sortEmailMatches
- `internal/config/file.go`
  - Line 20: cyclomatic complexity 21 for function LoadConfigFile
- `internal/api/relay_jobs_persistence.go`
  - Line 62: cyclomatic complexity 21 for function (*relayJobStore).load
- `internal/mcpserver/server.go`
  - Line 539: cyclomatic complexity 20 for function makeEmailQuery
- `internal/mcpserver/workflows.go`
  - Line 275: cyclomatic complexity 16 for function (*Service).waitForEmail
- `internal/mailserver/attachment_migration.go`
  - Line 240: cyclomatic complexity 44 for function preflightAttachmentMigration
  - Line 83: cyclomatic complexity 41 for function MigrateLocalAttachments

### misspell

- `docs/it/README.md`
  - Line 34: "compatibile" is a misspelling of "compatible"
  - Line 57: "definito" is a misspelling of "definition"
- `docs/it/internal/API_Refactoring_Record.md`
  - Line 13: "contributo" is a misspelling of "contribution"
  - Line 23: "completi" is a misspelling of "complexity"
  - Line 36: "contributo" is a misspelling of "contribution"
- `.github/SECURITY.de.md`
  - Line 28: "Reproduktion" is a misspelling of "Reproduction"
  - Line 47: "Reproduktion" is a misspelling of "Reproduction"
  - Line 47: "Reproduktion" is a misspelling of "Reproduction"
  - Line 63: "Konfiguration" is a misspelling of "Configuration"
- `docs/de/Sendmail.md`
  - Line 29: "Konfiguration" is a misspelling of "Configuration"
- `docs/it/Comparison-and-Migration.md`
  - Line 13: "contributo" is a misspelling of "contribution"
  - Line 24: "interfacce" is a misspelling of "interfaces"
  - Line 25: "correnti" is a misspelling of "correcting"
  - Line 27: "completi" is a misspelling of "complexity"
  - Line 40: "contributo" is a misspelling of "contribution"
- `web/app.js`
  - Line 650: "compatibile" is a misspelling of "compatible"
- `.github/CODE_OF_CONDUCT.fr.md`
  - Line 9: "insultes" is a misspelling of "insults"
- `README.fr.md`
  - Line 18: "transforme" is a misspelling of "transfer"
  - Line 52: "persistance" is a misspelling of "persistence"
  - Line 66: "marrage" is a misspelling of "marriage"
  - Line 138: "exemple" is a misspelling of "example"
  - Line 152: "exemple" is a misspelling of "example"
  - Line 217: "commandes" is a misspelling of "commands"
  - Line 223: "Origines" is a misspelling of "Origins"
  - Line 230: "Origines" is a misspelling of "Origins"
  - Line 240: "existant" is a misspelling of "existent"
  - Line 246: "marrage" is a misspelling of "marriage"
  - Line 284: "marrage" is a misspelling of "marriage"
  - Line 295: "marrage" is a misspelling of "marriage"
  - Line 299: "marrage" is a misspelling of "marriage"
  - Line 300: "exemple" is a misspelling of "example"
  - Line 532: "marrage" is a misspelling of "marriage"
  - Line 539: "marrage" is a misspelling of "marriage"
  - Line 616: "existant" is a misspelling of "existent"
- `docs/de/internal/API_Refactoring_Record.md`
  - Line 156: "Konfiguration" is a misspelling of "Configuration"
  - Line 157: "Konfiguration" is a misspelling of "Configuration"
  - Line 158: "Konfiguration" is a misspelling of "Configuration"
  - Line 159: "Konfiguration" is a misspelling of "Configuration"
  - Line 204: "Konfiguration" is a misspelling of "Configuration"
  - Line 205: "Konfiguration" is a misspelling of "Configuration"
  - Line 206: "Konfiguration" is a misspelling of "Configuration"
- `docs/fr/internal/API_Refactoring_Record.md`
  - Line 3: "constitue" is a misspelling of "constitutes"
  - Line 13: "contribuer" is a misspelling of "contribute"
  - Line 23: "complets" is a misspelling of "completes"
  - Line 27: "contribuer" is a misspelling of "contribute"
- `.github/CODE_OF_CONDUCT.de.md`
  - Line 9: "Empathie" is a misspelling of "Empathize"
- `.github/CONTRIBUTING.fr.md`
  - Line 1: "Contribuer" is a misspelling of "Contribute"
- `.github/SECURITY.it.md`
  - Line 28: "possibile" is a misspelling of "possible"
  - Line 73: "infrastrutture" is a misspelling of "infrastructure"
- `README.it.md`
  - Line 20: "limitato" is a misspelling of "limitation"
  - Line 26: "Convertire" is a misspelling of "Converter"
  - Line 28: "automatico" is a misspelling of "automation"
  - Line 59: "integrato" is a misspelling of "integration"
  - Line 61: "distribuito" is a misspelling of "distribution"
  - Line 68: "destinato" is a misspelling of "destination"
  - Line 207: "Comando" is a misspelling of "Commando"
  - Line 223: "incompatibile" is a misspelling of "incompatible"
  - Line 235: "compatibile" is a misspelling of "compatible"
  - Line 250: "certificato" is a misspelling of "certification"
  - Line 258: "certificato" is a misspelling of "certification"
  - Line 265: "automatico" is a misspelling of "automation"
  - Line 266: "automatico" is a misspelling of "automation"
  - Line 267: "automatico" is a misspelling of "automation"
  - Line 277: "certificato" is a misspelling of "certification"
  - Line 396: "configurato" is a misspelling of "configuration"
  - Line 457: "incluse" is a misspelling of "includes"
  - Line 460: "configurato" is a misspelling of "configuration"
  - Line 487: "certificato" is a misspelling of "certification"
  - Line 491: "Automatico" is a misspelling of "Automation"
  - Line 494: "automatico" is a misspelling of "automation"
  - Line 503: "automatico" is a misspelling of "automation"
  - Line 565: "comando" is a misspelling of "commando"
  - Line 580: "compatibile" is a misspelling of "compatible"
  - Line 587: "correnti" is a misspelling of "correcting"
  - Line 688: "concesso" is a misspelling of "concession"
- `docs/fr/README.md`
  - Line 33: "persistance" is a misspelling of "persistence"
  - Line 59: "Contribuer" is a misspelling of "Contribute"
- `docs/it/Sendmail.md`
  - Line 1: "compatibile" is a misspelling of "compatible"
- `.github/SECURITY.fr.md`
  - Line 103: "responsable" is a misspelling of "responsible"
- `README.de.md`
  - Line 42: "Deterministische" is a misspelling of "Deterministic"
  - Line 55: "Konfiguration" is a misspelling of "Configuration"
  - Line 69: "Produktions" is a misspelling of "Productions"
  - Line 69: "higer" is a misspelling of "higher"
  - Line 399: "Konfiguration" is a misspelling of "Configuration"
  - Line 424: "Konfiguration" is a misspelling of "Configuration"
  - Line 425: "Konfiguration" is a misspelling of "Configuration"
  - Line 426: "Konfiguration" is a misspelling of "Configuration"
  - Line 447: "Konfiguration" is a misspelling of "Configuration"
  - Line 448: "Konfiguration" is a misspelling of "Configuration"
  - Line 449: "Konfiguration" is a misspelling of "Configuration"
  - Line 538: "Konfiguration" is a misspelling of "Configuration"
  - Line 600: "Konfiguration" is a misspelling of "Configuration"
  - Line 714: "historisch" is a misspelling of "historic"
- `docs/README.md`
  - Line 106: "compatibile" is a misspelling of "compatible"
- `docs/fr/Comparison-and-Migration.md`
  - Line 13: "contribuer" is a misspelling of "contribute"
  - Line 19: "essentiels" is a misspelling of "essentials"
  - Line 27: "complets" is a misspelling of "completes"
  - Line 31: "contribuer" is a misspelling of "contribute"
- `docs/fr/Sendmail.md`
  - Line 43: "certificats" is a misspelling of "certificates"
  - Line 48: "journalis" is a misspelling of "journals"

---

_Generated by [Go Report Card](https://github.com/soulteary/goreportcard-action)._

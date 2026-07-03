# Security & Compliance Rules

This is a **living rule set**. Logres is self-hosted commerce software: operators
run their own instance, hold their own customer data, and inherit their own
regulatory exposure. Our job is to make the software *capable* of compliant
operation — correct primitives, honest defaults, no dark patterns — not to claim
certification on the operator's behalf.

**Working rule:** whenever a change touches authentication, personal data,
money, logging, or a third-party integration, add or update the matching control
below and cite it in the PR. If a standard has no concrete control yet, that is a
gap to file, not a line to leave blank.

Legend for control status: `enforced` (in code + CI), `partial` (mechanism
exists, coverage incomplete), `planned` (design agreed, not built).

---

## 0. Cross-cutting engineering controls (RFC-based)

These are the primitives every feature reuses. They are the concrete,
verifiable backbone the regulatory sections below lean on.

| # | Control | Standard | Status |
|---|---------|----------|--------|
| X1 | Passwords hashed with **Argon2id** (RFC 9106) — never bcrypt/sha/plaintext. Parameters checked in on introduction. | OWASP ASVS 2.4 | planned |
| X2 | Session/access tokens are **JWT** with the algorithm pinned server-side (reject `alg=none`; RFC 7519/8725); secret ≥ 256-bit; `exp` always set. | RFC 8725 | planned |
| X3 | Refresh tokens are opaque, stored **hashed** (SHA-256), single-use, with reuse-detection that revokes the family (RFC 6749 §10.4). | RFC 6749 | planned |
| X4 | Auth cookies use `__Host-` prefix, `HttpOnly`, `Secure`, `SameSite=Lax`. | OWASP | planned |
| X5 | Every request carries a **UUIDv7** request id (RFC 9562); it appears in logs and in the `instance` field of error responses. | RFC 9562 | partial (uuidv7 in models) |
| X6 | API errors use **`application/problem+json`** (RFC 9457); success uses the `{code, message, data}` envelope. Internal error text/stack traces never reach clients. | RFC 9457 | planned |
| X7 | All SQL is **parameterised**; identifiers interpolated into SQL are validated against an allowlist (see `migrationx` table-name check). | CWE-89 | enforced |
| X8 | Money is `shopspring/decimal` end to end — never `float64`. Currency is explicit (ISO 4217). | — | enforced (models) |
| X9 | Structured logging via `slog`; **PII and secrets are redacted before the log call**, never after. | GDPR Art. 25 | planned |
| X10 | Dependencies scanned every CI run: `govulncheck` (CVEs) + `gosec` (SAST); a finding is a red build. | ISO 27001 A.8.8 | enforced |

---

## 1. ISO/IEC 27001:2022 (ISMS)

Operator owns the ISMS; we supply technical controls that map to Annex A.

- **A.8.8 Vulnerability management** — `govulncheck` + `gosec` + Dependabot are wired into CI (`enforced`). No merge with an open high finding.
- **A.8.24 Cryptography** — crypto choices are centralised (X1–X4), never hand-rolled per feature. Use `crypto/*` and vetted libraries only.
- **A.8.15 Logging** — security-relevant events (admin login, password change, permission change, data export/delete) emit an audit log entry with actor, action, target id, timestamp, request id. `planned`.
- **A.8.5 Secure authentication** — X1–X4.
- **A.5.34 Privacy & PII protection** — see §2.
- **Data at rest** — SQLite file permissions `0600`; document operator guidance for full-disk / volume encryption. `planned`.

## 2. GDPR (EU) & CCPA/CPRA (California)

Both are data-subject-rights regimes; the software must make rights *executable*.

- **Lawful basis / data minimisation** — a new column storing personal data (name, email, phone, address, IP) requires a one-line justification in the PR and must map to a documented purpose. No "collect it just in case".
- **Right to erasure / deletion** — every PII-bearing entity needs a deletion path. Soft-delete (`deleted_at`) is not erasure; a hard-delete/anonymise routine is required before that entity ships to production. Adding PII storage without a deletion path is a **blocker**.
- **Right to access / portability** — customer data must be exportable in a machine-readable form (JSON). `planned`.
- **Privacy by design (Art. 25)** — PII never in logs (X9); analytics/telemetry off by default and opt-in.
- **Consent** — first-time collection of a new PII category requires explicit UI consent copy.
- **CCPA "Do Not Sell/Share"** — Logres does not sell data; keep it that way. Any third-party data flow (payments, shipping, email) is disclosed and configurable.

## 3. PCI DSS (card data)

**Primary control: descope.** Logres must **not** store, process, or transmit
Primary Account Numbers.

- Card entry uses a PCI-compliant processor via **tokenisation / hosted fields / redirect** (e.g. provider iframe). We store only the processor's token and last-4/brand for display. `planned`.
- No PAN, CVV, or full track data in the database, logs, or error payloads — ever. A code path that could persist card data is a **blocker**.
- Payment webhooks verify provider signatures before acting.
- This keeps operators in **SAQ A** territory rather than full DSS.

## 4. UU PDP — Law No. 27/2022 (Indonesia)

Indonesia's PDP law tracks GDPR closely; the §2 controls largely satisfy it.

- **Consent & purpose limitation** — as §2. Support Bahasa Indonesia consent copy.
- **Data subject rights** — access, correction, erasure, withdrawal of consent (reuse §2 mechanisms).
- **Breach notification** — retain the audit/event log (A.8.15) needed to reconstruct an incident within the 3×24-hour notification window. `planned`.
- **Cross-border transfer** — because Logres is self-hosted, data residency is the operator's choice; document which integrations (payment, email) egress data abroad.

## 5. OJK / Bank Indonesia (Indonesian financial services)

Applies only if an operator runs regulated financial/payment features.

- **Do not become a payment processor.** Integrate licensed providers; keep fund movement outside Logres (see §3).
- **Auditability** — immutable, append-only audit log for anything money-adjacent (order state changes, refunds); order status transitions go through the (planned) FSM with recorded actor + reason.
- **Data residency** — flag integrations that move data outside Indonesia so a regulated operator can disable them.

## 6. HIPAA (US health data)

Out of default scope — Logres is not a health product. If an operator sells
regulated health goods:

- Any PHI is treated at the §2 PII bar **plus** encryption at rest and access
  logging (A.8.15).
- No PHI in logs, URLs, or analytics. This is advisory until a real use case exists.

## 7. FATF AML / CFT

Relevant only for high-value or regulated goods.

- **Screening hooks** — provide extension points for sanctions/PEP screening at checkout rather than baking a specific vendor in. `planned`.
- **Record keeping** — customer identity + transaction records retained per the audit-log control; retention window is operator-configurable.
- **Suspicious activity** — velocity/anomaly signals surface to the operator; Logres does not auto-file reports.

---

## Definition of done for a security-relevant change

1. The matching control above is added/updated and its status is honest.
2. Inputs validated at the boundary (`go-playground/validator`); no raw user input in SQL, logs, or templates.
3. `make ci` green: `gosec`, `govulncheck`, `-race`, coverage gate all pass.
4. No PII/secret can reach a log line or a client error body.
5. New third-party data flow is disclosed and operator-configurable.

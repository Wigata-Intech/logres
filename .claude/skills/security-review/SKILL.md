---
name: security-review
description: Review Logres changes against the project's security & compliance standards (RFCs, ISO 27001, GDPR/CCPA/UU PDP, PCI descope, OJK/BI, HIPAA, FATF) before merge. Use when a diff touches auth, personal data, money, logging, SQL, or a third-party integration.
---

# Security review — Logres

Review the change against `.claude/rules/security-compliance.md`, which is the
source of truth. This skill is the *procedure*; that file is the *content*.

## When to run
Any diff that touches authentication/authorization, personal data (PII), money,
logging, SQL, cryptography, or a new dependency / external data flow.

## Procedure
1. Identify which controls the diff touches (rules Section 0 cross-cutting RFCs, Section 1–Section 7 regimes).
2. Verify each applicable control holds — cite it by id (e.g. `X2`, GDPR erasure).
3. Flag gaps as findings: `[BLOCKER]` (exploitable / illegal), `[MAJOR]` (standard violated), `[MINOR]` (hardening).
4. Confirm the change updated the rules file if it introduced or altered a control.

## Non-negotiables (fast checks)
- No PAN/CVV stored or logged (PCI descope). No PII or secrets in any log line or client error body.
- SQL parameterised; interpolated identifiers allowlist-validated.
- Passwords Argon2id; JWT alg pinned + `exp` set; refresh tokens hashed + single-use.
- Auth cookies `__Host-` HttpOnly Secure SameSite. New PII column has a documented purpose + deletion path.
- `make ci` green: gosec, govulncheck, `-race`, coverage gate.

## Contract: keep the standards current
**Whenever a change adopts a new RFC or standard (a new auth RFC, a crypto
primitive, a regulatory obligation), you MUST update BOTH:**
1. `.claude/rules/security-compliance.md` — add/adjust the control + its status.
2. This skill's fast-checks list, if the new standard warrants a merge-time check.

A new standard that isn't reflected here is an incomplete change.

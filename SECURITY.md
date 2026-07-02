# Security Policy

## Reporting a Vulnerability

At Nightona, we take security seriously. If you believe you have found a security vulnerability in any Nightona-owned repository or service, please report it responsibly.

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report vulnerabilities privately through [GitHub's private vulnerability reporting](https://github.com/nightona-co/nightona/security/advisories/new):

**https://github.com/nightona-co/nightona/security/advisories/new**

Please include:

- Description of the vulnerability
- Steps to reproduce
- Impact assessment
- Any relevant screenshots or proof-of-concept

Nightona is a community-maintained project. We aim to acknowledge receipt within 5 business days and provide an initial assessment as soon as possible after that, on a best-effort basis.

## Scope

The following assets and areas are in scope for vulnerability reports:

- **Nightona platform** — the self-hostable platform in this repository, including the API, dashboard, proxy, runner, and management interfaces
- **API and SDK** — all documented and undocumented API endpoints, client SDKs
- **Sandbox runtime isolation** — escape from sandbox to host, cross-tenant access, isolation boundary bypasses
- **Authentication and authorization** — SSO, API key management, session handling, privilege escalation across accounts or organizations
- **Secrets management** — scoped secret injection, unauthorized access to secrets, leakage across sandbox boundaries
- **Public GitHub repositories** — the [nightona-co/nightona](https://github.com/nightona-co/nightona) repository

## Excluded Submission Types

The following categories are excluded from this program. Reports in these categories will be closed without further assessment unless they demonstrate impact beyond what is described.

1. **In-sandbox privilege escalation, root access, or capability use** — Nightona sandboxes provide full root access within user-namespace isolation by design. Findings that chain to host escape or cross-sandbox access remain in scope.
2. **Findings within the reporter's own sandbox** that do not demonstrate impact beyond that sandbox's isolation boundary.
3. **Denial of service** — DoS, DDoS, resource exhaustion, volumetric testing, or network flooding.
4. **Rate limiting observations** that do not demonstrate resource exhaustion, financial impact, or abuse potential.
5. **Social engineering** — phishing, vishing, pretexting, or any form of social engineering targeting Nightona employees or users.
6. **Physical security testing** of offices, data centers, or personnel.
7. **Upstream Daytona sites and services** — findings against daytona.io, docs.daytona.io, app.daytona.io, or any other infrastructure operated by upstream Daytona. Nightona does not operate these; report them to the upstream Daytona project.
8. **Third-party services** — vulnerabilities in services or platforms not owned or operated by Nightona.
9. **Known public files or directories** — e.g., robots.txt, .well-known, or other intentionally public resources.
10. **DNSSEC or TLS cipher suite configuration suggestions** without a demonstrated exploit path.
11. **Missing Secure/HTTPOnly flags** on non-sensitive cookies.
12. **CSRF on unauthenticated or public-facing forms.**
13. **Outdated browsers and platforms** — vulnerabilities only affecting unpatched or end-of-life software.
14. **Automated scan output** — reports generated solely by automated tools without validated proof of impact.
15. **Best practice recommendations** without demonstrable security impact.
16. **Spam or service degradation** — testing that results in sending unsolicited messages or degradation of service to other users.

## Supported Versions

We accept vulnerability reports for the latest stable release of Nightona.

## Safe Harbor

Nightona supports safe harbor for security researchers who act in good faith and in accordance with this policy.

We will not pursue legal action against researchers who:

- Make a good-faith effort to avoid privacy violations, data destruction, and service disruption
- Only access data to the extent necessary to demonstrate the vulnerability
- Do not exfiltrate, retain, or disclose any user data encountered during research
- Report findings promptly through the channels listed above
- Do not disclose findings publicly before coordinated resolution (see Disclosure Timeline below)
- Comply with all applicable laws

If legal action is initiated by a third party against a researcher for activities conducted in accordance with this policy, we will take steps to make it known that the research was authorized.

This safe harbor applies to all Nightona services and assets listed in the Scope section.

## Disclosure Timeline

We follow a coordinated disclosure process:

- **90 days** — We target remediation within 90 days of a validated report. Complex issues may require additional time, and we will communicate timelines transparently.
- **30 days post-patch** — After a fix is released, we ask that researchers wait 30 days before public disclosure to allow users to update.
- **No response** — If we fail to acknowledge or respond to a report within 90 days, the researcher may proceed with public disclosure after providing 14 days advance written notice through the same [GitHub private vulnerability reporting channel](https://github.com/nightona-co/nightona/security/advisories/new).

## Rewards

Nightona is a community-maintained open-source project and does not currently operate a paid bug bounty program. Valid, original findings are credited in the resulting security advisory and release notes. Duplicate reports are credited to the first submission.

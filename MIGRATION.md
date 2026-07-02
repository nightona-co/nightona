# Migrating from Daytona v0.190.0 to Nightona

This guide is for operators of a **self-hosted Daytona v0.190.0 deployment** (Docker Compose or Kubernetes) who want to move to Nightona, and for developers whose applications use the Daytona SDKs or CLI against that deployment.

Nightona starts from the exact Daytona v0.190.0 codebase, so the runtime behavior, API surface, and data model are unchanged. What changes is **naming**: binaries, packages, module paths, environment variables, and configuration locations. This document lists every rename and the compatibility shims that ease the transition.

> **Bootstrap-phase disclaimer**: Nightona is mid-rebrand. Some of the compatibility behavior described here may not yet be present in every component of a given build, and Nightona has not yet published its own packages or container images. Items marked **needs verification** must be checked against your actual build before migrating a production deployment.

## TL;DR checklist

1. Back up your PostgreSQL database, Redis data, and object storage before switching.
2. Point your deployment at Nightona builds (built from this repository — no prebuilt Nightona artifacts exist yet).
3. Keep your existing `DAYTONA_*` environment variables — they continue to work — and migrate to `NIGHTONA_*` at your own pace.
4. Update SDK dependencies to the Nightona package names when they are published; until then, build the SDKs from source.
5. Set the API URL explicitly in SDK configuration or via `NIGHTONA_API_URL` — the default is no longer the Daytona SaaS.
6. Review the [Persisted state caveats](#persisted-state-caveats-needs-verification) section before starting the new stack against old data.

## Environment variables

All `DAYTONA_*` environment variables have `NIGHTONA_*` equivalents. A compatibility shim keeps the legacy names working:

- **`NIGHTONA_*` is preferred.** New documentation, examples, and defaults use the `NIGHTONA_` prefix.
- **`DAYTONA_*` is still honored.** When a component looks up `NIGHTONA_FOO` and it is unset, it falls back to the legacy `DAYTONA_FOO` twin. If both are set, `NIGHTONA_*` wins.
- In the Python SDK this also applies to values loaded from `.env` / `.env.local` files (precedence: runtime env, then `.env.local`, then `.env`; within each level `NIGHTONA_*` beats `DAYTONA_*`).

Common variables:

| Legacy (Daytona) | Preferred (Nightona) |
| ---------------------------- | ------------------------------ |
| `DAYTONA_API_KEY` | `NIGHTONA_API_KEY` |
| `DAYTONA_API_URL` | `NIGHTONA_API_URL` |
| `DAYTONA_SERVER_URL` (deprecated alias) | `NIGHTONA_SERVER_URL` (deprecated alias) |
| `DAYTONA_JWT_TOKEN` | `NIGHTONA_JWT_TOKEN` |
| `DAYTONA_ORGANIZATION_ID` | `NIGHTONA_ORGANIZATION_ID` |
| `DAYTONA_TARGET` | `NIGHTONA_TARGET` |
| `DAYTONA_CONFIG_DIR` | `NIGHTONA_CONFIG_DIR` |

> **Needs verification**: the fallback shim is being rolled out per component during the rebrand. Verify it in the specific SDK/app version you deploy (the Python SDK implements it in `libs/sdk-python/src/nightona/_utils/env.py`). If a component in your build only reads `NIGHTONA_*`, set both names during the transition — that is safe in all cases.

## CLI configuration directory

The CLI (now installed as `nightona` instead of `daytona`) stores its configuration under a `nightona` directory inside your platform's user config directory (e.g. `~/.config/nightona` on Linux, `~/Library/Application Support/nightona` on macOS), or wherever `NIGHTONA_CONFIG_DIR` points.

For migration, the legacy Daytona config location (e.g. `~/.daytona` / `~/.config/daytona`) is read as a fallback when the Nightona config directory does not exist yet, so existing profiles and API keys keep working (**needs verification** — the fallback is part of the rebrand rollout; check your CLI build).

If your build does not include the fallback, migrate manually — the file format is unchanged:

```bash
# Linux example; adjust paths for your platform
mkdir -p ~/.config/nightona
cp -r ~/.config/daytona/* ~/.config/nightona/ 2>/dev/null || cp -r ~/.daytona/* ~/.config/nightona/
```

## Package renames

Nightona-named packages are **not published yet** (see [PUBLISHING.md](PUBLISHING.md)); until they are, build from this repository. The planned coordinates:

| Ecosystem | Daytona v0.190.0 (upstream, published) | Nightona (this repo / planned) |
| --------- | -------------------------------------- | ------------------------------ |
| npm (TypeScript SDK) | `@daytonaio/sdk` | `@nightona/sdk` |
| PyPI (Python SDK) | `daytona` | `nightona` |
| RubyGems (Ruby SDK) | `daytona` | `nightona` |
| Go module (Go SDK) | `github.com/daytonaio/daytona/libs/sdk-go` | `github.com/Amartuvshins0404/nightona/libs/sdk-go` |
| Java / Maven (Java SDK) | `io.daytona:sdk` | `io.nightona:sdk` (with `io.nightona:api-client`, `io.nightona:toolbox-api-client`) |
| CLI binary | `daytona` | `nightona` |

Import paths and top-level namespaces change accordingly, e.g.:

```python
# before
from daytona import Daytona
# after
from nightona import Nightona
```

```typescript
// before
import { Daytona } from '@daytonaio/sdk'
// after
import { Nightona } from '@nightona/sdk'
```

## SDK default API URL

The Daytona SDKs defaulted to the Daytona SaaS endpoint (`https://app.daytona.io/api`). Nightona has no SaaS, so the SDK default is now the local self-hosted API:

```
http://localhost:3000/api
```

If your API is reachable elsewhere, set it explicitly — either `NIGHTONA_API_URL` (or legacy `DAYTONA_API_URL`) in the environment, or `apiUrl` / `api_url` in the SDK constructor configuration.

> **Needs verification**: the new default has landed in the Python SDK; confirm the TypeScript, Go, Ruby, and Java SDKs in your build no longer default to `https://app.daytona.io/api` before relying on the implicit default. Setting the URL explicitly avoids the question entirely.

## Container images

Nightona does not yet operate its own image registries. Deployments continue to pull the **upstream published Daytona images** where images are referenced at runtime, most notably the default sandbox snapshot image:

- `daytonaio/sandbox:<tag>` (e.g. `DEFAULT_SNAPSHOT=daytonaio/sandbox:0.5.0-slim` in `docker/docker-compose.yaml`)
- `ghcr.io/daytonaio/...` images referenced by charts/compose files

These are unmodified upstream artifacts pulled from Daytona's registries and will keep working as long as upstream keeps them published. Once Nightona publishes its own images (see the [roadmap](README.md#roadmap)), the references will be switched and this guide updated. If you need to insulate yourself from upstream registry availability now, mirror the images into your own registry and override the references.

## Persisted state caveats (needs verification)

The case-preserving rename was applied across the whole source tree, which means some **identifier strings that end up in persisted state** may differ between a Daytona v0.190.0 deployment's data and what the renamed Nightona code expects. Known examples found in the tree:

- **Seeded database identifiers**: the admin user id is now `nightona-admin` (`apps/api/src/app.service.ts`), and a migration references `id = 'nightona-admin'`; an existing Daytona database contains `daytona-admin`.
- **Kafka consumer group**: the API consumer `groupId` is now `nightona` (`apps/api/src/main.ts`); upstream used `daytona`. A group id change resets consumer offsets for that group.
- **Queue / Redis / job identifiers**: queue names, Redis key prefixes, lock names, and scheduled-job identifiers that embed the product name may have changed the same way, orphaning in-flight jobs or cached state from the old deployment.
- **TypeORM migration history**: the `migrations` table records executed migrations by name/timestamp; these are expected to be unchanged, but verify that the renamed codebase does not attempt to re-run or rename any migration.

**Before migrating production data**, diff the identifier strings your deployment persists (database rows, Redis keys, Kafka topics/groups, object-storage paths) against what the Nightona build emits, and prepare data-fix SQL / key-rename steps where they diverge. None of this is verified end-to-end yet — treat a migration on real data as untested until the Nightona project publishes a validated migration script. Recommended approach:

1. Take full backups (PostgreSQL dump, Redis RDB/AOF, object storage).
2. Restore into a staging environment and start the Nightona stack against it.
3. Verify: admin login, existing organizations/users, sandbox listing, snapshot operations, and background job processing.
4. Only then cut over production, keeping the backups until the new stack has run cleanly.

## Getting help

- Issues: https://github.com/Amartuvshins0404/nightona/issues
- Discussions: https://github.com/Amartuvshins0404/nightona/discussions

If you hit a migration problem not covered here — especially in the persisted-state area — please open an issue so this guide can be corrected.

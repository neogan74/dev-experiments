# GraphQL Data Utilities

This directory now contains two Go command line helpers for seeding and reporting on a GraphQL API. Both tools share a minimal HTTP-based GraphQL client and use [`gofakeit`](https://github.com/brianvoe/gofakeit) to fabricate member data.

## Prerequisites

- Go 1.21 or newer.
- Access to a GraphQL API that exposes mutations for creating members and queries for listing them. The default operations expect a schema similar to Hasura's `MemberInput`/`members` patterns—override them with the flags shown below if your schema differs.
- Run `go mod tidy` the first time to download dependencies.

## Seed Fake Members

```
go run ./cmd/loader \
  -endpoint https://your-graphql-endpoint/v1/graphql \
  -auth-token <optional bearer token> \
  -count 50
```

Key flags and environment variables:

- `-endpoint` (or `GRAPHQL_ENDPOINT`) – GraphQL HTTP endpoint.
- `-auth-token` (or `GRAPHQL_BEARER_TOKEN`) – Bearer token for authenticated servers.
- `-count` (or `MEMBER_COUNT`) – Number of fake members to send (default 25).
- `-mutation-file` – Path to a `.graphql` file with a custom mutation if the default `createMember` mutation is incompatible.
- `-dry-run` – Generate and print payloads without sending them.

Each generated member includes contact details, demographic fields, signup timestamp, and a plan tier-specific monthly spend.

## Build Member Reports

```
go run ./cmd/report \
  -endpoint https://your-graphql-endpoint/v1/graphql \
  -auth-token <optional bearer token> \
  -limit 500 \
  -top 10
```

The reporter paginates through members, summarises plan/country distributions, calculates average age and monthly spend, and highlights the top spenders. Flags and environment variables:

- `-limit` (or `REPORT_LIMIT`) – Hard cap on the number of members to pull (0 = no cap).
- `-page-size` (or `REPORT_PAGE_SIZE`) – Page size for the GraphQL query (default 100).
- `-query-file` – Override the default `members` query with a custom `.graphql` document.
- `-json` – Emit the report summary as formatted JSON.
- `-top` (or `REPORT_TOP`) – Number of top spenders shown in the summary.

## Directory Layout

- `cmd/loader` – CLI that seeds the API with fake member data.
- `cmd/report` – CLI that fetches members and renders a textual or JSON report.
- `internal/graphqlclient` – Shared lightweight GraphQL HTTP client.

Feel free to tailor the GraphQL operations or extend the summariser to match the shape of your production schema.


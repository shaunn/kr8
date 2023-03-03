domainFilters:
- example.com
- example.io
- bad
- example.work
extraEnv:
  CF_API_EMAIL: admin@example.com
  CF_API_KEY: some_key
logLevel: debug
nameOverride: external-dns
policy: sync
provider: cloudflare
rbac:
  create: true
registry: txt
txtOwnerId: cursedone


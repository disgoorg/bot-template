[![Docker](https://github.com/DisgoOrg/bot-template/actions/workflows/docker-build.yml/badge.svg)](https://github.com/DisgoOrg/bot-template/actions/workflows/docker-build.yml)
[![Test](https://github.com/DisgoOrg/bot-template/actions/workflows/go-test.yml/badge.svg)](https://github.com/DisgoOrg/bot-template/actions/workflows/go-test.yml)

# bot-template

This is a simple bot template for creating a bot which includes a config, postgresql database, slash commands and event listeners.

Optional CLI Flags:
- `--sync-commands=true`: Synchronize commands with the discord.
- `--sync-db=true`: Synchronize database.
- `--exit-after=true`: Exit after db & commands sync.

# Release Telegram Action

Clean up and send github auto generated release notes to a telegram channel

Currently it only works on linux runners. The cleanup is rudimentary, and specific to github generated release notes (the ones created using `.github/release.yml`)

### Example workflow

```yaml
name: Post changelog
on:
  release:
    types: published

jobs:
  post-changelog:
    runs-on: ubuntu-latest
    steps:
      - name: Send changelog to telegram channel
        uses: tyagdit/release-telegram-action@latest
        with:
          token: ${{ secrets.TELEGRAM_BOT_TOKEN }}
          channel: ${{ secrets.CHANGELOG_CHANNEL_ID }}
          changelog: ${{ github.event.release.body }}
```

The release notes are obtained from `github.event.release.body` which is available only on the release event

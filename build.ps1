param(
    [string]$Version = "dev"
)

New-Item -Path "outputs" -ItemType Directory -Force

go build `
  -ldflags "-X main.version=$Version" `
  -o .\outputs\youtube-watcher-cli.exe .
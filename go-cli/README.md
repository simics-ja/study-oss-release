# OSS Release Demo CLI

A simple command-line tool demonstrating binary distribution via GitHub Releases.

## Installation

### Download Pre-built Binaries

Download the appropriate binary for your platform from the [Releases](https://github.com/simics-ja/study-oss-release/releases) page.

#### macOS
```bash
# Intel Mac
curl -LO https://github.com/simics-ja/study-oss-release/releases/latest/download/oss-release-demo-cli-darwin-amd64
chmod +x oss-release-demo-cli-darwin-amd64
./oss-release-demo-cli-darwin-amd64 --version

# Apple Silicon (M1/M2)
curl -LO https://github.com/simics-ja/study-oss-release/releases/latest/download/oss-release-demo-cli-darwin-arm64
chmod +x oss-release-demo-cli-darwin-arm64
./oss-release-demo-cli-darwin-arm64 --version
```

#### Linux
```bash
# AMD64
curl -LO https://github.com/simics-ja/study-oss-release/releases/latest/download/oss-release-demo-cli-linux-amd64
chmod +x oss-release-demo-cli-linux-amd64
./oss-release-demo-cli-linux-amd64 --version

# ARM64
curl -LO https://github.com/simics-ja/study-oss-release/releases/latest/download/oss-release-demo-cli-linux-arm64
chmod +x oss-release-demo-cli-linux-arm64
./oss-release-demo-cli-linux-arm64 --version
```

#### Windows
```powershell
# Download from releases page
Invoke-WebRequest -Uri https://github.com/simics-ja/study-oss-release/releases/latest/download/oss-release-demo-cli-windows-amd64.exe -OutFile oss-release-demo-cli.exe
.\oss-release-demo-cli.exe --version
```

### Build from Source

```bash
# Clone repository
git clone https://github.com/simics-ja/study-oss-release.git
cd study-oss-release/go-cli

# Build for current platform
make build

# Build for all platforms
make build-all
```

## Usage

```bash
# Show version
./oss-release-demo-cli --version

# Show system information
./oss-release-demo-cli --system

# Greet someone
./oss-release-demo-cli --greet World

# Greet in uppercase
./oss-release-demo-cli --greet World --upper

# Output in JSON format
./oss-release-demo-cli --system --json
```

## Development

### Requirements

- Go 1.20 or higher
- Make (for using Makefile)

### Building

```bash
# Build for current platform
go build -o oss-release-demo-cli

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o oss-release-demo-cli-linux-amd64

# Cross-compile for Windows
GOOS=windows GOARCH=amd64 go build -o oss-release-demo-cli-windows-amd64.exe

# Cross-compile for macOS
GOOS=darwin GOARCH=amd64 go build -o oss-release-demo-cli-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o oss-release-demo-cli-darwin-arm64
```

## License

MIT
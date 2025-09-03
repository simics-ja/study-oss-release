# OSS Release Demo Docker Application

A simple Node.js web application demonstrating Docker image publication to GitHub Container Registry.

## Local Development

```bash
# Run locally
npm start

# Build Docker image
docker build -t oss-release-demo-app .

# Run Docker container
docker run -p 3000:3000 oss-release-demo-app
```

## GitHub Container Registry

This image is published to GitHub Container Registry (ghcr.io).

### Pull the image

```bash
docker pull ghcr.io/simics-ja/oss-release-demo-app:latest
```

### Run the container

```bash
docker run -p 3000:3000 ghcr.io/simics-ja/oss-release-demo-app:latest
```

## Endpoints

- `/` - Main web interface with system information
- `/health` - Health check endpoint returning JSON status

## Environment Variables

- `PORT` - Server port (default: 3000)

## License

MIT
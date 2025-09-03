# GitHub Container Registry Setup Guide

## Prerequisites

1. GitHub Personal Access Token (PAT) with `write:packages` scope
2. Docker installed locally

## Manual Publishing Steps

### 1. Create Personal Access Token

1. Go to GitHub Settings > Developer settings > Personal access tokens
2. Generate new token (classic) with scopes:
   - `write:packages`
   - `read:packages`
   - `delete:packages` (optional)

### 2. Authenticate with GHCR

```bash
# Login to GitHub Container Registry
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin
```

### 3. Build and Tag Image

```bash
# Build the image
docker build -t oss-release-demo-app .

# Tag for GHCR
docker tag oss-release-demo-app ghcr.io/simics-ja/oss-release-demo-app:latest
docker tag oss-release-demo-app ghcr.io/simics-ja/oss-release-demo-app:1.0.0
```

### 4. Push to GHCR

```bash
# Push the images
docker push ghcr.io/simics-ja/oss-release-demo-app:latest
docker push ghcr.io/simics-ja/oss-release-demo-app:1.0.0
```

## Automated Publishing with GitHub Actions

See `.github/workflows/docker-publish.yml` for automated publishing on releases.

## Making Image Public

1. Go to your GitHub profile > Packages
2. Click on the package
3. Package settings > Change visibility > Public

## Using the Published Image

```bash
# No authentication needed for public images
docker pull ghcr.io/simics-ja/oss-release-demo-app:latest
docker run -p 3000:3000 ghcr.io/simics-ja/oss-release-demo-app:latest
```
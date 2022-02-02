# edgex-snap-metadata
This repo contains the metadata of EdgeX snaps along with utility scripts to generate dummy snaps that can used to update the metadata of those snaps on https://snapcraft.io.

Utility scripts:
* [create-snapcraft](create-snapcraft.go): Reads [template.yaml](template.yaml) and populates it with the [metadata](metadata) of a snap.
* [upload-metadata](upload-metadata.sh): Builds a snap out of the resulting `snap/snapcraft.yaml` file and uploads it to the store, replacing any existing data.

## Metadata

The metadata directory contains the metadata of snaps.

Each `.md` file contains the summary and description, separated by a markdown horizontal line (`---`):
```
summary
---
description
```

Notes on the formatting:
* Use `**bold**` for section header
* Section headers do not need a trailing colon
* Use double-spaces at the end of a line to make a line-break
* Use two line-break to start a new paragraph in text or after the section headers
* URLs get converted to hyperlinks automatically. Adding hyperlinks as `[title](url)` work on Github and in the snapcraft.io preview but will not be rendered on the final snapcraft.io listing!

## Requirements
Snapcraft and Go:
```bash
sudo snap install snapcraft go
```

## Usage
Login once:
```
snapcraft login
```

Create the `snapcraft.yaml` file:
```bash
go run create-snapcraft.go --name=edgex-device-gpio
```

Set `--default-icon` to include the default icon in case it needs to be updated as part of the metadata upload.

Build snap and upload metadata:
```bash
./upload-metadata.sh edgex-device-gpio
```
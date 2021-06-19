# SCA Wishlist Notifier

## Introduction

SCA Wishlist notifier parses your Super Cheap Auto (SCA) Wishlist and sends notifications with pushover when there is a product on sale via pushover.

## Getting Started

The application is wrapped using docker, using environment variables.

### docker

```
docker create \
  --name=sca-wishlist-notifier \
  -e WL_SCA_WLID=### \
  -e WL_PUSHOVER_APP=### \
  -e WL_PUSHOVER_RECIPIENT=### \
  -e WL_CRON="0 9 * * *" \
  --restart unless-stopped \
  atajsic/sca-wishlist-notifier
```

### docker-compose

```
version: "3"
services:
  sca-wishlist-notifier:
    image: atajsic/sca-wishlist-notifier
    environment:
      - WL_SCA_WLID=###
      - WL_PUSHOVER_APP=###
      - WL_PUSHOVER_RECIPIENT=###
      - WL_CRON="0 9 * * *"
    restart: unless-stopped
```

## License

This software is distributed under the MIT License.  See the LICENSE file for more details.
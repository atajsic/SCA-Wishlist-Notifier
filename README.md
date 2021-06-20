# SCA Wishlist Notifier

## Introduction

SCA Wishlist notifier parses your Super Cheap Auto (SCA) Wishlist and sends notifications with pushover when there is a product on sale.

## Getting Started

The application is wrapped using docker, using environment variables.

### Supercheap Wishlist ID

You can find this under you account, wishlist, share link. In the URL WishlistID=###

### Cron

See documentation here: http://godoc.org/github.com/robfig/cron

### docker-compose

```
version: "3"
services:
  sca-wishlist-notifier:
    image: atajsic/sca-wishlist-notifier
    environment:
      WL_SCA_WLID: ###
      WL_PUSHOVER_APP: ###
      WL_PUSHOVER_RECIPIENT: ###
      WL_CRON: "0 0 9 * * *"
    volumes:
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
    restart: unless-stopped
```

## License

This software is distributed under the MIT License.  See the LICENSE file for more details.
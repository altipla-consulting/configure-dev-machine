#!/bin/bash

set -eu

sudo apt install -y jq

curl -L -o /tmp/configure-dev-machine $(curl --silent 'https://api.github.com/repos/altipla-consulting/configure-dev-machine/releases/latest' | jq -r '.assets[0].browser_download_url')
chmod +x /tmp/configure-dev-machine

/tmp/configure-dev-machine install

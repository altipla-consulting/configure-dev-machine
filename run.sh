#!/bin/bash

set -eu

sudo apt install -y jq

curl -L -o ~/bin/configure-dev-machine $(curl --silent 'https://api.github.com/repos/altipla-consulting/reloader/releases/latest' | jq -r '.assets[0].browser_download_url')
chmod +x ~/bin/configure-dev-machine

~/bin/configure-dev-machine


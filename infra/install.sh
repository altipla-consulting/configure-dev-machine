#!/bin/bash

set -eu

mkdir -p ~/bin
source ~/.bashrc

curl https://tools.altipla.consulting/bin/configure-dev-machine > ~/bin/configure-dev-machine
chmod +x ~/bin/configure-dev-machine

configure-dev-machine install

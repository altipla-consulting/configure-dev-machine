#!/bin/bash

set -eu

. /opt/ci-toolset/functions.sh

GOOGLE_PROJECT=altipla-tools

configure-google-cloud

mkdir -p tmp
run "go build -o tmp/configure-dev-machine -ldflags=\"-X 'main.Version=$(build-tag)'\" ./cmd/configure-dev-machine"
run "gsutil -h 'Cache-Control: no-cache' cp tmp/configure-dev-machine gs://tools.altipla.consulting/bin/configure-dev-machine"

run "echo $(build-tag) > tmp/version"
run "gsutil -h 'Cache-Control: no-cache' cp tmp/version gs://tools.altipla.consulting/version-manifest/configure-dev-machine"

run "gsutil -h 'Cache-Control: no-cache' cp infra/install.sh gs://tools.altipla.consulting/install/configure-dev-machine"

git-tag

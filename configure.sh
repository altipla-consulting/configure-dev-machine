#!/bin/bash

set -eu


# ------------------------------------------------------------------------------
function install_apt {
  sudo apt update
  sudo apt install -y wget tar curl
}

echo ">>> update apt..."
install_apt


# ------------------------------------------------------------------------------
function install_go {
  sudo rm -rf /usr/local/go
  wget -q -O /tmp/go.tar.gz "https://dl.google.com/go/$1.linux-amd64.tar.gz"
  sudo tar -C /usr/local -xzf /tmp/go.tar.gz
  rm /tmp/go.tar.gz
  go version

  if [ -z "${GOROOT-}" ]; then
    echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
    echo 'export PATH=$PATH:$GOROOT/bin' >> ~/.bashrc
  fi
}

WANTED_VERSION=go1.13.8
if ! hash go 2>/dev/null; then
  echo ">>> install go..."
  install_go $WANTED_VERSION
fi
CURRENT_VERSION=$(go version | { read _ _ v _; echo $v; })
if [ "$CURRENT_VERSION" != "$WANTED_VERSION" ]; then
  echo ">>> update go [$CURRENT_VERSION -> $WANTED_VERSION]..."
  install_go $WANTED_VERSION
fi


# ------------------------------------------------------------------------------
function install_docker_compose {
  sudo curl -L "https://github.com/docker/compose/releases/download/$1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
  docker-compose --version
}

WANTED_VERSION=1.25.4
if ! hash docker-compose 2>/dev/null; then
  echo ">>> install docker-compose..."
  install_docker_compose $WANTED_VERSION
fi
CURRENT_VERSION=$(docker-compose version --short)
if [ "$CURRENT_VERSION" != "$WANTED_VERSION" ]; then
  echo ">>> update docker-compose [$CURRENT_VERSION -> $WANTED_VERSION]..."
  install_docker_compose $WANTED_VERSION
fi
if [ -z "${USR_ID-}" ]; then
  echo ">>> install dc alias..."
  echo 'export USR_ID=$(id -u)' >> ~/.bashrc
  echo 'export GRP_ID=$(id -g)' >> ~/.bashrc
  echo "alias dc='docker-compose'" >> ~/.bashrc
fi
if ! type dcrun 2>/dev/null; then
  echo ">>> install dcrun alias..."
  echo "alias dcrun='docker-compose run --rm'" >> ~/.bashrc
fi


# ------------------------------------------------------------------------------
function install_reloader {
  go get -v -u github.com/altipla-consulting/reloader
}

echo ">>> update reloader..."
install_reloader


# ------------------------------------------------------------------------------
function install_actools {
  sudo curl https://tools.altipla.consulting/bin/actools -o /usr/bin/actools
  chmod +x /usr/bin/actools

  actools pull
}

echo ">>> update actools..."
install_actools


# ------------------------------------------------------------------------------
echo ">>> done!"

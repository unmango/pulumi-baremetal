#!/bin/bash
set -eu

# https://stackoverflow.com/questions/45125516/possible-values-for-uname-m
function arch() {
	case "$(uname -m)" in
		x86_64) echo 'amd64';;
		arm64|aarch64_be|aarch64|armv8b|armv8l) echo 'arm64';;
		arm) echo 'arm';;
	esac
}

function os() {
	case "$(uname -s)" in
		Linux) echo 'linux';;
		*)
		>&2 echo 'OS=UNSUPPORTED # ...jk'
		# idk what values `uname -s` can have but the provisioner doesn't support windows
		echo 'darwin';;
	esac
}

function assertTool() {
	if ! which "$1" 2>/dev/null; then
		>&2 echo "'$1' not found"
		return 1
	fi
}

DEV_MODE="${DEV_MODE:-false}"
GIT="$(assertTool 'git')"
JQ="$(assertTool 'jq')"

# Supporting PULUMI_COMMAND_* prefixes allows for safer sshd AcceptEnv= configurations.
INSTALL_DIR="${INSTALL_DIR:-${PULUMI_COMMAND_INSTALL_DIR:-}}"
CONFIG_DIR="${CONFIG_DIR:-${PULUMI_COMMAND_CONFG_DIR:-}}"
LISTEN_ADDRESS="${LISTEN_ADDRESS:-${PULUMI_COMMAND_LISTEN_ADDRESS:-}}"
LISTEN_NETWORK="${LISTEN_NETWORK:-${PULUMI_COMMAND_LISTEN_NETWORK:-}}"
SYSTEMD_SERVICE_FILE="${SYSTEMD_SERVICE_FILE:-${PULUMI_COMMAND_SYSTEMD_SERVICE_FILE:-}}"
WORK="${WORK:-${PULUMI_COMMAND_WORK:-}}"
VERSION="${VERSION:-${PULUMI_COMMAND_VERSION:-}}"
SRC_BIN="${SRC_BIN:-${PULUMI_COMMAND_SRC_BIN:-}}"
BIN_NAME="${BIN_NAME:-${PULUMI_COMMAND_BIN_NAME:-}}"

if GIT_ROOT=$($GIT rev-parse --show-toplevel 2>/dev/null); then
	IS_GIT=true

	if $DEV_MODE; then
		echo 'Using git repository root'
		WORK="$GIT_ROOT/hack/.work"
		mkdir -p "$WORK"
		INSTALL_DIR="${INSTALL_DIR:-$WORK}"
		CONFIG_DIR="${CONFIG_DIR:-$WORK}"
		LISTEN_ADDRESS='localhost:6969'
		SYSTEMD_SERVICE_FILE="${SYSTEMD_SERVICE_FILE:-$WORK/baremetal.service}"
	else
		echo 'Skipping when in a git repository'
		exit 0
	fi
else
	IS_GIT=false
fi

if [ -n "${CI:-}" ]; then
	echo "🙋 Hello from CI! Do the emoji's make you wanna puke? 🤮"
	if [ -n "${GITHUB_ACTIONS:-}" ]; then
		WORK="$GITHUB_WORKSPACE"
	else
		echo "💀 I'm not sure how to handle this situation"
		exit 1
	fi
fi

WORK="${WORK:-"$(mktemp --tmpdir -d pulumi-baremetal-XXXX)"}"
LOGS="$WORK/logs.txt"
OS="$(os)"
ARCH="$(arch)"
LATEST_RELEASE_JSON="$WORK/release.json"
GITHUB='https://github.com/unmango/pulumi-baremetal'
GITHUB_API='https://api.github.com/repos/unmango/pulumi-baremetal'
VERSION="${VERSION:-}"
SRC_BIN="${SRC_BIN:-provisioner}"
BIN_NAME="${BIN_NAME:-$SRC_BIN}"
LISTEN_ADDRESS="${LISTEN_ADDRESS:-0.0.0.0}"
LISTEN_NETWORK="${LISTEN_NETWORK:-tcp}"
ARCHIVE_TMPL="pulumi-resource-baremetal-\$VERSION-$OS-$ARCH.tar.gz"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
SYSTEMD_DIR='/etc/systemd/system'
SYSTEMD_SERVICE_FILE="${SYSTEMD_SERVICE_FILE:-$SYSTEMD_DIR/baremetal-provisioner.service}"
CONFIG_DIR="${CONFIG_DIR:-$SYSTEMD_SERVICE_FILE.d}"
VARS_FILE="$CONFIG_DIR/vars.env"

function log() {
	tee -a "$LOGS" <<<"$1"
}

function latestRelease() {
	if [ ! -s "$LATEST_RELEASE_JSON" ]; then
		curl "$GITHUB_API/releases/latest" >"$LATEST_RELEASE_JSON"
	fi

	cat "$LATEST_RELEASE_JSON"
}

function printVars() {
	echo '🔥 Running with vars:'
	log "VERSION=$VERSION"
	log "OS=$OS, ARCH=$ARCH"
	log "INSTALL_DIR=$INSTALL_DIR"
	log "LISTEN_ADDRESS=$LISTEN_ADDRESS"
	log "LISTEN_NETWORK=$LISTEN_NETWORK"
	log "WORK=$WORK"
	log "BIN_NAME=$BIN_NAME"
	log "SYSTEMD_SERVICE_FILE=$SYSTEMD_SERVICE_FILE"
	log "SRC_BIN=$SRC_BIN"
	log "GITHUB=$GITHUB"
	log "GITHUB_API=$GITHUB_API"
	log "VARS_FILE=$VARS_FILE"
}

function installProvisioner() {
	BIN="$INSTALL_DIR/$BIN_NAME"
	if [ -f "$BIN" ] && [ "$($BIN --version 2>/dev/null)" == "$VERSION" ]; then
		log "$BIN"
		echo "👻 Up to date"
		return 0
	fi

	ARCHIVE_NAME="$(VERSION="$VERSION" envsubst <<<"$ARCHIVE_TMPL")"
	log "ARCHIVE_NAME=$ARCHIVE_NAME"

	URL="$GITHUB/releases/download/$VERSION/$ARCHIVE_NAME"
	log "URL=$URL"
	echo '🧬 Downloading...'
	curl -L "$URL" | tar -zx -C "$INSTALL_DIR" "$SRC_BIN" --transform "s/$SRC_BIN/$BIN_NAME/"

	chmod +x "$BIN"
	echo 'Provisioner downloaded'
	log "$BIN"
}

function main() {
	if $DEV_MODE; then
		echo '🚧 DEV_MODE=true, Carry on friend'
	elif theyReadTheDocs; then
		echo '🤗 Thanks for reading scripts before you execute!'
	fi

	if [ -z "$VERSION" ]; then
		log 'VERSION unset, fetching latest release'
		VERSION="$(latestRelease | $JQ -r '.tag_name')"
	fi

	printVars

	log 'Ensuring install directory'
	mkdir -p "$INSTALL_DIR"

	installProvisioner

	export PROVISIONER_BIN="$BIN"
	export LISTEN_ADDRESS
	export LISTEN_NETWORK

	if $DEV_MODE && $IS_GIT; then
		FILE="$GIT_ROOT/provider/cmd/provisioner/baremetal-provisioner.service"
		echo "🌱 Sourcing local unit file: $FILE"
		envsubst <"$FILE" >"$SYSTEMD_SERVICE_FILE"
	else
		URL="$GITHUB/releases/download/$VERSION/baremetal-provisioner.service"
		log 'Fetching systemd unit file'
		curl -fL "$URL" | envsubst >"$SYSTEMD_SERVICE_FILE"
	fi

	if $DEV_MODE; then
		echo '🚧 Skipping daemon-reload and service enable in dev mode'
		exit 0
	fi

	log 'Reloading systemd'
	systemctl daemon-reload

	log 'Enabling service'
	systemctl enable --now baremetal-provisioner

	echo '⭐ Done'
}

function theyReadTheDocs() {
	if [ -z "${IREADTHEDOCS:-${PULUMI_COMMAND_IREADTHEDOCS:-}}" ]; then
		echo '🚨 Please take a look through the script here before proceeding'
		echo 'https://github.com/unmango/pulumi-baremetal/tree/main/provider/cmd/provisioner/install.sh'
		echo ''
		echo 'This script would like to create or modify these files:'
		echo "$INSTALL_DIR/$BIN_NAME"
		echo "$SYSTEMD_SERVICE_FILE"
		echo "$VARS_FILE"
		echo ''
		echo 'And make requests to these urls:'
		echo "$GITHUB"
		echo "$GITHUB_API"
		echo ''
		echo 'If this is expected, update your env and run this script again'
		exit 1
	fi
}

main

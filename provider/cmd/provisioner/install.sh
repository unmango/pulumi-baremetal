#!/bin/bash
set -eu

function require() {
	if [ -z "$1" ]; then
		echo "$2 is required"
		exit 1
	fi
}

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

if GIT_ROOT=$($GIT rev-parse --show-toplevel 2>/dev/null); then
	IS_GIT=true
else
	IS_GIT=false
fi

if $IS_GIT; then
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

require "${LISTEN_ADDRESS:-}" 'LISTEN_ADDRESS'
WORK="${WORK:-"$(mktemp --tmpdir -d pulumi-baremetal-XXXX)"}"
LOGS="$WORK/logs.txt"

function log() {
	tee -a "$LOGS" <<<"$1"
}

OS="$(os)"
ARCH="$(arch)"
LATEST_RELEASE_JSON="$WORK/release.json"
GITHUB='https://github.com/unmango/pulumi-baremetal'
GITHUB_API='https://api.github.com/repos/unmango/pulumi-baremetal'
VERSION="${VERSION:-}"
SRC_BIN="${SRC_BIN:-provisioner}"
BIN_NAME="${BIN_NAME:-$SRC_BIN}"
ARCHIVE_TMPL="pulumi-resource-baremetal-\$VERSION-$OS-$ARCH.tar.gz"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
SYSTEMD_DIR='/etc/systemd/system'
SYSTEMD_SERVICE_FILE="${SYSTEMD_SERVICE_FILE:-$SYSTEMD_DIR/baremetal-provisioner.service}"
CONFIG_DIR="${CONFIG_DIR:-$SYSTEMD_SERVICE_FILE.d}"
VARS_FILE="$CONFIG_DIR/vars.env"

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
	log "WORK=$WORK"
	log "BIN_NAME=$BIN_NAME"
	log "SYSTEMD_SERVICE_FILE=$SYSTEMD_SERVICE_FILE"
	log "SRC_BIN=$SRC_BIN"
	log "GITHUB=$GITHUB"
	log "GITHUB_API=$GITHUB_API"
	log "VARS_FILE=$VARS_FILE"
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

	BIN="$INSTALL_DIR/$BIN_NAME"
	if [ -f "$BIN" ]; then
		log "$BIN"
		echo "👻 Up to date"
		exit 0
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

	if [ ! -f "$SYSTEMD_SERVICE_FILE" ]; then
		if $DEV_MODE && $IS_GIT; then
			FILE="$GIT_ROOT/provider/cmd/provisioner/baremetal-provisioner.service"
			sed "s@\$VARS_FILE@$VARS_FILE@g" "$FILE" >"$SYSTEMD_SERVICE_FILE"
		else
			URL="$GITHUB/releases/download/$VERSION/baremetal-provisioner.service"
			log 'Fetching systemd unit file'
			curl -fL "$URL" | sed "s@\$VARS_FILE@$VARS_FILE@g" >"$SYSTEMD_SERVICE_FILE"
		fi
	fi

	log 'Ensuring config directory'
	mkdir -p "$CONFIG_DIR"

	log 'Creating vars file'
	TMP_VARS="$(mktemp --tmpdir XXXX.env)"
	echo "PROVISIONER_BIN=$BIN" | tee -a "$TMP_VARS"
	echo "LISTEN_ADDRESS=$LISTEN_ADDRESS" | tee -a "$TMP_VARS"
	mv "$TMP_VARS" "$VARS_FILE"

	log 'Reloading systemd'
	systemctl daemon-reload

	log 'Enabling service'
	systemctl enable --now baremetal-provisioner

	echo '⭐ Done'
}

function theyReadTheDocs() {
	if [ -z "${IREADTHEDOCS:-}" ]; then
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

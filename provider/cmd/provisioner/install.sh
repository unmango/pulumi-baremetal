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
	else
		echo 'Skipping when in a git repository'
		exit 0
	fi
fi

if [ -n "${CI:-}" ]; then
	echo "🙋 Hello from CI! Do the emoji's make you puke? 🤮"
	if [ -n "${GITHUB_ACTIONS:-}" ]; then
		WORK="$GITHUB_WORKSPACE"
	else
		echo "💀 I'm not sure how to handle this situation"
		exit 1
	fi
fi

WORK="${WORK:-"$(mktemp --tmpdir -d pulumi-baremetal-XXXX)"}"
LOGS="$WORK/logs.txt"

function log() {
	tee -a "$LOGS" <<<"$1"
}

OS="$(os)"
ARCH="$(arch)"
LATEST_RELEASE_JSON="$WORK/release.json"
BASE_URL='https://api.github.com/repos/unmango/pulumi-baremetal'
VERSION="${VERSION:-}"
SRC_BIN="${SRC_BIN:-provisioner}"
BIN_NAME="${BIN_NAME:-$SRC_BIN}"
ARCHIVE_TMPL="pulumi-resource-baremetal-\$VERSION-$OS-$ARCH.tar.gz"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

function latestRelease() {
	if [ ! -s "$LATEST_RELEASE_JSON" ]; then
		curl "https://api.github.com/repos/unmango/pulumi-baremetal/releases/latest" >"$LATEST_RELEASE_JSON"
	fi

	cat "$LATEST_RELEASE_JSON"
}

function main() {
	if [ -z "$VERSION" ]; then
		log 'VERSION unset, fetching latest release'
		VERSION="$(latestRelease | $JQ -r '.tag_name')"
	fi

	echo '🔥 Running with vars:'
	log "VERSION=$VERSION"
	log "OS=$OS, ARCH=$ARCH"
	log "INSTALL_DIR=$INSTALL_DIR"
	log "WORK=$WORK"
	log "BIN_NAME=$BIN_NAME"
	log "SRC_BIN=$SRC_BIN"
	log "BASE_URL=$BASE_URL"

	log 'Testing install directory'
	if ! mkdir -p "$INSTALL_DIR"; then
		log "Failed to mkdir -p $INSTALL_DIR"
		exit 1
	fi

	BIN="$INSTALL_DIR/$BIN_NAME"
	if [ -f "$BIN" ]; then
		log "$BIN"
		echo "👻 Up to date"
		exit 0
	fi

	ARCHIVE_NAME="$(VERSION="$VERSION" envsubst <<<"$ARCHIVE_TMPL")"
	log "ARCHIVE_NAME=$ARCHIVE_NAME"

	URL="https://github.com/unmango/pulumi-baremetal/releases/download/$VERSION/$ARCHIVE_NAME"
	log "URL=$URL"
	curl -L "$URL" | tar -zx -C "$INSTALL_DIR" "$SRC_BIN" --transform "s/$SRC_BIN/$BIN_NAME/"

	chmod +x "$BIN"
	echo '⭐ Done'
	log "$BIN"
}

main

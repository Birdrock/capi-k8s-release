#!/usr/bin/env bash

set -ex

CF_FOR_K8s_DIR="${CF_FOR_K8s_DIR:-${HOME}/workspace/cf-for-k8s/}"
SCRIPT_DIR=$(dirname $0)
REPO_BASE_DIR="${SCRIPT_DIR}/.."

if [ -z "$1" ]
  then
    echo "No values file was supplied!"
		exit 1
fi

if [ -z "$2" ]
  then
    echo "No image destination was supplied!"
		exit 1
fi

KBLD_TMP="$(mktemp)"
ytt -f "${REPO_BASE_DIR}/dev-templates/" -v kbld.destination="${2}" > "${KBLD_TMP}"

VALUES_TMP="$(mktemp)"
echo "#@data/values" > "${VALUES_TMP}"

kbld -f "${KBLD_TMP}" -f $1 > "${VALUES_TMP}"

cat "${VALUES_TMP}" > "$1"
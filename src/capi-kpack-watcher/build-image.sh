#! /bin/sh

BASEDIR=$(dirname $0)

case $# in
	0) docker build --build-arg GIT_REV=$(git rev-parse HEAD) -t awittrock/test-capi-kpack-watcher $BASEDIR  ;;
	2) docker build --build-arg BUILD_IMAGE=$1 --build-arg RUN_IMAGE=$2 --build-arg GIT_REV=$(git rev-parse HEAD) -t awittrock/test-capi-kpack-watcher $BASEDIR ;;
	*) echo "Usage: build-image.sh [BUILD_IMAGE RUN_IMAGE]" ;;

esac

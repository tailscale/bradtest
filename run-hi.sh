#!/usr/bin/env bash

set -eu
(cd $GITHUB_ACTION_PATH && go build)

exec $GITHUB_ACTION_PATH/bradtest "$@"

#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

source ./hack/lib/test_lib.sh
source ./hack/lib/image_lib.sh

test_dir=./test
tests=$test_dir/e2e-go

export TRACE=1
export GO111MODULE=on

###########################################################################
### DO NOT UNCOMMENT THESE LINES UNLESS YOU KNOW WHAT YOU'RE DOING !!!! ###
###                                                                     ###
### They cause the integration image not to be loaded into kind in      ###
### TravisCI.                                                           ###
###                                                                     ###
###########################################################################
###
###    #prepare_staging_dir $tmp_sdk_root
###    #fetch_envtest_tools $tmp_sdk_root
###
###########################################################################
setup_envs $tmp_sdk_root

go test $tests -v -ginkgo.v

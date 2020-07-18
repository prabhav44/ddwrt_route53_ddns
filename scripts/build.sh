#!/bin/bash

TOOLCHAIN='/usr/local/toolchain-arm_cortex-a9_gcc-8.2.0_musl_eabi'

STRIP="$TOOLCHAIN/bin/arm-openwrt-linux-muslgnueabi-strip"

GOOS=linux GOARCH=arm GOARM=5 go build -o bin/route53_ddns cmd/main.go
$STRIP bin/route53_ddns
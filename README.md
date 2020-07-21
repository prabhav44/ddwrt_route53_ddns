# Route 53 Dynamic DNS Updater

This program is meant to be run on Armv7 based devices, I am currently developing this for the Netgear R8500 and RaspberryPi 3B. The makefile builds the binary and uploads it via SCP to the target device.

The DDNS API it is targeting was generated via this repo: https://github.com/awslabs/route53-dynamic-dns-with-lambda and the associated cloudformation template.

Currently I am only implementing the set IP and get IP functions of the API. I plan on adding IPv6 eventually. 

## Build:
~~~bash
TOOLCHAIN='/usr/local/toolchain-arm_cortex-a9_gcc-8.2.0_musl_eabi'

STRIP="$TOOLCHAIN/bin/arm-openwrt-linux-muslgnueabi-strip"
~~~
I'm using the ARM Cortex toolchain for DD-WRT routers found [here](https://download1.dd-wrt.com/dd-wrtv2/downloads/toolchains/toolchains.tar.xz) to strip the binary and reduce size. You could probably use the default `strip` command in your OS.

Feel free to modify the Makefile for local target testing, I currently have it run `scp` to copy the binary to my rpi-3b.

## CLI Commands:
#### Set IPv4 Address of Hostname to current WAN network's IPv4:
~~~bash
[pparashar@rpi-3b golang-projects]$ ./route53_ddns set -hostname=HOSTNAME_TO_BE_UPDATED
                                                       -api-key=YOUR_DDNS_API_KEY
                                                       -shared-secret=SHARED_SECRET_FOR_HOSTNAME
                                                       -api-url=YOUR_DDNS_API_URL
~~~
#### Output:
~~~bash
2020/07/20 22:35:32 success
2020/07/20 22:35:32 Your IP address matches the current Route53 DNS record.
~~~
#### Get public IPv4 Address of current WAN network:
~~~bash
[pparashar@rpi-3b golang-projects]$ ./route53_ddns get -api-key=YOUR_DDNS_API_KEY
                                                       -api-url=YOUR_DDNS_API_URL
~~~
#### Output:
~~~
123.123.123.XXX
~~~
# Route 53 DDNS Updater

This program is meant to be run on Armv7 based devices, I am currently developing this for the Netgear R8500 and RaspberryPi 3B. The makefile builds the binary and uploads it via SCP to the target device.
The DDNS API it is targeting was generated via this repo: https://github.com/awslabs/route53-dynamic-dns-with-lambda and the associated cloudformation template.
TOOLCHAIN=/usr/local/toolchain-arm_cortex-a9_gcc-8.2.0_musl_eabi

STRIP=$(TOOLCHAIN)/bin/arm-openwrt-linux-muslgnueabi-strip

ddwrt_route53_ddns: ddwrt_route53_ddns.go Makefile
	GOOS=linux GOARCH=arm GOARM=5 go build -o ddwrt_route53_ddns ddwrt_route53_ddns.go
	$(STRIP) ddwrt_route53_ddns
	scp ddwrt_route53_ddns root@obi-wan:/jffs

clean:
	rm -f ddwrt_route53_ddns

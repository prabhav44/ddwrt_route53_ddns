route53_ddns: 
	./scripts/build.sh
	scp bin/route53_ddns pparashar@rpi-3b:/home/pparashar/golang-projects/

clean:
	rm -f bin/route53_ddns

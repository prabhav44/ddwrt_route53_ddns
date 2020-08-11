route53_ddns: 
	./scripts/build.sh
	scp bin/route53_ddns pi@rpi-3b:/home/pi/golang-projects/

clean:
	rm -f bin/route53_ddns

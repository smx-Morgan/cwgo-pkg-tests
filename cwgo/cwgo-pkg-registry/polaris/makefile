prepare:
	docker pull polarismesh/polaris-standalone:v1.17.2
	docker run -d --privileged=true \
	--name polaris-standalone \
	-p 15010:15010 \
	-p 8101:8101 \
	-p 8100:8100 \
	-p 8080:8080 \
	-p 8090:8090 \
	-p 8091:8091 \
	-p 8093:8093 \
	-p 8761:8761 \
	-p 8848:8848 \
	-p 9848:9848 \
	-p 9090:9090 \
	-p 9091:9091 polarismesh/polaris-standalone:latest



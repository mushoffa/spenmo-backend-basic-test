all:
	docker-compose -f docker-compose-db.yml up -d ;\
	make -C user-service/ all ;\
	make -C card-service/ all ;\
	make -C wallet-service/ all ;\
	make -C transaction-service/ all ;\
	make -C api-gateway/ all ;\

clean:
	make -C user-service/ clean ;\
	make -C card-service/ clean ;\
	make -C wallet-service/ clean ;\
	make -C transaction-service/ clean ;\
	make -C api-gateway/ clean ;\

docker:
	DOCKER_BUILDKIT=0 docker-compose up -d --build
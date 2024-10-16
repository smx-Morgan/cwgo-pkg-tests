NODE1=172.18.31.33
REGISTRY=bitnami/etcd
ETCD_VERSION=latest

# run single node in docker 
prepare:
	docker network create app-tier --driver bridge
	docker pull ${REGISTRY}:${ETCD_VERSION}
	docker run -d --name Etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest

# run cluster
prepare-cluster:
	docker-compose up -d

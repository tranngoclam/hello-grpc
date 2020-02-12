build-server:
	DOCKER_BUILDKIT=1 docker build -f helloserver/Dockerfile -t helloserver ./helloserver

build-client:
	DOCKER_BUILDKIT=1 docker build -f helloclient/Dockerfile -t helloclient ./helloclient

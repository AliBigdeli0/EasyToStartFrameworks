BIN_PATH=./bin
CLIENT_OUT=client
SERVER_OUT=server
FLAGS = -mod=vendor
ARGS= --config=config.yml --base=$(PWD)/bin/

GO=go
GDB=gdb

DOCKER_COMPILE=go_1_8_compile
DOCKER_PYTHON=geo_active_python
DOCKER_GRPC=geo_active_grpc

DOCKER_KAFKA_UI_NAME=obsidiandynamics/kafdrop
DOCKER_REDIS_UI=rediscommander/redis-commander:latest
DOCKER_DOCKER_UI=portainer/portainer-ce
DOCKER_FILE_ARGS_DATABASE=-f build/docker/docker-compose.database.yml
DOCKER_FILE_ARGS= $(DOCKER_FILE_ARGS_DATABASE) -f build/docker/kafka-docker/docker-compose.yml -f build/docker/docker-compose.req.yml 

################################
## client

client: perbuild build_client
	@$(BIN_PATH)/$(CLIENT_OUT).out $(ARGS)

dc: perbuild build_client
	@$(GDB) $(BIN_PATH)/$(CLIENT_OUT).out

build_client: 
	@$(GO) build $(FLAGS) -o $(BIN_PATH)/$(CLIENT_OUT).out ./cmd/$(CLIENT_OUT)/startup.go

################################
## server
server: perbuild build_server
	@$(BIN_PATH)/$(SERVER_OUT).out  $(ARGS)

ds: perbuild build_server
	@$(GDB) $(BIN_PATH)/$(SERVER_OUT).out

build_server:
	@$(GO) build $(FLAGS) -o $(BIN_PATH)/$(SERVER_OUT).out ./cmd/$(SERVER_OUT)/startup.go



################################
## common
perbuild: 
	@mkdir -p $(BIN_PATH)
	@cp ./configs/config.yml $(BIN_PATH)
	@cp ./configs/log_config.json $(BIN_PATH)
 

################################
## test
test: 
	@$(GO) test ./test/...

.PHONY: client server perbuild test dc ds

################################
## golang compiler 
docker_build:
	@docker build -f ./build/docker/Dockerfile_compile -t $(DOCKER_COMPILE) .

docker_run:
	@docker run -it --rm --pid=host -p 8095:80 -p 8096:81 -p 8097:82 --network=docker_active_geo  --privileged -v $(PWD):/app -u root --name $(DOCKER_COMPILE)_client $(DOCKER_COMPILE) 

docker_run_server:
	@docker run -it --rm --pid=host -p 8098:80 -p 8099:81 -p 8100:82 --network=docker_active_geo  --privileged -v $(PWD):/app -u root --name $(DOCKER_COMPILE)_server $(DOCKER_COMPILE) 

docker_rm:
	@docker image rm $(DOCKER_COMPILE) -f

################################
## docker python
docker_python_build:
	@docker build -f ./build/docker/python/Dockerfile_python -t $(DOCKER_PYTHON) .

docker_python_run:
	@docker run -it --rm --pid=host -p 8096:80 --network=docker_active_geo  --privileged -v $(PWD):/app -u root --name $(DOCKER_PYTHON)_run $(DOCKER_PYTHON) 

docker_python_rm:
	@docker image rm $(DOCKER_PYTHON) -f


################################
## grpc
docker_grpc_build:
	@docker build -f ./build/docker/Dockerfile_grpc -t $(DOCKER_GRPC) .

docker_grpc_run:
	@docker run -it --rm  -v $(PWD):/app -u root --name $(DOCKER_GRPC)_run $(DOCKER_GRPC) 

docker_grpc_rm:
	@docker image rm $(DOCKER_GRPC) -f	

################################
## docker compose
docker_compose_up_all: 
	@docker-compose  $(DOCKER_FILE_ARGS) -f build/docker/kafka-docker/docker-compose.yml up

docker_compose_down_all: 
	@docker-compose $(DOCKER_FILE_ARGS) -f build/docker/kafka-docker/docker-compose.yml down	

docker_compose_rm: 
	@docker image remove docker_go_client
	@docker image remove docker_go_server

docker_compose_up: 
	@docker-compose  $(DOCKER_FILE_ARGS)  up
docker_compose_down: 
	@docker-compose $(DOCKER_FILE_ARGS) down	

################################
## docker tools 

#kafka ui 

docker_kafka_ui_run: 
	@docker run --rm -p 9000:9000 --name=kaka_ui  \
    -e KAFKA_BROKERCONNECT=172.20.200.5:9092 \
    -e JVM_OPTS="-Xms32M -Xmx64M" \
    -e SERVER_SERVLET_CONTEXTPATH="/" \
	--network=docker_active_geo\
    $(DOCKER_KAFKA_UI_NAME)

#.kafka ui

#redis ui

docker_redis_ui_run: 
	@docker run --rm -p 9001:8081 --name redis-gui \
    -e REDIS_HOSTS=local:redis:6379 \
	--network=docker_active_geo\
    $(DOCKER_REDIS_UI)

#.redis ui

#mongo ui 
mongo_ui_run: 
	@docker-compose $(DOCKER_FILE_ARGS_DATABASE) exec mongodb1 mongo "mongodb://mongodb1:27017, mongodb2:27017, mongodb3:27017/?replicaSet=activegeo_set"
#.mongo ui


#docker ui

docker_docker_ui:
	@docker run --rm -p 9002:9000 --name=docker_gui -v /var/run/docker.sock:/var/run/docker.sock -v $(PWD)/third_party/docker_gui:/data $(DOCKER_DOCKER_UI)
#.docker ui
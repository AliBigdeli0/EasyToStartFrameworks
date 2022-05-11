# Golang Template 
There are several futures in this template including : 
* client/server architecture is implemented
    * to see client startup files check [client startup.go](./cmd/client/startup.go) and [internal client startup.go](./internal/app/client/startup.go)
    * to see client startup files check [server startup.go](./cmd/server/startup.go) and [internal server startup.go](./internal/app/server/startup.go)

* All actions can be executed in make file. For eample to run client part of program you can run `make build_client`

* Unit test framework with two sampale file is implemented. for mote information see [this](./test/unit_testing/file_test.go) sample

* External configure is availible to use in Yaml format. to see configure structure see [config.go](./internal/pkg/config.go) file, and to see Yaml file check  [config.yaml](./configs/config.yml)

* Golang compiler is set to docker file. 
    * To see golang compiler docker file check [this link](./build/docker/Dockerfile_compile). As you see in the file all compiler needs is designed pure ubuntu docker file

* Container based kakfa to run with docker compose 

* Container based redis to run with docker compose

* Container based mongodb with 3 replication set

* Container based postgresql database 

* All projects are putted into `Docker Compose` file. As a result, All projects and its requirements, including kafka, redis, replicated mongodb database, client and server could be runned with running `make docker_compose_up_all` .
    * To run requirements you can simply run `make docker_compose_up` in your terminal. 

* GRPC based comminucation with container to generate `GRPC` go files.
    * see proto file in [ping_grpc.proto](./third_party/grpc/ping_grpc.proto)
    * the generated file will put into [./pkg/public_lib/grpc](./pkg/public_lib/grpc) link. to generate simply run `make docker_grpc_run`

* Python interpreter to install and future works
    * to run python docker execute `make docker_python_run` in your terminal

<br/>
<br/>

# List Make Commands

In the following section, available `make` commands will be represented: 
<br/>
<br/>
## Client Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| client | this target is compile and execute client section | N/A |
| dc | this target is execute a debugger for client  | `make client` must be runned before |
| build_client | this target is build all client part and put results into `./bin` path.the file name will be `client.out` | N/A |

<br/>
<br/>

## Server Actions

| Target | Description | Requirements |
| :---: | :---: | :---: |
| server | this target is compile and execute server section | N/A |
| ds | this target is execute a debugger for server  | `make server` must be runned before |
| build_client | this target is build all server part and put results into `./bin` path. The file name will be `server.out` | N/A |

<br/>
<br/>

## PerBuild Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| perbuild | this target is execute before compile in both server and client and copy some important files including `config.yml` and `log_config.json` into the `./bin` path | N/A |

<br/>
<br/>

## Test Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| test | this target is execute all unit tests file in [tests folder](./test)| N/A | 

<br/>
<br/>

## Golang Compiler Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| docker_build | this target is create a docker image with golang v.1.17 and bash script and other debug tools | N/A | 
| docker_run | this target is create a container for client  | N/A | 
| docker_run_server | this target iscreate a container for server | N/A | 
| docker_rm | this target is remove docker image | N/A | 

<br/>
<br/>

## Python Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| docker_python_build | this target is create a docker image for python actions | N/A | 
| docker_python_run | this target is create a container for python actions  | N/A | 
| docker_python_rm | this target is remove the python image | N/A | 

<br/>
<br/>

## GRPC Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| docker_grpc_build | this target is create a docker image for grpc actions | N/A | 
| docker_grpc_run | this target is create a container for compile file and put grpc into the `./pkg/public_lib/grpc/` path  | N/A | 
| docker_grpc_rm | this target is remove the grpc image | N/A | 

<br/>
<br/>

## Docker Compose Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| docker_compose_up_all | this target run all containers into a one docker compose, the containers is including : Mongodb database, kafka, posgresql, redis, client and server  | N/A | 
| docker_compose_down_all | this target shutdown all running containers | N/A | 
| docker_compose_rm | this target remove client and server images | N/A | 

<br/>
<br/>

## Kafka UI Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| docker_kafka_ui_run | this target is run a kafka container | N/A |

<br/>
<br/>


## Redis UI Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| docker_redis_ui_run | this target is run a redis container | N/A |

<br/>
<br/>

## Redis UI Actions
| Target | Description | Requirements |
| :---: | :---: | :---: |
| mongo_ui_run | this target is run a mongodb container | N/A |

<br/>
<br/>

**NOTE**: All containers are isolated environment for development.

<br/>
<br/>

# LICENSE
This project is licensed under the GNU GENERAL PUBLIC LICENSE - see the [LICENSE](../LICENSE) file for details


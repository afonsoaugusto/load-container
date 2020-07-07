# LOAD-CONTAIENR

## Use submodule for pattern make and deploy

`git submodule add git@github.com:afonsoaugusto/base-ci.git`

`git submodule update --init --recursive`

```sh
docker run --rm -it \
    -v $PWD:/src/ \
    -e "GOPATH=/src" \
    golang bash

make build && make scan && docker run  --rm -it -p 9000:9000 afonsoaugusto/load-container:b8e61d5-master

docker build -t api .
docker run --rm -it -p 9000:9000 --name api api

curl -iL localhost:9000/
curl -iL localhost:9000/mem-usage

docker build -t api . && \
docker run --rm -it -p 9000:9000 --name api  --cpus="1" --cpu-shares=2 --cpuset-cpus=1 api


docker run --rm -it -p 9000:9000 --name api -m 10m --memory-swap -1 api

docker run --rm -it -p 9000:9000 --name api  --cpus="1" --cpu-shares=2 --cpuset-cpus=1 afonsoaugusto/load-container:latest
```

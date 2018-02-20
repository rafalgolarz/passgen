[![Build Status](https://travis-ci.org/rafalgolarz/passgen.svg?branch=master)](https://travis-ci.org/rafalgolarz/passgen)
[![CircleCI](https://circleci.com/gh/rafalgolarz/passgen/tree/master.svg?style=svg)](https://circleci.com/gh/rafalgolarz/passgen/tree/master)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/rafalgolarz/passgen)
[![GoDoc](https://godoc.org/github.com/rafalgolarz/passgen?status.svg)](https://godoc.org/github.com/rafalgolarz/passgen)
[![codebeat badge](https://codebeat.co/badges/3cadc60b-3642-46bc-9118-1595e354aa6d)](https://codebeat.co/projects/github-com-rafalgolarz-passgen-master)
[![Docker Automated buil](https://img.shields.io/docker/automated/jrottenberg/ffmpeg.svg)](https://hub.docker.com/r/rafalgolarz/passgen-builds/builds/)

# Passgen

**API to generate secure passwords**

[![asciicast](https://asciinema.org/a/141109.png)](https://asciinema.org/a/141109?speed=2)

## Configurarion

**default.toml** stores minimum default required settings for generated passwords

Currently it contains two sections: 

```toml
[default]
min_length = 8
min_special_characters = 2
min_digits = 2
min_lowercase = 1
min_uppercase = 1
results = 1

[strong]
min_length = 16
min_special_characters = 4
min_digits = 4
min_lowercase = 2
min_uppercase = 2
results = 1
```

You can choose which section should be loaded as default one, in **init.go**:

```toml
passwordType = "default"
```

Each of the params can be overwritten from the url.

## Running

I recommend to use docker:

```!/bin/bash
docker run --rm -p 8080:8080 --name=passgen rafalgolarz/passgen
```

or run the dev version:

```!/bin/bash
docker run --rm -p 8080:8080 --name=passgen rafalgolarz/passgen:dev
```

You can also build a docker image locally:

```!/bin/bash
docker build -f Dockerfile-dev -t rafalgolarz/passgen .
docker run --rm -p 8080:8080 --name=passgen rafalgolarz/passgen
```

Dockerfile-dev has debug flags on. Use **Dockerfile** to have them off or set them with -e param:

Next, open the url:

<http://localhost:8080/v1/passwords>

By default, it generates one password meeting criteria defined in config.toml but you can overwrite any of the params.

Generate 3 passwords:

<http://localhost:8080/v1/passwords?res=3>

Generate 20 passwords. Each of the passwords should have:

- minimum length of 25 characters
- minimum of 2 special characters
- minimum of 2 digits
- minimum 4 lower and 4 upper case letters

<http://localhost:8080/v1/passwords/?min-length=25&min-specials=2&min-digits=2&min-lowers=4&min-uppers=4&res=20>

## Running tests

```!/bin/bash
cd $GOPATH/src/github.com/rafalgolarz/passgen
go test -v ./... -bench=./...
```

If you don't know/have $GOPATH, run **go env** to display all Go enviromental variables

## Deploying to Kubernetes

In **k8s** folder, you can find sample kubernetes manifests that can help you to deploy it to your K8s cluster.

I recommend to use different namespaces for dev and production versions.

To create two new namespaces (**dev** and **production**) you can use my yaml files:

```!/bin/bash
kubectl create -f k8s/production/production-ns.yaml
kubectl create -f k8s/dev/dev-ns.yaml
```

List available namespaces:

```!/bin/bash
kubectl get ns
```

Deploy passgen API to **dev** namespace:

```!/bin/bash
kubectl -n dev apply -f k8s/dev/api-passgen-deployment.yaml
kubectl -n dev apply -f k8s/api-passgen-svc.yaml
```

Deploy passgen API to **production** namespace:

```!/bin/bash
kubectl -n production apply -f k8s/production/api-passgen-deployment.yaml
kubectl -n production apply -f k8s/api-passgen-svc.yaml
```

Once everything is up and running, you can access the API through kubectl proxy:

```!/bin/bash
kubectl proxy
```

If you deployed it to the dev namespace, open the url for dev:
<http://localhost:8001/api/v1/proxy/namespaces/dev/services/api-passgen:80/v1/passwords/>

If you deployed it to the production namespace, open the url for production:
<http://localhost:8001/api/v1/proxy/namespaces/production/services/api-passgen:80/v1/passwords/>

## TODO

- [ ] allow passing the path to config.toml from the command line
- [ ] display just array of passwords (hide configuration info) by setting a new boolean param verbose to false
- [ ] add tests checking if the length of required subsets of characters match passed values
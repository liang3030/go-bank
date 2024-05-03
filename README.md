# Go Bank

Go Bank is a tutorial project for practice golang.

## Getting Started
You could set up it locally by following the steps. We assume you have golang and docker are installed in your local enviroment.

### Prerequisites

1. Setting database with docker. 

```shell

docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres

```

2. Install dependencies.

In `Makefile` , we setup command to install depencies. In the root repository of the project, run following command in your terminal.

```shell

make install

```

3. Export jwt secret

```shell

export JWT_SECRET="soent435werh"

```

## Usage

1. build

```shell
make build
```

2. seed the database. First time run the project, it needs to seed the database. 

```shell
./bin/go-bank --seed

```

2. run project

```shell
make run
```

## Others 

1. Thunder Client, it is a nice plugin of vscode to test http request.

## Reference


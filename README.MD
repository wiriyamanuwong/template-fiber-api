# Goland Template API Service 
with Fiber V2 Framework 

> The following is a folder structure pattern that I use 
> [Golang Clean Architecture - https://github.com/restuwahyu13/go-clean-architecture](https://github.com/restuwahyu13/go-clean-architecture) 
> 

## Table of Contents
- [Goland Template API Service](#goland-template-api-service)
  - [Table of Contents](#table-of-contents)
  - [Root Folder Structure](#root-folder-structure)
  - [Go Import Flow](#go-import-flow)
  - [Prerequisites](#prerequisites)
  - [How to setup](#how-to-setup)
  - [How to Run](#how-to-run)


## Root Folder Structure

<img src="./images/folder-structure.png" width="800px" alt="Folder Structure" />

## Go Import Flow
<img src="./images/go-import-flow.png" width="800px" alt="Folder Structure" />


## Prerequisites

- [Go version 1.21 or above](https://golang.org/)
- [taskfile.dev](https://taskfile.dev/) - same `makefile` but config with `yaml` file
- [Cobra-cli](https://github.com/spf13/cobra-cli)  - create cli with project
- [Swag Go](https://github.com/swaggo/swag) - Swag converts Go annotations to Swagger Documentation 2.0. 



## How to setup

1. Init Go Module
    ```bash
    go mod init github.com/xxx/xxx
    ```

2. Replace `default.env` to [`.env`](./.env)
    ```bash
    mv default.env .env
    ```
3. Change go import `github.com/attapon-th/template-fiber-api` to your go module
    
    <img src="./images/chage-go-import.png" width="600px" alt="Folder Structure" />

4. Download Go Module
  
    with `taskfile`
    ```bash
    task mod
    ```
    or with `go mod`
    ```bash
    go mod download
    go mod tidy
    go mod vendor
    ```

5. Rename [`default.env`](./default.env) to [`.env`](./.env)
    ```bash
    mv default.env .env
    ```

## How to Run

Run `dev` for start server dev mode:

```bash
task dev
```

Run `swag` for generate swagger docs:

```bash
task swag
```

Run `build` for build go binary file

(output directory: `./build/`)
```bash
task build
```

Run `docker-build` fors Build Docker Image:

```bash
task docker-build
```
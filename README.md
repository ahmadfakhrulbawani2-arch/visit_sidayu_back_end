# Go, Gin, Gorm, Postgre REST-API Template Made Ez

<div align="center">
  <img src="https://img.shields.io/badge/Go-%23363636.svg?style=for-the-badge&logo=go&logoColor=white&color=00ADD8">
  <img src="https://img.shields.io/badge/JWT-%23F8CC46.svg?style=for-the-badge&logo=jsonwebtokens&logoColor=white&color=000000">
  <img src="https://img.shields.io/badge/PostgreSQL-%23007ACC.svg?style=for-the-badge&logo=postgresql&logoColor=white&color=4169E1">
  <img src="https://img.shields.io/github/repo-size/ahmadfakhrulbawani2-arch/simple_go_gin_gorm_postgres_be_template?style=for-the-badge&color=orange">
  <img src="https://img.shields.io/github/stars/ahmadfakhrulbawani2-arch/simple_go_gin_gorm_postgres_be_template?style=for-the-badge&color=yellow">
  <img src="https://img.shields.io/github/forks/ahmadfakhrulbawani2-arch/simple_go_gin_gorm_postgres_be_template?style=for-the-badge&color=purple">
</div>

## ℹ️ Overview

This monorepo is my basic Go Backend REST API. The design architecture is my simplified version of more advance Go REST API template out there. The feature is complete and passed for MVP (Minimum Viable Product)

Here is the list of this repo feature:<br />
✅ Standard Go Framework using GIN-Gonic and Go ORM (GORM)<br />
✅ Standard Json Web Token (JWT) authentication<br />
✅ Unprotected and protected authorization routes<br />
✅ Standard JSON response (header, data, meta)<br />
✅ Plug and Play Postgre Integration<br />
✅ Plug and Play Imagekit for media Storage<br />

What's missing:<br />
❌ Standard security middlewares (xss protection, cors, etc.) for simplicity and flexibility <br />
❌ Comprehensive file upload to SDK such as Cloudflare R2 and AWS S3<br />
❌ Async API features<br />
❌ Infra and DevOps<br />
❌ Granular migration (will be added for next development, currently only up and down all database tables) <br />
I'll cover these missing feature in upcoming Go async app template.<br />
To simplify this `README.md`, we move brief description inside [design-pattern.md](./design-pattern.md)

## 🚀 Get Started

Before you can developing with this repo, there's some prerequisites to be fulfilled.

### 🛠️ Prerequisites

There's no order of operation below, but please `install .msi software for windows and any extension on Linux/MacOs`. To check if everything work, run the code snippet given below in powershell or terminal.

1. Install git in [here](https://git-scm.com/install/windows). Check with:

```sh
git -v
```

2. Install code editor like VsCode
3. Install Golang in [here](https://go.dev/). Check with:
4. (Optional) Install air for live reload.

```go
go install github.com/air-verse/air@latest
```

```go
go version
```

Below are optional extension for visual studio compatible IDE/code editor but good to have:

1. Prettier extension, download [here](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
2. Go extension, donwload [here](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### ⚙️ Start Your Development

There's order of operation below, make sure you follow all steps sequentially. All error handling are provided at the footer.

1. Make new folder/directory for this template, name it with your own project.
2. Enter to that directory, then open terminal and run:

```bash
git clone https://github.com/ahmadfakhrulbawani2-arch/simple_go_gin_gorm_postgres_be_template.git .
```

3. On current directory, install dependecy needed by running:

```go
go mod tidy
```

4. Setup `go.mod` by changing the module name:

```go
module simple_go_gin_gorm_postgres_be_template // change this repo name to your project dirname

go 1.26.5 // change to version you see on your `go version` to match with your go version

require (
	github.com/gin-gonic/gin v1.12.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/imagekit-developer/imagekit-go/v2 v2.9.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.54.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.2
)

```

5. Setup environment variable. You need to create new file `.env` and not renaming `.env.example` (keep placeholder file so other know the environment of your project). Then copy paste all `.env.example` to your `.env`. Here's .env.example overview:

```py
RUN_MODE=development # change between development or production, I use just this to change the ip and port

# database, you can add different database for different mode
DATABASE_URI=

# jwt, you can change to your preference secret
JWT_SECRET="jwt_secret_rahasia_saya"

# host and port
DEV_HOST=127.0.0.1
DEV_PORT=8080
PROD_HOST=
PROD_PORT=

# image/video media storage sdk
IMAGEKIT_URL_ENDPOINT="https://ik.imagekit.io/your-imagekit-account-id"
IMAGEKIT_PUBLIC_KEY=
IMAGEKIT_PRIVATE_KEY=
```

6. Fix other code import with your `go.mod` module name.
7. You can try to run the app with:

```go
go run cmd/api/main.go
```

or use air live reload

```sh
air
```

## 👩‍💻 Development Guide

Development guide is explained briefly in `design-pattern.md` [here](./design-pattern.md). There I explain what the project structure are and how I develop with consistent pattern.

## 🔃 Contributing

Feel free to fork or clone this repo with some rules below:

1. Do not change module name, leave it as is
2. Do not change .air.toml, if you need change, make new file of it and put it in .gitignore
3. You can just open an issue or directly make pull request.

---

simple_go_gin_gorm_postgres_be_template licensed under MIT License. Made with love 💖 by Ahmad Fakhrul Bawani

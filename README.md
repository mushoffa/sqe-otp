# SQE OTP

# Getting Started
The infrastructure stack for this repository project comprises the following: Postgres, Redis, and Liquibase for database migration. If you have already installed Postgres & Redis running on your system, please configure the port accordingly in the [.env](config/config.go) file.

## Deployment
Please run the following command at the root of this folder repository if you have already cloned it.
```shell
make deploy
```
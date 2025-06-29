# Finance Manager
App made to handle finances

Features currently suppoerted
* Google Login
* Working on Loading KBC Transactions


## How to use
### ENV Variables
the env.go file will read in a .env file. This can be used with the dev build tag.
This will also read in PRIVATE_KEY and PUBLIC_KEY if put in root folder of the backend

#### Database
* DB_HOSTNAME : URL of DB server
* DB_PORT : Porf of DB server
* DB_PASSWORD : Password of DB server
* DB_NAME: Name of DB to use on server

#### Webserver
* REST_HOSTNAME: URL of Webserver
* REST_PORT: Port of Webserver
* PRIVATE_KEY: Security private key
* PUBLIC_KEY: Security public key
# Preparation

Before running the service you need to add in the root of this app a file called ports.json

# Running the service

To run this app you must have docker installed

Tests will run automatically as a build step

Run the following command in the root of the application
```bash
docker-compose up --build
```

This command will run 3 docker services
* mongodb
* PortDomainService
* ClientAPI

# API

You can access the client api on localhost:5000

The api contains only 1 endpoint **/ports** with queryparams **limit=uint&offset=uint** for simple pagination

The maximum amount of ports returned is 100

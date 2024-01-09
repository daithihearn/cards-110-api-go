# Cards 110 API
[![codecov](https://codecov.io/gh/daithihearn/cards-110-api-go/graph/badge.svg?token=OlyCq0RvGe)](https://codecov.io/gh/daithihearn/cards-110-api-go)

The API layer for the [Cards 110 application](https://github.com/daithihearn/cards-110)

# Requirements
To run this application you need a MongoDB instance to point to.
You will also need an application and API configured in Auth0.
You will also need a cloudinary account.

All of the above can be configured in the `.env` file.

# Technical Stack
- Go
- Swagger
- MongoDB

# Building
To build locally `make build`
To build the docker image `make image`

# Running
To run locally `make run`

# GoCoreFoundation

## Overview

Update Later

## How to run application on your local (execute only if the previous step was successful)

1. Clone the repo and cd into it
2. Copy `.example.env` file to `.env` (create first) file and fill in your input
3. Create instance database locally based on your database credentials (remember update env file)
4. Run `make migrate-up` for setup your local database or `make migrate-down` if you want to reset database
5. Run `make run` to run app on your local
6. Test api with BASE_URL is `http://localhost:8080`

## How to run application on Docker (execute only if the previous step was successful)

1. Clone the repo and cd into it
2. Copy `.example.docker.env` file to `.env.docker` (create first) file and fill in your input
3. Run `make docker-compose-up` for running container docker
4. Run `make migrate-up-docker` for setup your local database or `make migrate-down-docker` if you want to reset database
5. Test api with BASE_URL is `http://localhost:8080`

## Author Contact

Contact me with any questions!<br>

Email: anquoctpdev@gmail.com
Facebook: https://www.facebook.com/tranphuocanhquoc2003

<p style="text-align:center">Thank You so much for your time!!!</p>

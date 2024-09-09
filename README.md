## Env

Normally I wouldn't commit any of my .env files. But I didn't since it's easier than sharing access to a secure note in 1Password and having you create one yourself.

## Login

- Username: user@notreal.com or admin@notreal.com
- Password: testing

## PSQL Local Install

- Install postgres

  - Mac:

    - `brew install postgresql` but this will spin up a db
    - `brew install libpq` will install only the client

  - Ubuntu: `sudo apt install postgresql` `sudo apt-get install -y postgresql-client`

- Stop all postgres instances, sometimes they'll conflict with the docker instance.
  - `sudo systemctl stop postgresql`
- Connect to Postgres via terminal:
  - `psql postgresql://devchic:logchild@localhost:5432/resume`

## Docker Next JS Examples

- https://github.com/vercel/next.js/tree/canary/examples/with-docker-compose

## Docker - Web (Local)

- Create an image from the dockerfile
  - docker build -t test-image -f dockerfile.prod .
- Initally run the image as a container, exposing port 3000 to your local machine
  - docker run -d --name test-image --restart=always -p 3000:3000 web-test

## Docker - Nuke

- Note if you alter the schemas & init scripts in `db` then you may need to nuke your volumes, etc.

```
    rm -rf ./db/_docker_ && \
    docker compose down --volumes && \
    docker volume prune -af && \
    docker system prune -af && \
    docker compose up -d --build --renew-anon-volumes
```

<!-- If you alter this list then make sure to run the following commands: -->
<!-- docker-compose -f docker-compose.yml -f docker-development.yml down -->
<!-- docker-compose -f docker-compose.yml -f docker-development.yml stop users-db -->
<!-- docker volume prune -->
<!-- docker-compose -f docker-compose.yml -f docker-development.yml up --build -->

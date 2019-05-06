FROM postgres:latest

ENV POSTGRES_PASSWORD docker
COPY init.sql /docker-entrypoint-initdb.d/10-init.sql

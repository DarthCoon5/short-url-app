FROM postgres:latest

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres
ENV POSTGRES_HOST_AUTH_METHOD=trust

COPY ./database/ /docker-entrypoint-initdb.d
FROM node:current-alpine3.17 as build
WORKDIR /app/
RUN wget https://dl.min.io/client/mc/release/linux-amd64/mc
RUN chmod +x mc

FROM node:current-alpine3.17 as main
COPY --from=build /app/mc /usr/local/bin/mc
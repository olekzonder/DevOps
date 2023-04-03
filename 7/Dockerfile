FROM ubuntu:22.04

ARG host=8.8.8.8
ARG count=2
RUN apt update && apt -y install iputils-ping

ENV HOST $host
ENV COUNT $count
CMD ping $HOST -c  $COUNT
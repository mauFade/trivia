FROM golang:1.20.0

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y bash

CMD ["bash"]
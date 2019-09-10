FROM golang:1.13.0-alpine

COPY ./ /workspace/app/

ENTRYPOINT ["echo"]

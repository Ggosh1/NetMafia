FROM ubuntu:latest
LABEL authors="radig"

ENTRYPOINT ["top", "-b"]
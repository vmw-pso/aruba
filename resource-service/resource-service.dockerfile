# build tiny docker image
FROM alpine:latest
RUN mkdir /app
COPY resourceService /app
CMD ["app/resourceService"]
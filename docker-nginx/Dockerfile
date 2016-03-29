FROM alpine:3.2 
MAINTAINER Lee Chang <leetchang@gmail.com>

# Install nginx
RUN apk add --update nginx && \
    rm -rf /var/cache/apk/*

# Expose the ports for nginx
EXPOSE 80 443

RUN ln -sf /dev/stdout /var/log/nginx/access.log && \
    ln -sf /dev/stderr /var/log/nginx/error.log

CMD ["nginx", "-g", "daemon off;"]

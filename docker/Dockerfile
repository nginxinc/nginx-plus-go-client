# syntax=docker/dockerfile:1.3
FROM debian:buster-slim

LABEL maintainer="NGINX Docker Maintainers <docker-maint@nginx.com>"

ARG NGINX_PLUS_VERSION

# Install NGINX Plus
# Download certificate and key from the customer portal (https://my.f5.com)
# and copy to the build context
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN --mount=type=secret,id=nginx-repo.crt,dst=/etc/ssl/nginx/nginx-repo.crt,mode=0644 \
    --mount=type=secret,id=nginx-repo.key,dst=/etc/ssl/nginx/nginx-repo.key,mode=0644 \
    apt-get update \
    && apt-get install --no-install-recommends --no-install-suggests -y ca-certificates gnupg curl apt-transport-https \
    && curl -sSL https://cs.nginx.com/static/keys/nginx_signing.key | gpg --dearmor > /etc/apt/trusted.gpg.d/nginx_signing.gpg \
    && curl -sSL -o /etc/apt/apt.conf.d/90pkgs-nginx https://cs.nginx.com/static/files/90pkgs-nginx \
    && printf "%s\n" "deb https://pkgs.nginx.com/plus/debian buster nginx-plus" > /etc/apt/sources.list.d/nginx-plus.list \
    && apt-get update && apt-get install -y nginx-plus-${NGINX_PLUS_VERSION} \
    && apt-get remove --purge --auto-remove -y gnupg \
    && rm -rf /var/lib/apt/lists/* \
    && rm /etc/apt/apt.conf.d/90pkgs-nginx /etc/apt/sources.list.d/nginx-plus.list


# Forward request logs to Docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80

STOPSIGNAL SIGQUIT

RUN rm -rf /etc/nginx/conf.d/*
COPY test.conf /etc/nginx/conf.d/
COPY nginx.conf /etc/nginx/

CMD ["nginx", "-g", "daemon off;"]

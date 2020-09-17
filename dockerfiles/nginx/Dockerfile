ARG RUN_IMAGE=ubuntu:bionic

FROM ${RUN_IMAGE}
ARG GIT_SHA
RUN if [ -z "$GIT_SHA" ]; then echo "GIT_SHA not set"; exit 1; else : ; fi

RUN mkdir -p /nginx
WORKDIR /nginx
COPY *.tar.gz ./
COPY *.patch ./

# install only what is needed to run everything but `./configure` for build optimization
RUN apt-get update && apt-get install -y \
  patch

RUN tar xzvf nginx-upload-module-2.3.0.tar.gz
RUN cd nginx-upload-module-2.3.0 && \
  patch < ../upload_module_put_support.patch

RUN tar xzvf nginx-1.16.1.tar.gz
RUN tar xzvf pcre-8.42.tar.gz

# install whatever is needed for "./configure"
RUN apt-get install -y \
  gcc \
  g++ \
  libssl-dev \
  make \
  openssl \
  zlibc \
  zlib1g \
  zlib1g-dev

RUN cd nginx-1.16.1 && \
  ./configure \
    --with-debug \
    --with-pcre=../pcre-8.42 \
    --add-module=../nginx-upload-module-2.3.0 && \
  make -j $(nproc) && \
  make -j $(nproc) install

RUN groupadd -g 1000 cnb
RUN useradd -r -u 1000 -g 1000 cnb

CMD ["/usr/local/nginx/sbin/nginx", "-g", "daemon off;", "-c", "/etc/nginx/nginx.conf" ]

LABEL org.opencontainers.image.source=https://github.com/Birdrock/cf-k8s-networking.git \
            org.opencontainers.image.revision=$GIT_SHA \
            org.cloudfoundry.project.name=cf-k8s-networking \
            org.cloudfoundry.project.component.name=routecontroller \
            org.cloudfoundry.image.base.run=$RUN_IMAGE

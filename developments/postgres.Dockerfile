FROM postgres:13-alpine

RUN \
    apk --no-cache add \
    build-base \
    clang-dev \
    llvm15 \
    lz4-dev \
    msgpack-c-dev \
    zstd-dev

RUN \
    apk del \
    build-base \
    clang \
    clang-dev \
    llvm15

EXPOSE 5432
CMD ["postgres"]
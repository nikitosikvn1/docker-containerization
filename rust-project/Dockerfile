FROM rust:latest as builder

WORKDIR /app

RUN apt-get update \
    && apt-get install -y musl-tools \
    && rustup target add x86_64-unknown-linux-musl

COPY Cargo.toml Cargo.lock ./

RUN mkdir src

RUN echo "fn main() {}" > ./src/main.rs

RUN cargo build --release

RUN rm -rf ./src/

COPY . .

RUN cargo build --target x86_64-unknown-linux-musl --release

FROM gcr.io/distroless/cc-debian11:latest

COPY --from=builder /app/target/x86_64-unknown-linux-musl/release/rust-project /

CMD ["/rust-project"]

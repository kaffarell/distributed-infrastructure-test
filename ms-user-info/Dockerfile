FROM rustlang/rust:nightly as builder

# Create a new cargo project
RUN USER=root cargo new --bin ms-user-info
WORKDIR ./ms-user-info
# Copy only the Cargo.toml file and install the dependencies
COPY ./Cargo.toml ./Cargo.toml
RUN cargo build --release
# Remove the default main.rs created with cargo new
RUN rm src/*.rs

ADD . ./

# Remove the binary of the application build previously (without our code)
# The binaries of the dependencies stay
RUN rm ./target/release/deps/ms_user_info*

# Make another build with our code
RUN cargo build --release


FROM debian:buster-slim
ARG APP=/usr/src/app

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 8000

# Create new user, add to group and give folder permission
ENV APP_USER=appuser

RUN groupadd $APP_USER \
    && useradd -g $APP_USER $APP_USER \
    && mkdir -p ${APP}

# Copy the final executable from the builder to the debian instance
COPY --from=builder /ms-user-info/target/release/ms-user-info ${APP}/ms-user-info

# Set exetutable bit for user
RUN chown -R $APP_USER:$APP_USER ${APP}

USER $APP_USER
WORKDIR ${APP}

CMD ["./ms-user-info"]

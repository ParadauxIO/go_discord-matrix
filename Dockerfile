FROM alphine:3.14.2
MAINTAINER Rían Errity <rian@paradaux.io>

# Setup Envoirnment
RUN mkdir -p /app

# Copy files in
COPY go_discord-matrix /app/go_discord-matrix 
COPY config.json /app/config.json

WORKDIR /app
ENTRYPOINT ["./go_discord-matrix"]

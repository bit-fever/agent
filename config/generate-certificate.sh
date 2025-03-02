#!/bin/sh
 
# Create a key for the agent
openssl genrsa -out agent.key 2048

# Generate the Certificate Signing Request 
openssl req -new -key agent.key -out agent.csr -subj "/C=EU/ST=Italy/L=Rome/O=BitFever/OU=BitFeverAgent/CN=bitfever-agent" 

echo "subjectAltName=DNS:bitfever-agent" > altsubj.ext

openssl x509 -req -in agent.csr -key agent.key \
    -days 20000 -sha256 \
    -extfile altsubj.ext \
    -out agent.crt

rm altsubj.ext

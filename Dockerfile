FROM alpine:3.19.0
COPY ddns /app/ddns
# RUN apk add --no-cache bash
# RUN chmod +x /app/ddns
ENTRYPOINT ["/app/ddns","start"]
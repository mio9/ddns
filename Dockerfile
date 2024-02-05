FROM debian:bookworm-slim
# COPY ddns /app/ddns
# RUN apk add --no-cache bash
# RUN chmod +x /app/ddns
ENTRYPOINT ["/app/ddns","start"]
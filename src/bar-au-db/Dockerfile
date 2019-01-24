# Base image
FROM postgres:11.1-alpine

# Expose ports
EXPOSE 5432

# Add liveness probe
ADD ./liveness-probe.sh /usr/local/bin/liveness-probe
RUN chmod +x /usr/local/bin/liveness-probe
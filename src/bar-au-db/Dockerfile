# migrate image
FROM migrate/migrate:v4.2.2 as migrate

# Base image
FROM postgres:11.1-alpine

# Expose ports
EXPOSE 5432

# Copy migrations files
ADD migrations /migrations

# Copy migrate executable, will be executed on PostStart hook
COPY --from=migrate /migrate /usr/local/bin/migrate 

# Add script allow to wait the postgre database to be up
ADD ./wait-for-postgres.sh /usr/local/bin/wait-for-postgres
RUN chmod +x /usr/local/bin/wait-for-postgres

# Add migrate database script
ADD ./migrate-up.sh /usr/local/bin/migrate-up
RUN chmod +x /usr/local/bin/migrate-up

# Add liveness probe
ADD ./liveness-probe.sh /usr/local/bin/liveness-probe
RUN chmod +x /usr/local/bin/liveness-probe
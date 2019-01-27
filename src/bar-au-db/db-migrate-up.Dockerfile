FROM migrate/migrate:v4.2.2

# add bash and postgre client
RUN apk add --no-cache bash postgresql-client

# print executed commands
RUN set -x

# add the migration files
ADD migrations /migrations

# add script for waiting the database is up
ADD wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh

# run migrate up
#  CMD migrate -verbose -path "/migrations" -database ${DATABASE_URL up
ENTRYPOINT  /wait-for-postgres.sh && \
    ./migrate -verbose -path "/migrations" -database $DATABASE_URL up

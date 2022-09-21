# Base image
FROM mysql:5.7

# Copy custom conf files
# from ./conf (relative host path from this Dockerfile)
# to /etc/mysql/conf.d/ in the VM
COPY ./conf/* /etc/mysql/conf.d/

COPY ./debezium-config/db.sql /docker-entrypoint-initdb.d/

# Make the conf files copied not writeable so mysql will read them
RUN chmod a-w /etc/mysql/conf.d/*
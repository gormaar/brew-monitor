FROM mysql:latest as builder

ADD brew-monitor_db.sql /docker-entrypoint-initdb.d/
RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

RUN ["sed", "-i", "s/exec \"$@\"/echo \"not running $@\"/", "/usr/local/bin/docker-entrypoint.sh"]

ARG DB_USER
ARG DB_PASSWORD
ARG DB_DATABASE
ENV MYSQL_USER=${DB_USER}
ENV MYSQL_PASSWORD=${DB_PASSWORD}
ENV MYSQL_ROOT_PASSWORD=${DB_PASSWORD}
ENV MYSQL_DATABASE=${DB_DATABASE}

WORKDIR /app

RUN ["/usr/local/bin/docker-entrypoint.sh", "mysqld", "--datadir", "/initialized-db"]


FROM mysql:latest

COPY --from=builder /initialized-db /var/lib/mysql

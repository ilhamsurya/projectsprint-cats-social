FROM postgres:14.5
RUN rm -rf /docker-entrypoint-initdb.d/*

ADD ./pkg/database/postgre/migration/1_schema.sql /docker-entrypoint-initdb.d

RUN chmod a+r /docker-entrypoint-initdb.d
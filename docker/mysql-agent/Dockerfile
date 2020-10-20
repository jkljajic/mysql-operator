FROM mysql/mysql-server:8.0.22

#adding community mysql
RUN groupadd --gid 999 mysqlcom
RUN useradd -u 999 mysqlcom --gid 999
COPY bin/linux_amd64/mysql-agent /

#USER mysql

ENTRYPOINT ["/mysql-agent"]

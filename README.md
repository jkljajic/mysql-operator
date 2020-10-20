# MySQL Operator

The MySQL [Operator][1] creates, configures and manages MySQL InnoDB clusters running on Kubernetes. It is not for MySQL NDB Cluster.

[![issues](https://img.shields.io/github/issues/oracle/mysql-operator.svg)](https://github.com/oracle/mysql-operator/issues)
[![tags](https://img.shields.io/github/tag/oracle/mysql-operator.svg)](https://github.com/oracle/mysql-operator/tags)
[![wercker status](https://app.wercker.com/status/cc1710e8b354d1a22f36b04c8313eac9/s/master "wercker status")](https://app.wercker.com/project/byKey/cc1710e8b354d1a22f36b04c8313eac9)
[![Go Report Card](https://goreportcard.com/badge/github.com/oracle/mysql-operator)](https://goreportcard.com/report/github.com/oracle/mysql-operator)

The MySQL Operator is opinionated about the way in which clusters are configured.
We build upon [InnoDB cluster][3] and [Group Replication][4] to provide a complete high
availability solution for MySQL running on Kubernetes.

**While fully usable, this is currently alpha software and should be treated as
such.  You are responsible for your data and the operation of your database clusters. There may be backwards incompatible changes up until the first major
release.**

## Getting started

See the [tutorial][5] which provides a quick-start guide for users of the Oracle MySQL Operator.

## Features

The MySQL Operator provides the following core features:

- Create and delete highly available MySQL InnoDB clusters in Kubernetes with minimal effort
- Automate database backups, failure detection, and recovery
- Schedule automated backups as part of a cluster definition
- Create "on-demand" backups.
- Use backups to restore a database

## Requirements

 * Kubernetes 1.8.0 +

## Contributing

`mysql-operator` is an open source project. See [CONTRIBUTING](CONTRIBUTING.md) for
details.

Oracle gratefully acknowledges the contributions to this project that have been made
by the community.

## License

Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.

`mysql-operator` is licensed under the Apache License 2.0.

See [LICENSE](LICENSE) for more details.

TODO:
Add Dynamc generation ServiceAgentName and clusterrole & binder
80% Add service for primary and secondary selector by label ROLE
ADD refresh servuce when some changes acured 
ADD CLEAN UP AFTER DOWN resources


[1]: https://coreos.com/blog/introducing-operators.html
[2]: https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/
[3]: https://dev.mysql.com/doc/refman/5.7/en/mysql-innodb-cluster-userguide.html
[4]: https://dev.mysql.com/doc/refman/5.7/en/group-replication.html
[5]: docs/tutorial.md

kubectl run -it --rm sysbench-client --image=perconalab/sysbench:latest --restart=Never -- bash
sysbench oltp_read_only --tables=10 --table_size=1000000  --mysql-host=mysql-server --mysql-user=root --mysql-password=SGI8c0DNQYEP3s0o --mysql-db=sbtest   prepare
sysbench oltp_read_only --tables=10 --table_size=1000000  --mysql-host=mysql-server --mysql-user=root --mysql-password=SGI8c0DNQYEP3s0o --mysql-db=sbtest  --mysql-ignore-errors=1180,1213,1020,1205,3100,3101 --time=300 --threads=16 --report-interval=1 run
sysbench oltp_read_write --tables=10 --table_size=1000000  --mysql-host=mysql-server --mysql-user=root --mysql-password=SGI8c0DNQYEP3s0o --mysql-db=sbtest --mysql-ignore-errors=1180,1213,1020,1205,3100,3101 --time=300 --threads=16 --report-interval=1 run


kubectl run -it --rm sysbench-client --image=perconalab/sysbench:latest --restart=Never -- bash
#./tpcc.lua --mysql-host=172.16.0.11 --mysql-user=sbtest --mysql-password=sbtest --mysql-db=sbtest --time=300 --threads=64 --report-interval=1 --tables=10 --scale=100 --db-driver=mysql --use_fk=0 --force_pk=1 --trx_level=RC prepare
#./tpcc.lua --mysql-host=172.16.0.12,172.16.0.13 --mysql-user=sbtest --mysql-password=sbtest --mysql-db=sbtest --time=$time --threads=$i --report-interval=1 --tables=10 --scale=100 --trx_level=RR  --db-driver=mysql --report_csv=yes --mysql-ignore-errors=1180,1213,1020,1205,3100,3101 run




sysbench oltp_read_only --tables=10 --table_size=1000000  --mysql-host=cluster1-haproxy -uroot -proot_password --mysql-user=root --mysql-password=root_password --mysql-db=sbtest   prepare
sysbench oltp_read_only --tables=10 --table_size=1000000  --mysql-host=cluster1-haproxy --mysql-user=root --mysql-password=root_password --mysql-db=sbtest  --time=300 --threads=16 --report-interval=1 run
sysbench oltp_read_write --tables=10 --table_size=1000000  --mysql-host=mcluster1-haproxy --mysql-user=root --mysql-password=root_password --mysql-db=sbtest  --time=300 --threads=16 --report-interval=1 run


helm upgrade mysql-operator --namespace kube-operators ./mysql-operator/
#!/bin/bash

set -e

function usage
{
    echo "usage: `basename $0` (-c || -s) [-f CASSANDRA_REPLICATION -b BUCKET_NAME]"
    echo "   ";
    echo "  -c | --cluster           : Cluster initialization";
    echo "  -s | --standalone        : Standalone initialization";
    echo "  -f | --cassandra_factor  : Cassandra replication factor (default: 1 for standalone, 3 for cluster)";
    echo "  -b | --mino_bucket_name  : The name of minio bucket (default: snapshot-alert-feature-db)";
    echo "  -h | --help              : This message";
}

function parse_args
{
    while [ "$1" != "" ]; do
    	case "$1" in
            -c | --cluster )              cluster="true";;
            -s | --standalone )           standalone="true";;
            -f | --cassandra_factor )     cassandra_factor="$2";;
            -b | --mino_bucket_name )     bucket_name="$2";;
            -h | --help )                 usage;                     exit;;
        esac
        shift
    done

    if [[ -z "${cluster}" && -z "${standalone}" ]] || [[ -n "${cluster}" && -n "${standalone}" ]]; then
        echo -e "should set -c or -s\n"
        usage
        exit;
    fi

    if [[ -z "${bucket_name}" ]]; then
        bucket_name="snapshot-alert-feature-db"
    fi

    if [[ -n "${cluster}" ]]; then
        if [[ -z "${cassandra_factor}" ]]; then
            cassandra_factor="3"
        fi
    fi

    if [[ -n "${standalone}" ]]; then
        cassandra_factor="1"
    fi
}

function init
{
/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "CREATE KEYSPACE IF NOT EXISTS viper_test WITH replication = {'class':'SimpleStrategy', 'replication_factor' : ${cassandra_factor} };"

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "
CREATE TABLE IF NOT EXISTS viper_test.static_feature_dbs(
    shard_key int,
	db_id uuid,
	object_type text,
	name text,
	feature_version int,
	description text,
	creation_time timestamp,
	indexes map<uuid, text>,
	deleted boolean,
	max_size bigint,
	PRIMARY KEY (shard_key, db_id),
);"

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "
CREATE TABLE IF NOT EXISTS viper_test.static_features(
    index_id uuid,
    seq_id bigint,
    feature_version int,
    creation_time timestamp,
    metadata blob,
    image_id text,
    payload text,
    feature blob,
    user_key text,
    PRIMARY KEY (index_id, seq_id),
);"

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "
ALTER TABLE viper_test.static_features WITH gc_grace_seconds = 86400;" || true

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "
CREATE MATERIALIZED VIEW IF NOT EXISTS viper_test.static_features_by_user_key AS
    SELECT user_key, feature_version, creation_time, metadata, image_id, payload FROM static_features
    WHERE user_key IS NOT NULL AND index_id IS NOT NULL AND seq_id IS NOT NULL
    PRIMARY KEY (index_id, user_key, seq_id);"
}

function run
{
	parse_args $@
	echo "bucket_name: ${bucket_name}"
	echo "cassandra_factor: ${cassandra_factor}"
    init
}

run $@

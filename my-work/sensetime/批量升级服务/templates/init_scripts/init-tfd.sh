#!/bin/bash

set -e

function usage
{
    echo "usage: `basename $0` (-c || -s) [-cf CASSANDRA_REPLICATION -kf KAFKA_REPLICATION -p TOPIC_PARTITIONS -b BUCKET_NAME -v FEATURE_VERSION]"
    echo "   "
    echo "  -c  | --cluster          : Cluster initialization"
    echo "  -s  | --standalone       : Standalone initialization"
    echo "  -cf | --cassandra_factor : Cassandra replication factor (default: 1 for standalone, 3 for cluster)"
    echo "  -kf | --kafka_factor     : Kafka replication factor (default: 1 for standalone, 3 for cluster)"
    echo "  -p  | --partitions       : The partitions num of the topic (default: 8 for standalone, 32 for cluster)"
    echo "  -b  | --mino_bucket_name : The name of minio bucket (default: snapshot-timespace-feature-db)"
    echo "  -v  | --feature_version  : The version of feature (default: 24702)"
    echo "  -dn | --deploy_num       : The number of tfd deployment (default: 1 for standalone, 2 for cluster)"
    echo "  -h  | --help             : This message"
}

function parse_args
{
    while [ "$1" != "" ]; do
    	case "$1" in
            -c  | --cluster )          cluster="true";;
            -s  | --standalone )       standalone="true";;
            -cf | --cassandra_factor ) cassandra_factor="$2";;
            -kf | --kafka_factor )     kafka_factor="$2";;
            -p  | --partitions )       partitions="$2";;
            -b  | --mino_bucket_name ) bucket_name="$2";;
            -v  | --feature_version )  feature_version="$2";;
            -dn | --deploy_num )       deploy_num="$2";;
            -h  | --help )             usage; exit;;
        esac
        shift
    done

    if [[ -z "${cluster}" && -z "${standalone}" ]] || [[ -n "${cluster}" && -n "${standalone}" ]]; then
        echo -e "should set -c or -s\n"
        usage
        exit;
    fi

    if [[ -z "${bucket_name}" ]]; then
        bucket_name="snapshot-timespace-feature-db"
    fi

    if [[ -z "${feature_version}" ]]; then
        feature_version="24702"
    fi

    if [[ -n "${cluster}" ]]; then
        if [[ -z "${cassandra_factor}" ]]; then
            cassandra_factor="3"
        fi
        if [[ -z "${kafka_factor}" ]]; then
            kafka_factor="3"
        fi
        if [[ -z "${partitions}" ]]; then
            partitions="32"
        fi
        if [[ -z "${deploy_num}" ]]; then
            deploy_num=2
        fi
    fi

    if [[ -n "${standalone}" ]]; then
        cassandra_factor="1"
        kafka_factor="1"
        deploy_num=1
        if [[ -z "${partitions}" ]]; then
            partitions="8"
        fi
    fi
}

function init_cassandra
{
suffix=$1

keyspace="face_${feature_version}_${suffix}"

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "CREATE KEYSPACE IF NOT EXISTS ${keyspace} WITH replication = {'class':'SimpleStrategy', 'replication_factor' : ${cassandra_factor} };"

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "
CREATE TABLE IF NOT EXISTS ${keyspace}.indexes (
	region_id int,
	index_id uuid,
	first_time timestamp,
	last_time timestamp,
	shard_size int,
	status int,
	worker_id text,
	PRIMARY KEY (region_id, index_id),
);"

/usr/bin/kubectl exec -it -n component cassandra-default-0 -- cqlsh -e "
CREATE TABLE IF NOT EXISTS ${keyspace}.features (
	region_id int,
	captured_date int,
	captured_time timestamp,
	camera_idx int,
	sequence int,
	annotation blob,
	cluster_id bigint,
	extra_info text,
	feature blob,
	panoramic_image_url text,
	portrait_image_url text,
	PRIMARY KEY ((region_id, captured_date), captured_time, camera_idx, sequence),
);"
}

function init
{
for i in `seq 1 ${deploy_num}`; do
    init_cassandra ${i}
done

/usr/bin/kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --create --zookeeper zookeeper-default:2181/kafka -topic stream.features.face_${feature_version} -replication-factor ${kafka_factor} -partitions ${partitions} --if-not-exists

/usr/bin/kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --create --zookeeper zookeeper-default:2181/kafka -topic sync.stream.features.face_${feature_version} -replication-factor ${kafka_factor} -partitions ${partitions} --if-not-exists

/usr/bin/kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --create --zookeeper zookeeper-default:2181/kafka -topic senseguard.bulk.tool_${feature_version} -replication-factor ${kafka_factor} -partitions ${partitions} --if-not-exists

}

function run
{
    parse_args $@
    echo "cassandra_factor: ${cassandra_factor}"
    echo "kafka_factor: ${kafka_factor}"
    echo "partitions: ${partitions}"
    echo "feature_version: ${feature_version}"
    echo "bucket_name: ${bucket_name}"
    echo "deploy_num: ${deploy_num}"
    init
}

run $@

#!/bin/bash

set -e

function usage
{
    echo "usage: `basename $0` (-c || -s) [-f KAFKA_REPLICATION -p TOPIC_PARTITIONS]"
    echo "   ";
    echo "  -c | --cluster      : Cluster initialization";
    echo "  -s | --standalone   : Standalone initialization";
    echo "  -f | --kafka_factor : Kafka replication factor(default: 1 for standalone, 3 for cluster)";
    echo "  -p | --partitions   : The partitions num of the topic(default: 8 for standalone, 32 for cluster)";
    echo "  -h | --help         : This message";
}

function parse_args
{
    while [ "$1" != "" ]; do
    	case "$1" in
            -c | --cluster )        cluster="true";;
            -s | --standalone )     standalone="true";;
            -f | --kafka_factor )   kafka_factor="$2";;
            -p | --partitions )     partitions="$2";;
            -h | --help )           usage;                     exit;;
        esac
        shift
    done

    if [[ -z "${cluster}" && -z "${standalone}" ]] || [[ -n "${cluster}" && -n "${standalone}" ]]; then
        echo -e "should set -c or -s\n"
        usage
        exit;
    fi

    if [[ -n "${cluster}" ]]; then
        if [[ -z "${kafka_factor}" ]]; then
            kafka_factor="3"
        fi
        if [[ -z "${partitions}" ]]; then
            partitions="32"
        fi
    fi

    if [[ -n "${standalone}" ]]; then
        kafka_factor="1"
        if [[ -z "${partitions}" ]]; then
            partitions="8"
        fi
    fi
}

function init
{
/usr/bin/kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --create --zookeeper zookeeper-default:2181/kafka -topic stream.rws.td.comparison -replication-factor ${kafka_factor} -partitions ${partitions} --if-not-exists
}

function run
{
	parse_args $@
	echo "kafka_factor: ${kafka_factor}"
	echo "partitions: ${partitions}"
    init
}

run $@

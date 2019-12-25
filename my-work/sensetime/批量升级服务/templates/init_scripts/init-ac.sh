#!/bin/bash

set -e

function usage
{
    echo "usage: `basename $0` (-c || -s) [-o OSG_REPLICATION -k TOPIC_FACTOR]"
    echo "   ";
    echo "  -c | --cluster           : Cluster initialization";
    echo "  -s | --standalone        : Standalone initialization";
    echo "  -o | --osg_replication   : OSG replication factor(default: 000 for standalone, 010 for cluster)";
    echo "  -k | --kafka_factor      : Kafka topic factor(default: 1 for standalone, 3 for cluster)";
    echo "  -p | --partitions        : Kafka topic partition number(default: 8 for standalone, 32 for cluster)";
    echo "  -h | --help              : This message";
}

function parse_args
{
  while [ "$1" != "" ]; do
      case "$1" in
          -c | --cluster )              cluster="true";;
          -s | --standalone )           standalone="true";;
          -o | --osg_replication )      osg_replication="$2";;
          -k | --kafka_factor )         kafka_factor="$2";;
          -p | --partitions )           partitions="$2";;
          -h | --help )                 usage;                     exit;;
      esac
      shift
  done

  if [[ -z "${cluster}" && -z "${standalone}" ]] || [[ -n "${cluster}" && -n "${standalone}" ]]; then
      echo -e "should set -c or -s\n"
      usage
      exit;
  fi

  if [[ -n "${cluster}" ]]; then
	  if [[ -z "${osg_replication}" ]]; then
	      osg_replication="010"
	  fi
	  if [[ -z "${kafka_factor}" ]]; then
	      kafka_factor="3"
	  fi
    if [[ -z "${partitions}" ]]; then
        partitions="32"
    fi
  fi

  if [[ -n "${standalone}" ]]; then
      osg_replication="000"
      kafka_factor="1"
      partitions="8"
  fi
}

function init
{
  kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --create --zookeeper zookeeper-default:2181/kafka -topic stream.sensekeeper.rwstsdb -replication-factor "${kafka_factor}" -partitions "${partitions}" --if-not-exists || true
  kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --create --zookeeper zookeeper-default:2181/kafka -topic stream.sensekeeper.biz -replication-factor "${kafka_factor}" -partitions "${partitions}" --if-not-exists || true

  curl -X PUT -d '{"bucket_info":{"name":"keeper_face", "attrs":{"ttl":"90d", "replication":"'"${osg_replication}"'"}}}' http://0.0.0.0/components/osg-default/v1
}

function run
{
  parse_args "$@"
  init
}

run "$@";

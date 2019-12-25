#!/bin/bash

function usage
{
    echo "usage: `basename $0` --host host [[-t | --ttl] expired-days]"
    echo "   ";
    echo "  --host        : bucket host";
    echo "  -t | --ttl    : expired-days";
    echo "  -h | --help   : This message";
}

function is_number
{
    re='^[0-9]+$'
    if [[ $1 =~ $re ]] ; then
        return 0
    else
        return 1
    fi
}

function parse_args
{
    while [ "$1" != "" ]; do
        case "$1" in
            --host )      host="$2";;
            -t | --ttl  ) ttl="$2";;
            -h | --help ) usage;  exit;;
        esac
        shift
    done

    if [ -z ${host} ]; then
        usage
        exit 1
    fi

    is_number ${ttl}
    if [ $? -ne 0 ]; then
        echo "ttl is not a positive number"
        exit 1
    fi
}


function create_bucket
{
    bucket_name=$1
    bucket_ttl=$2

    req_body="\"name\":\"${bucket_name}\""
    if [ -n "${bucket_ttl}" ]; then
        req_body="${req_body}, \"attrs\":{\"ttl\":\"${bucket_ttl}d\"}"
    fi
    #echo "{\"bucket_info\": {${req_body}}}"
    curl -X PUT "http://${host}/components/osg-default/v1" --data "{\"bucket_info\": {${req_body}}}"
}

parse_args $@
create_bucket "GLOBAL" ""
create_bucket "video_face" ${ttl}
create_bucket "video_panoramic" ${ttl}
create_bucket "keeper_face" ${ttl}
create_bucket "OTHER" ${ttl}
create_bucket "GUEST" ${ttl}
create_bucket "FRONTEND" ${ttl}

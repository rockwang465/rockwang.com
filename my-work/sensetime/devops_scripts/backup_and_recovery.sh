#!/bin/bash

#mysql backup && recovery脚本
#执行环境:直接在部署有mysql的k8s集群上执行脚本
#使用方法
#backup
#./backup_and_recovery.sh -m backup -p /xxx/xxx/xxx/backup_dir
#recovery
#./backup_and_recovery.sh -m restore -f /xxx/xxx/xxx/xx.sql

#set -x

DATABASES="console mysql"
EXECCMD="/usr/bin/kubectl -n component exec"
RESERVE_DAYS=7
mariadb_pod_name="mariadb-component-master-0"

function usage
{
    echo "usage: `basename $0` -m [backup | restore] [-p path | -f file] [--pod-name <real-pod-name>]"
    echo "   ";
    echo "  -m | --method: backup or restore";
    echo "  -p | --path: the path where backup to";
    echo "  -f | --file: the file which restore from";
    echo "  --pod-name: the pod name of mariadb: one of mariadb-component-master-0 and mariadb-component-slave-0";
}

function parse_args
{
    while [ "$1" != "" ]; do
    	case "$1" in
            -m | --method ) method="$2";;
            -p | --path )   backup_path="$2";;
            -f | --file )   backup_file="$2";;
            --pod-name )   mariadb_pod_name="$2";;
            -h | --help )   usage;  exit;;
        esac
        shift
    done

    if [[ -z "${method}" ]]; then
        echo -e "should set -m\n"
        usage
        exit;
    fi

    if [[ -z "${mariadb_pod_name}" ]]; then
        echo -e "should set --pod-name\n"
        usage
        exit;
    fi

    if [[ "${method}" == "backup" ]] && [[ -z ${backup_path} ]]; then
       	echo -e "where is backup_path?\n"
       	usage
        exit;
    fi

    if [[ "${method}" == "restore" ]] && [[ -z ${backup_file} ]]; then
       	echo -e "where is backup_file?\n"
       	usage
        exit;
    fi
}

function is_cmd_exist {
    cmd=$1
    which "${cmd}" > /dev/null
    if [ $? -ne 0 ]; then
        echo "${cmd} not exists, exit"
        exit 1
    fi
}

function pre_check {
    is_cmd_exist "/usr/bin/kubectl"
}

function get_token {
    mysql_cnf="~/.my.cnf"
    if [[ ! -f ${mysql_cnf} ]]; then
        ${EXECCMD} -it ${mariadb_pod_name} -- /bin/bash -c '. /.bashrc && echo -e "[client]\npassword=${MARIADB_ROOT_PASSWORD}" > /bitnami/mariadb/.my.cnf'
    fi
}

function get_mariadb_server {
    pod_number=$(kubectl get pods -n component ${mariadb_pod_name} | grep mariadb |wc -l)
    if [[ pod_number -lt 1 ]]; then
        echo "nil Running mariadb pod"
        exit 2
    fi
}

function check_error {
    if [ $1 -ne 0 ]; then
        echo "$2"
        exit $3
    fi
}


function backup_database {
    l_db=$1
    l_path=$2
    l_date=$3
    l_file="${l_path}/${l_db}-${l_date}.sql"
    echo "backup_database ${l_db} to ${l_file}"
    ${EXECCMD} -it ${mariadb_pod_name} -- mysqldump --defaults-extra-file=/bitnami/mariadb/.my.cnf -uroot ${l_db} > ${l_file}
}

function clean_old_backup {
    n_days_ago=`date +%Y%m%d -d"${RESERVE_DAYS} days ago"`
    old_tgz="${backup_path}/mariadb_backup_${n_days_ago}.tgz"
    if [ -f ${old_tgz} ]; then
        rm ${old_tgz}
    fi
}

function backup {
    backup_date=`date +%Y%m%d`
    ok=true
    for db in ${DATABASES}; do
        backup_database "${db}" "${backup_path}" "${backup_date}"
        if [ $? -ne 0 ]; then
            ok=false
        fi
    done
    if [ ${ok} == true ]; then
        cd ${backup_path}
        tar zcf mariadb_backup_${backup_date}.tgz *-${backup_date}.sql
        if [ $? -ne 0 ]; then
            echo "tar zcf mariadb_backup_${backup_date}.tgz fail"
            exit 1
        fi
        rm *-${backup_date}.sql
    else
        echo "backup_database fail!"
        exit 1
    fi
    clean_old_backup
}

function restore_database {
    l_db=$1
    l_file=$2
    echo "restore_database ${l_db} from ${l_file}"
    ${EXECCMD} -i ${mariadb_pod_name} -- mysql --defaults-extra-file=/bitnami/mariadb/.my.cnf -uroot ${l_db} < ${l_file}
}

function restore {
    base_dir=`dirname ${backup_file}`
    backup_date=`basename ${backup_file} | sed 's/mariadb_backup_\(.*\).tgz/\1/'`
    tar zxvf ${backup_file} -C "${base_dir}"

    for db in ${DATABASES}; do
        backup_sql_file="${base_dir}/${db}-${backup_date}.sql"
        if [ ! -f ${backup_sql_file} ]; then
            echo "ERROR: ${backup_sql_file} not exist"
            exit 1
        fi
        restore_database "${db}" "${backup_sql_file}"
        rm "${backup_sql_file}"
    done
}

function run {
    parse_args $@
    pre_check
    get_mariadb_server
    get_token

    set -u
    if [ "${method}" == "backup" ]; then
        backup
    else
        restore
    fi
}

run $@
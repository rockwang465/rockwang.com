#!/bin/bash

#######################################
# author:  DengJunYu                  #
# time:    2020-4-28                  #
# update:                             #
# version: v.1.0                      #
#######################################
function usage
{
    echo "usage: `basename $0`  [-i IP -h ]"
    echo "   ";
    echo "  -i | --ip                : Network ipv4 address(default: local host ip automatic acquisition)";
    echo "  -h | --help              : This message";

}

function parse_args
{
    while [ "$1" != "" ]; do
        case "$1" in
          -i | --ip )                   ip_addr="$2";;
          -h | --help )                 usage;                     exit;;

        esac
        shift
    done

    if [ -z "${ip_addr}"  ]; then
        ip_addr=`hostname -i`
    fi

}

function encryption_rules
{
   num1=`echo ${ip_addr}|awk -F '.' '{print int($2)}'`
   num2=`echo ${ip_addr}|awk -F '.' '{print int($3)}'`
   num3=`echo ${ip_addr}|awk -F '.' '{print int($4)}'`
   #echo $num1 $num2 $num3
   sum=$((${num1}+${num2}+${num3}))
   #echo $sum
   if [ `expr $sum / 1000` -ge 1 ];then
        filling=$sum
   elif [ `expr $sum / 100` -ge 1 ];then
        filling=`expr $sum \* 10`
   elif [ `expr $sum / 10` -ge 1 ];then
        filling=`expr $sum \* 100`
   fi
   #echo $filling
   encryption="IdeA!#$filling@$"
   echo "this server root password is: $encryption"
   echo "this server root password is: $encryption" > /root/password.txt



}

function change_password
{
   echo $encryption |passwd --stdin root
}

function run
{
  parse_args "$@"
  encryption_rules
  change_password
}

run "$@";

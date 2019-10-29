#!/bin/sh

args_num=$#
args1=$1
args2=$2

function judge_args(){
    if [${args_num} == 1 ]; then
        if [ ${args1} == "-h" || ${args1} == "--help" ]; then
            echo -e "Usage: $0 vm_name vm_disk_name"
        fi
    elif [ ${args_num} != 2 ]; then
        echo -e "Error : Must input 2 arguments."

    else
        echo -e "Error: Input argument error, you can use '-h' to help"
    fi
    echo -e "\n"
}

function create_disk(){
    echo -e "vm disk is [$args2]"
    vm_disk=$args2

    if [ -e /data/data_root/vm-images/${vm_disk}.img ]; then
        echo -e "Error : Already exists /data/data_root/vm-images/${vm_disk}.img"
        exit 1
    else
        echo -e "create vm image ${vm_disk}"
        #qemu-img create -f qcow2 /data/data_root/vm-images/${vm_disk}.img 240G
    fi
    echo -e "\n"
}


function install_vm(){
    echo "e"



}


judge_args
create_disk

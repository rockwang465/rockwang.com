#!/bin/bash
user=root
crontab=/var/spool/cron/crontabs/${user}
home_dir=`pwd`
clean_script=cron_clean.py
guard_script=guard_sky_eye_v2.sh

function set_cron(){
    printf "******** setting crontab ********\n"
    #set guard
    cron_count=`cat $crontab | grep ${guard_script} | grep -v grep | wc -l`
    if [ $cron_count == 0 ]; then
        echo "* * * * * /bin/bash ${home_dir}/${guard_script}"
        echo "* * * * * /bin/bash ${home_dir}/${guard_script} >${home_dir}/guard.log" >> $crontab
    fi


    #set clean data
    cron_count=`cat $crontab | grep ${clean_script} | grep -v grep | wc -l`
    if [ $cron_count == 0 ]; then
        echo "* * * * * /home/youtuapp/anaconda2/bin/python  ${home_dir}/${clean_script}"
        echo "* * * * * /home/youtuapp/anaconda2/bin/python  ${home_dir}/${clean_script} 30" >> $crontab
    fi
    sudo chown ${user}:${user} ${crontab}
    sudo chmod 600 ${crontab}


    #check set crontab status
    num2=`cat $crontab | grep ${guard_script} | grep -v grep | wc -l`
    num3=`cat $crontab | grep ${clean_script} | grep -v grep | wc -l`
    if [ ${num2} -eq 1 -a ${num3} -eq 1 ];then
        echo "set crontab task has been successful."
        #echo "num2:$num2,num3:$num3"
    else
        echo "set crontab task error......"
        exit
    fi
    echo -e "\n"
    sleep 3
}

set_cron

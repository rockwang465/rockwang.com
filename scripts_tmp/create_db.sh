#！/bin/bash

#create mysql database
host="localhost"   
port="3306"
username="root"
password="skyeye"
dbname="skyeye"

#create database skyeye
function create_db(){
  #1.check skyeye database exist
  mysql -uroot -p${password} -e "use skyeye;" 
  #import database
  if [ $? -ne 0 ];then
      create_db_sql="create database IF NOT EXISTS ${dbname}"
      mysql -h${host}  -P${port}  -u${username} -p${password} -e "${create_db_sql}"
      echo "mysql database  create successful."
  else
      echo "skyeye database already exist."
  fi
  sleep 2

  #2.check tables user_info exist
  mysql -uroot -p${password} -e "desc skyeye.user_info;" >/dev/null 2>&1 
  #impor  tables
  num1=`echo $?`
  if [ $? -ne 0 ];then
      mysql -h${host} -u${username} -p${password} ${dbname} < skyeye.sql 
      echo "mysql table create successful"
  else
      echo "user_info table already exist."
  fi
  sleep 2

  #3.check mysql user:root@%
  num2=`mysql -uroot -pskyeye -e "select * from mysql.user where user='root';" | grep  "^%.*root" | wc -l`
  #set remote connect
  if [ ${num2} -eq 0 ];then
      mysql -h${host}  -P${port}  -u${username} -p${password} -e "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'skyeye' WITH GRANT OPTION;flush privileges;"
      echo "mysql GRANT root user add successful"
  elif [ ${num2} -eq 1 ];then
      echo "mysql 'root@%' user already exist"
  else
      echo "mysql 'root@%' user has error......"
  fi
  echo -e "\n"
  sleep 3

    
  #4.check all
  if [ ${num1} -eq 0 -a ${num2} -eq 1 ];then
      echo "mysql table and user already exist."
  elif [ ${num1} -ne 0 -a $[num2] -eq 0 ];then
      cd ~/code/sky_eye_v2/stream_service
      bash install_srs.bash
      sudo cp srs/srs.conf /usr/local/srs/conf/srs.conf
      sudo ldconfig
  else
      echo "Waring: table status or mysql user status abnormal [need check]......"
  fi
}


create_db

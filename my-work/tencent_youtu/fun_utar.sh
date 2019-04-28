#!/bin/bash
now_dir=`pwd`
default_dir="/home/youtuapp/1.html-code/sky_eye_v2"

#mdkir -p default_dir
function mkdir_sky_eye(){
  if [ ! -d ${default_dir} ];then 
      mdkir -p ${default_dir}
      echo "${default_dir} path create successful."
  else
      echo "${default_dir} path already exist."
  fi
  echo -e "\n"
  sleep 4
}

#utar now_dir skyeyev2.4.1.zip
function skyeye_untar(){
  unzip skyeyev2.4.1.zip
  if [ $? -eq 0 ];then
      echo "unzip skyeyev2.4.1.zip successful."
      echo -e "\n"
      sleep 4
  fi
}

#copy *.tar.gz to /home/youtuapp/1.html-code/sky_eye_v2
function cp_all_file_to_sky_eye(){
  cd ./skyeyev2.4.1
  cp *.tar.gz ${default_dir}
  echo "copy *.tar.gz to ${default_dir} successful."
  echo -e "\n"
  sleep 4
}

#decompressing  /home/youapp/1.html-code/sky_eye_v2/*.tar.gz files
function untar(){
  cd ${default_dir}
#  find ${default_dir} -name '*.tar.gz' | grep -v scripts*.tar.gz | xargs -i tar -zxvf {} -C ${default_dir}
  find ${default_dir} -name '*.tar.gz' | grep -v scripts*.tar.gz | xargs -i tar -zxvf {} -C ${default_dir}

  value=`find ${default_dir} -name '*.tar.gz' | grep -v scripts*.tar.gz | wc -l`
  if [ ${value} -eq 0 ];then
      echo "Not found *.tar.gz files"
  else
      echo "untar ${default_dir} *.tar.gz successful."
  fi
  echo -e "\n"
  sleep 4
}

mkdir_sky_eye
skyeye_untar
cp_all_file_to_sky_eye
untar






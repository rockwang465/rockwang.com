#!/bin/bash
passwd="123456"
#home_dir=/home/youtuapp/1.html-code/sky_eye_v2
gpu_dir=`pwd`
cd ..
home_dir=`pwd`
backend_dir=${home_dir}/web_backend
video_dir=${home_dir}/video_service/bin/camera_server
save_video_dir=${home_dir}/video_service/bin/video_server
task_dir=${home_dir}/task_manager
stream_dir=${home_dir}/stream_service/stream_server
skyeye_dir_72=${home_dir}/sky_eye_gpu_72/bin/face_retrieve_server
skyeye_dir_76=${home_dir}/sky_eye_gpu_76/bin/face_retrieve_server
root_dir=${home_dir}/root_server
feature_dir_72=${home_dir}/feature_service_72/bin
feature_dir_76=${home_dir}/feature_service_76/bin
route_dir=${home_dir}/route_service
property_dir=${home_dir}/property_service

info_success="serivce has been successful."
info_error="serivce installation error..."


#run web_backend
function run_web_backend(){
  cd $backend_dir
  num=`ps aux | grep "sky_eye.py" | grep -v grep | wc -l`
  echo 'sky_eye number:'$num
  if [ "${num}" -ne 2 ];then
      ps aux | grep "sky_eye.py" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup /home/youtuapp/anaconda2/bin/python sky_eye.py >nohup.out 2>&1 &
  fi
  
  num2=`ps aux | grep "sky_eye.py" | grep -v grep | wc -l`
  if [ "${num2}" -eq 2 ];then
      echo "web_backend ${info_success}"
  else
      echo "web_backend ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run video_server
function run_video_server(){
  cd $save_video_dir
  num=`ps aux | grep "video_server" | grep -v grep | wc -l`
  echo 'video_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "video_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./video_server >nohup.out 2>&1 &
  fi

  num2=`ps aux | grep "video_server" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "video_server ${info_success}"
  else
      echo "video_server ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run task_manager
function run_task_manager(){
  cd $task_dir
  num=`ps aux | grep "task_manager_server" | grep -v grep | wc -l`
  echo 'task_manager_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "task_manager_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./task_manager_server >nohup.out 2>&1 &
  fi

  num2=`ps aux | grep "task_manager_server" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "task_manager_server ${info_success}"
  else
      echo "task_manager_server ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run stream_service
function run_stream_service(){
  cd $stream_dir
  num=`ps aux | grep "srs" | grep -v grep | wc -l`
  process=`ps aux | grep "srs" | grep -v grep | awk '{print $2}'`
  echo 'stream srs number:'$num
  if [ "${num}" -ne 1 ];then
      #ps aux | grep "camera_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      kill -9 $process
      /etc/init.d/srs start
  fi
  num2=`ps aux | grep "srs" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "srs ${info_success}"
  else
      echo "srs ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4

  cd $stream_dir
  num=`ps aux | grep "server.py" | grep -v grep | wc -l`
  echo 'stream_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "server.py" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup /home/youtuapp/anaconda2/bin/python server.py >nohup.out 2>&1 &
  fi
  num2=`ps aux | grep "server.py" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "stream_service ${info_success}"
  else
      echo "stream_service ${info_error},is [/home/youtuapp/anaconda2/bin/server.py] script problem..."
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run sky_eye_gpu root server
function run_root_server(){
  cd $root_dir
  num=`ps aux | grep "root_server" | grep -v grep | wc -l`
  echo 'root_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "root_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./root_server --port=7201 &
  fi
  num2=`ps aux | grep "root_server" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "root_server ${info_success}"
  else
      echo "root_server ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4

}



#run face_retrieve_server_gpu
function run_face_retrieve_gpu_72(){
  cd $gpu_dir
  killall face_retrieve_server_gpu
  bash start_gpu_72.sh 
  #caution: -port=7202   Not --port ##############################################
  num2=`ps -ef | grep "face_retrieve_server_gpu -port=7202"| grep -v grep  | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "face_retrieve_server_gpu_72 ${info_success}"
  else
      echo "face_retrieve_server_gpu_72 ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run sky_eye_gpu
function run_face_retrieve_gpu_76(){
  cd $skyeye_dir_76
  num=`ps aux | grep "face_retrieve_server_gpu --port=7601" | grep -v grep | wc -l`
  echo 'face_retrieve_server_gpu number:'$num
  if [ "${num}" -ne 1 ];then
      #ps aux | grep "face_retrieve_server_gpu" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./face_retrieve_server_gpu --port=7601 --gpu_id=0 &
  fi

  num2=`ps aux | grep "face_retrieve_server_gpu --port=7601" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "face_retrieve_server_gpu_76 ${info_success}"
  else
      echo "face_retrieve_server_gpu_76 ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run feature_service
function run_feature_service_72(){
  cd $feature_dir_72
  num=`ps aux | grep "feature_server --port=7200" | grep -v grep | wc -l`
  echo 'feature_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "feature_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./feature_server --port=7200 --get_attr=1 &
  fi

  num2=`ps aux | grep "feature_server --port=7200" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "feature_service_72 ${info_success}"
  else
      echo "feature_service_72 ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run feature_service
function run_feature_service_76(){
  cd $feature_dir_76
  num=`ps aux | grep "feature_server --port=7600" | grep -v grep | wc -l`
  echo 'feature_server number:'$num
  if [ "${num}" -ne 1 ];then
      #ps aux | grep "feature_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./feature_server --port=7600  &
  fi

  num2=`ps aux | grep "feature_server --port=7600" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "feature_service_76 ${info_success}"
  else
      echo "feature_service_76 ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run camera_server
function run_camera_server(){
  cd $video_dir
  num=`ps aux | grep "camera_server" | grep -v grep | wc -l`
  echo 'camera_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "camera_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./camera_server >nohup.out 2>&1 &
  fi

  num2=`ps aux | grep "camera_server" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "camera_server ${info_success}"
  else
      echo "camera_server ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run route_service
function run_route_service(){
  cd $route_dir
  num=`ps aux | grep "route_server" | grep -v grep | wc -l`
  echo 'route_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "route_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./route_server >nohup.out 2>&1 &
  fi

  num2=`ps aux | grep "route_server" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "route_server ${info_success}"
  else
      echo "route_server ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}

#run property_server
function run_property_server(){
  cd $property_dir
  num=`ps aux | grep "property_server" | grep -v grep | wc -l`
  echo 'property_server number:'$num
  if [ "${num}" -ne 1 ];then
      ps aux | grep "property_server" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
      nohup ./property_server >nohup.out 2>&1 &
  fi

  num2=`ps aux | grep "property_server" | grep -v grep | wc -l`
  if [ "${num2}" -eq 1 ];then
      echo "property_server ${info_success}"
  else
      echo "property_server ${info_error}"
      exit
  fi
  echo -e "\n"
  sleep 4
}


run_web_backend
run_video_server
run_task_manager
run_stream_service
run_root_server
run_face_retrieve_gpu_72
run_face_retrieve_gpu_76
run_feature_service_72
run_feature_service_76
run_camera_server
run_route_service
run_property_server


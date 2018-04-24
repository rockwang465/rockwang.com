#!/bin/bash
passwd="123456"
home_dir=/home/youtuapp/code/sky_eye_v2
gpu_dir=${home_dir}/scripts
backend_dir=${home_dir}/web_backend
video_dir=${home_dir}/video_service/bin/camera_server
save_video_dir=${home_dir}/video_service/bin/video_server
task_dir=${home_dir}/task_manager
stream_dir=${home_dir}/stream_service/stream_server
skyeye_dir_76=${home_dir}/sky_eye_gpu_76/bin/face_retrieve_server
root_dir=${home_dir}/root_server
feature_dir_72=${home_dir}/feature_service_72/bin
feature_dir_76=${home_dir}/feature_service_76/bin
route_dir=${home_dir}/route_service
property_dir=${home_dir}/property_service
gpu_num=`nvidia-smi -L|wc -l`


#unset hik_path
hik_path=${home_dir}/video_service/thirdparty/hiksdk/lib:${home_dir}/video_service/thirdparty/hiksdk/lib/HCNetSDKCom
#echo ${hik_path}
#export hik_path
export LD_LIBRARY_PATH=${hik_path}:${LD_LIBRARY_PATH}
#echo ${LD_LIBRARY_PATH}
#run web_backend
cd $backend_dir
num=`ps aux | grep "sky_eye.py" | grep -v grep | wc -l`
echo 'sky_eye number:'$num
if [ "${num}" -eq 0 ];then
	nohup /home/youtuapp/anaconda2/bin/python sky_eye.py >nohup.out 2>&1 &
fi

#run video_server
cd $save_video_dir
num=`ps aux | grep "video_server" | grep -v grep | wc -l`
echo 'video_server number:'$num
if [ "${num}" -eq 0 ];then
	nohup ./video_server >nohup.out 2>&1 &
fi


#run task_manager
cd $task_dir
num=`ps aux | grep "task_manager_server" | grep -v grep | wc -l`
echo 'task_manager_server number:'$num
if [ "${num}" -eq 0 ];then
	nohup ./task_manager_server >nohup.out 2>&1 &
fi

#run stream_service
cd $stream_dir
num=`ps aux | grep "srs" | grep -v grep | wc -l`
process=`ps aux | grep "srs" | grep -v grep | awk '{print $2}'`
echo 'stream srs number:'$num
if [ "${num}" -eq 0 ];then
    /etc/init.d/srs start
fi
cd $stream_dir
num=`ps aux | grep "server.py" | grep -v grep | wc -l`
echo 'stream_server number:'$num
if [ "${num}" -eq 0 ];then
	nohup /home/youtuapp/anaconda2/bin/python server.py >nohup.out 2>&1 &
fi

#run sky_eye_gpu
cd $root_dir
num=`ps aux | grep "root_server" | grep -v grep | wc -l`
echo 'root_server number:'$num
if [ "${num}" -eq 0 ];then
	nohup ./root_server --port=7201 &
fi


#run sky_eye_gpu
num=`ps aux | grep "face_retrieve_server_gpu" | grep -v grep | grep -v nohup |wc -l`
echo 'face_retrieve_server_gpu number:'$num
total_num=$[${gpu_num}+1]
if [ "${num}" -ne ${total_num} ];then
        #ps aux | grep "face_retrieve_server_gpu" | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
	killall face_retrieve_server_gpu
        cd ${gpu_dir}	
	bash ${gpu_dir}/start_gpu_72.sh
        cd $skyeye_dir_76
	nohup ./face_retrieve_server_gpu --port=7601 --gpu_id=0&
fi

#run feature_service
cd $feature_dir_72
num=`ps aux | grep "feature_server --port=7200" | grep -v grep | wc -l`
echo 'feature_server number:'$num
if [ "${num}" -eq 0 ];then
	nohup ./feature_server --port=7200 --get_attr=1 &
fi

#run feature_service
cd $feature_dir_76
num=`ps aux | grep "feature_server --port=7600" | grep -v grep | wc -l`
echo 'feature_server number:'$num
if [ "${num}" -eq 0 ];then
        nohup ./feature_server --port=7600  &
fi

#run camera_server
cd $video_dir
num=`ps aux | grep "camera_server" | grep -v grep | grep -v nohup | wc -l`
echo 'camera_server number:'$num
if [ "${num}" -eq 0 ];then
	nohup ./camera_server >nohup.out 2>&1 &
fi

#run route_service
cd $route_dir
num=`ps aux | grep "route_server" | grep -v grep | wc -l`
echo 'route_server number:'$num
if [ "${num}" -eq 0 ];then
        nohup ./route_server >nohup.out 2>&1 &
fi

#run property_server
cd $property_dir
num=`ps aux | grep "property_server" | grep -v grep | wc -l`
echo 'property_server number:'$num
if [ "${num}" -eq 0 ];then
        nohup ./property_server >nohup.out 2>&1 &
fi





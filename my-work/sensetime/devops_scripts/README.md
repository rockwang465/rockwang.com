# 脚本文件介绍

### delete_old_release.sh
+ 删除超过15天以上的release文件
+ 定时任务: 
```
# rsync release to 5.150
0 2 * * * bash /data/backup/crontab_scripts/rsync_release_to_5.150.sh
```

### rsync_release_to_5.150.sh
+ 同步release文件到5.150机器上
```
# delete old release
0 5 * * * bash /data/backup/crontab_scripts/delete_old_release.sh
```
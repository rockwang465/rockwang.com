<template>
    <div class="tmpl">
        <!--  组件名navBar -->
        <nav-bar title="图片详情"></nav-bar>
        <!-- 组件名:navbar -->
        <!-- 使用：navbar-->
        <div class="photo-title">
            <p v-text="imgInfo.title"></p>
            <span>发起日期：{{imgInfo.add_time | convertDate}}</span>
            <span>{{imgInfo.click}}次浏览</span>
            <span>分类：民生经济</span>
        </div>
        <ul class="mui-table-view mui-grid-view mui-grid-9">
            <li v-for="(img,index) in imgs" :key="index"
                class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
                <!--<img :src="img.src">-->  <!--下面使用vue-preview，所以这一行就注释了-->
                <!--下面一行是老师的代码，旧版本vue-preview: height="100"是设置页面中每个小图片的高度，不是点开后的高度-->
                <!--<img class="preview-img" :src="img.src" height="100" @click="$preview.open(index, imgs)">-->
            </li>
            <!--注意: 这里是imgs，而按老师的来，则会在for中套imgs，会导致按倍数显示图片，解决方法就是离开上面的for循环即可-->
            <vue-preview :slides="imgs" @close="handleClose" style="float: left"></vue-preview>
        </ul>
        <div class="photo-desc">
            <p v-html="imgInfo.content"></p>
        </div>

        <!-- 评论内容开始 -->
        <div class="photo-bottom">
            <ul>
                <li class="photo-comment">
                    <div>
                        <span>提交评论</span>
                        <span><a @click="goback">返回</a></span>
                    </div>
                </li>
                <li class="txt-comment">
                    <textarea v-model="msg"></textarea>
                </li>
                <li>
                    <mt-button @click="sendComment" size="large" type="primary">发表评论按钮</mt-button>
                </li>
                <li class="photo-comment">
                    <div>
                        <span>评论列表</span>
                        <span>66条评论</span>
                    </div>
                </li>
            </ul>
            <ul class="comment-list">
                <!--3.3 循环comments -->
                <li v-for="(comment,index) in comments" :key="index">
                    {{comment.user_name}}：{{comment.content}} {{comment.add_time|convertDate}}
                </li>

            </ul>
            <!--3.6 加载更多时，使用loadByPage自增，然后追加到第一个列表内容中显示。-->
            <mt-button type="danger" size="large" plain @click="loadByPage">加载更多按钮</mt-button>
        </div>
        <!-- 评论内容结束 -->

        <!--改变颜色<mt-button type="default">default</mt-button>
        <mt-button type="primary">primary</mt-button>
        <mt-button type="danger">danger</mt-button> -->

        <!-- 改变大小
        <mt-button size="small">small</mt-button>
        <mt-button size="large">large</mt-button>
        <mt-button size="normal">normal</mt-button>
        幽灵按钮
        <mt-button type="danger" plain>plain</mt-button> -->


        <!-- 使用评论子组件 -->
        <!--<comment :cid="pid"></comment>-->
    </div>
</template>
<script>
    export default {
        data() {
            return {
                imgs: [],//存放缩略图
                imgInfo: {},//详情数据对象
                comments: [], //3.2 存放评论
                pageIndex: 1, //3.5 存放默认页码，正常从1开始
                pid: this.$route.params.id, //3.7 记录当前图片id，这种用法之前没说过，这样下面都可以调用这个pid了。
                msg: '',
            }
        },
        created() {
            //1:获取photoShare.vue中用户点击页面内容的路由id参数
            // let pid = this.$route.params.id;
            //2:发起请求2个
            //2.1:获取图片对应的详情
            // this.$ajax.get('getimageInfo/' + this.pid)  //老师的代码
            this.$ajax.get('photoDetail/getimageInfo/' + this.pid)  //Rock的代码，因为没有api，所以这样写的
                .then(res => {
                    //一个id对应一个详情对象,所以用message[0]
                    // this.imgInfo = res.data.message[0];  //老师的代码
                    this.imgInfo = res.data[0];  //Rock的代码，因为没有api，所以这样写的
                })
                .catch(err => {
                    console.log(err)
                });
            //2.2:获取缩略图(多个缩略图)的地址
            // this.$ajax.get('getthumimages/' + this.pid)  //老师的代码
            this.$ajax.get('photoDetail/getthumimages/' + this.pid)  //Rock的代码，因为没有api，所以这样写的
                .then(res => {
                    // this.imgs = res.data.message;  //老师的代码
                    this.imgs = res.data;  //Rock的代码，因为没有api，所以这样写的
                    //forEach（es6的语法）设置每个图片的高、宽
                    this.imgs.forEach((ele) => {
                        // 下面是设置每个图片点开后的大小(即缩略图)
                        ele.w = 400;
                        ele.h = 500; //设置缩略图显示的高
                        ele.msrc = ele.src;  //新版本vue-preview要加这一行，msrc为小图片，src为点击后的大图片。
                    })
                })
                .catch(err => {
                    console.log(err)
                });
            // 3.4 评论操作开始  先执行加载一下评论内容,且用上面的pid。
            this.loadFirst(this.pid);
            // 评论操作结束
        },
        methods: {
            handleClose() {
                console.log('close event')
            },
            // 3.1 评论列表中先加载评论内容
            // 打图文详细的页面,在地址栏可以看到，如: http://localhost:9999/#/photo/detail/37?pageindex=1中的37为cid。
            // loadFirst中的cid为的上面地址中的37.
            loadFirst() {
                // 请求获取打开的图文详细页面(cid)中的第一页评论(pageindex=1)。
                // this.$ajax.get('postcomment/' + cid + '?pageindex=1')  //老师的代码
                this.$ajax.get('photoDetail/postcomment/' + this.pid + 'pageindex1')  //Rock的代码，没有api的原因
                    .then(res => {
                        // this.comments = res.data.message;  //老师的代码
                        this.comments = res.data;  //Rock的代码，没有api的原因
                    })
                    .catch(err => {
                        console.log(err);
                    });
            },
            sendComment() {

            },
            // 3.8 loadByPage 加载更多，使用concat追加内容到第一页内容中
            loadByPage() {
                // this.$ajax.get('postcomment/' + cid + '?pageindex=1')  //老师的代码
                // ++pageIndex自增
                this.$ajax.get('photoDetail/postcomment/' + this.pid + 'pageindex' + (++this.pageIndex))  //Rock的代码，没有api的原因
                    .then(res => {
                        // this.comments = res.data.message;  //老师的代码
                        this.comments = this.comments.concat(res.data);  //Rock的代码，没有api的原因
                    })
                    .catch(err => {
                        console.log(err);
                    });
            },
            goback() {
                this.route.go(-1);
            }
        }
    }
</script>
<style scoped>
    li {
        list-style: none;
    }

    ul {
        margin: 0;
        padding: 0;
    }

    .photo-title {
        overflow: hidden;
    }

    .photo-title,
    .photo-desc {
        border-bottom: 1px solid rgba(0, 0, 0, 0.2);
        padding-bottom: 5px;
        margin-bottom: 5px;
        padding-left: 5px;
    }

    .photo-title p {
        color: #13c2f7;
        font-size: 20px;
        font-weight: bold;
    }

    .photo-title span {
        margin-right: 20px;
    }

    .mui-table-view.mui-grid-view.mui-grid-9 {
        background-color: white;
        border: 0;
    }

    .mui-table-view.mui-grid-view.mui-grid-9 li {
        border: 0;
    }

    .photo-desc p {
        font-size: 18px;
    }

    .mui-table-view-cell.mui-media.mui-col-xs-4.mui-col-sm-3 {
        padding: 2px 2px;
    }

    /*评论样式 开始*/
    .photo-comment > div span:nth-child(1) {
        float: left;
        font-weight: bold;
        margin-left: 5px;
    }

    .photo-comment > div span:nth-child(2) {
        float: right;
    }

    .photo-comment {
        height: 30px;
        border-bottom: 1px solid rgba(0, 0, 0, 0.4);
        line-height: 30px;
        margin-bottom: 5px;
    }

    .txt-comment {
        padding: 5px 5px;
    }

    .txt-comment textarea {
        margin-bottom: 5px;
    }

    .comment-list li {
        border-bottom: 1px solid rgba(0, 0, 0, 0.2);
        padding-bottom: 5px;
        margin-bottom: 5px;
        padding-left: 5px;
    }

    li {
        list-style: none;
    }

    ul {
        margin: 0;
        padding: 0;
    }

    /*评论样式 结束*/
</style>

<template>
    <div>
        <!--上部分-轮播图-->
        <mt-swipe :auto="4000">
            <!--由于rock没有api，所以轮播图引用data中的images_info,但老师是使用created中获取的imgs列表的。-->
            <mt-swipe-item v-for="(img,index) in images_info" :key="index">
                <a :href="img.url"><img :src="img.img" alt=""></a>
            </mt-swipe-item>
        </mt-swipe>

        <!--下部分-九宫格-->
        <div class="mui-content">
		        <ul class="mui-table-view mui-grid-view mui-grid-9">
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><router-link :to="{name: 'news.list'}">
		                    <span class="mui-icon mui-icon-home"></span>
		                    <div class="mui-media-body">新闻资讯</div></router-link></li>
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><router-link :to="{name: 'photo.share'}">
		                    <span class="mui-icon mui-icon-email"></span>
		                    <div class="mui-media-body">图文分享</div></router-link></li>
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-chatbubble"></span>
		                    <div class="mui-media-body">商品展示</div></a></li>
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-location"></span>
		                    <div class="mui-media-body">留言反馈</div></a></li>
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-search"></span>
		                    <div class="mui-media-body">搜索资讯</div></a></li>
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-phone"></span>
		                    <div class="mui-media-body">联系我们</div></a></li>
		        </ul>
		</div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                //由于我自己没有api接口，这里只能手写一个了
                images_info: [
                    {
                        url: "http://www.baidu.com",
                        img: "http://10.5.1.80:6001/images/1.jpg"
                    },
                    {
                        url: "http://www.jd.com",
                        img: "http://10.5.1.80:6001/images/2.png"
                    },
                    {
                        url: "http://www.taobao.com",
                        img: "http://10.5.1.80:6001/images/3.jpg"
                    }
                ],
                imgs: [], //1.3 老师的轮播图列表
            }
        },
        created() { // 用created是因为要在DOM生成前做这个操作
            // 1.1 main.js中设置Axios.defaults.baseURL="http://10.5.1.80:6001/images"
            // 1.2 由于rock没有api接口，所以下面不用了，用data中定义的。正确的做法是用下面ajax.get获取内容并赋值给imgs列表的。
            // this.$ajax.get("getlunbo")
            //     .then(res => {
            //         console.log(res.data.message);
            //         this.imgs = res.data.message;
            //     })
        }
    }

</script>

<style scoped>
    /*1.4 为了不影响其他的页面，所以加下scoped*/
    /*1.5 轮播图的样式*/
    /*设置div的最大高度*/
    .mint-swipe {
        max-height: 187px;
    }

    /*设置图片充满整个div*/
    .mint-swipe img {
        height: 100%;
        width: 100%;
    }

    /*1.6 九宫格的样式*/
    /*注意: 下面是连起来的，不是分开的 .mui-table-view.mui-grid-view.mui-grid-9*/
    /*背景色设为白色*/
    .mui-table-view.mui-grid-view.mui-grid-9 {
        background-color: white;  /*设置白色背景*/
        border: none; /*并取消ul的边框*/
        margin-top: 0; /*顶部margin取消*/
    }

    /*取消每个li的边框*/
    .mui-table-view.mui-grid-view.mui-grid-9 li {
        border: none;
    }

    /*清除九宫格文字图标，并设置高度*/
    .mui-icon-home:before,
    .mui-icon-email:before,
    .mui-icon-chatbubble:before,
    .mui-icon-location:before,
    .mui-icon-search:before,
    .mui-icon-phone:before{
        content: '';
    }

    /*设置每个九宫格中的背景图片*/
    .mui-icon-home{
        background-image: url("../../static/img/news.png");
    }
    .mui-icon-email{
        background-image: url("../../static/img/picShare.png");
    }
    .mui-icon-chatbubble{
        background-image: url("../../static/img/goodShow.png");
    }
    .mui-icon-location{
        background-image: url("../../static/img/feedback.png");
    }
    .mui-icon-search{
        background-image: url("../../static/img/search.png");
    }
    .mui-icon-phone{
        background-image: url("../../static/img/callme.png");
    }

    /*设置图片的高度*/
    .mui-icon{
        height: 50px;
        width: 50px;
        background-repeat: round;  /*缩放图片到对应的大小背景图中，这里非常合适的用法*/
    }
</style>
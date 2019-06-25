<template>
    <div class="tmpl">

        <nav-bar title="新闻列表"></nav-bar>   <!--3.1 父组件向子组件传值-->

    <!-- MUI 图文列表 -->
        <ul class="mui-table-view">
            <li v-for="news in newsListRock" :key="news.id" class="mui-table-view-cell mui-media">
                <router-link :to="{name:'news.detail',query:{id:news.id} }">  <!-- 此处用于新闻详情中 去哪里的定义，去news.detail组件，并添加了query查询字符串，给予newsDetail.vue中引用拼接url的-->
                    <img class="mui-media-object mui-pull-left" :src="news.img_url">
                    <div class="mui-media-body">
                        <span v-text="news.title"></span>  <!-- v-text、v-html为变量用法 -->
                        <div class="news-desc">
                            <p>点击数:{{news.click}}</p>
                            <p>发表时间:{{news.add_time | convertDate}}</p>   <!-- 2.1使用了过滤器 -->
                        </div>
                    </div>
                </router-link>
            </li>
        </ul>
    </div>
</template>
<script>
export default {
    data(){
        return {
            newsList:[],//新闻列表数据
            newsListRock: [ //由于rock没有api页面，所以这里还是自己写一个做调用吧
                {
                    id: 1,
                    title: "外交部就习近平主席出席二十国集团领导人第十四次峰会举行中外媒体吹风会",
                    click: 113,
                    img_url: "http://10.5.1.80:6001/images/1.jpg",
                    add_time: "2019-05-20T20:14:56.000Z"
                },
                {
                    id: 2,
                    title: "美媒:美官方确认美军一架'海神'无人机被伊朗击落",
                    click: 44,
                    img_url: "http://10.5.1.80:6001/images/2.png",
                    add_time: "2019-06-21T09:15:32.000Z"
                },
                {
                    id: 3,
                    title: "加拿大总理特鲁多飞抵美国,将同特朗普会面",
                    click: 22,
                    img_url: "http://10.5.1.80:6001/images/3.jpg",
                    add_time: "2019-07-22T11:34:46.000Z"
                },
                                {
                    id: 4,
                    title: "普京:俄罗斯从未与任何国家'争吵',也没有这样的想法",
                    click: 113,
                    img_url: "http://10.5.1.80:6001/images/1.jpg",
                    add_time: "2019-05-20T20:14:56.000Z"
                },
                {
                    id: 5,
                    title: "美媒:美官方确认美军一架'海神'无人机被伊朗击落",
                    click: 44,
                    img_url: "http://10.5.1.80:6001/images/2.png",
                    add_time: "2019-06-21T09:15:32.000Z"
                },
                {
                    id: 6,
                    title: "加拿大总理特鲁多飞抵美国,将同特朗普会面",
                    click: 22,
                    img_url: "http://10.5.1.80:6001/images/3.jpg",
                    add_time: "2019-07-22T11:34:46.000Z"
                },
                                {
                    id: 7,
                    title: "普京:俄罗斯从未与任何国家'争吵',也没有这样的想法",
                    click: 113,
                    img_url: "http://10.5.1.80:6001/images/1.jpg",
                    add_time: "2019-05-20T20:14:56.000Z"
                },
                {
                    id: 8,
                    title: "美媒:美官方确认美军一架'海神'无人机被伊朗击落",
                    click: 44,
                    img_url: "http://10.5.1.80:6001/images/2.png",
                    add_time: "2019-06-21T09:15:32.000Z"
                },
                {
                    id: 9,
                    title: "加拿大总理特鲁多飞抵美国,将同特朗普会面",
                    click: 22,
                    img_url: "http://10.5.1.80:6001/images/3.jpg",
                    add_time: "2019-07-22T11:34:46.000Z"
                },
                                {
                    id: 10,
                    title: "普京:俄罗斯从未与任何国家'争吵',也没有这样的想法",
                    click: 113,
                    img_url: "http://10.5.1.80:6001/images/1.jpg",
                    add_time: "2019-05-20T20:14:56.000Z"
                },
                {
                    id: 11,
                    title: "美媒:美官方确认美军一架'海神'无人机被伊朗击落",
                    click: 44,
                    img_url: "http://10.5.1.80:6001/images/2.png",
                    add_time: "2019-06-21T09:15:32.000Z"
                },
                {
                    id: 12,
                    title: "加拿大总理特鲁多飞抵美国,将同特朗普会面",
                    click: 22,
                    img_url: "http://10.5.1.80:6001/images/3.jpg",
                    add_time: "2019-07-22T11:34:46.000Z"
                },
            ]
        }
    },
    created(){
        //发起请求
        // 1.1 正确写法是下面的，由于Rock没有api请求的页面，只好上面定义了
        // this.$ajax.get('getnewslist')
        // .then(res=>{
        //     this.newsList = res.data.message;
        // })
        // .catch(err=>{
        //     console.log(err);
        // })
    }
}

</script>
<style scoped>
.mui-media-body p {
    color: #0bb0f5;
}

.news-desc p:nth-child(1) {
    float: left;
}

.news-desc p:nth-child(2) {
    float: right;
}
</style>

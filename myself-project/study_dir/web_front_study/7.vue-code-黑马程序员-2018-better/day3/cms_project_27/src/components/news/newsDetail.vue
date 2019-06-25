<template>
    <div class="tmpl">
        <nav-bar title="新闻详情"></nav-bar>  <!-- 头部 -->
        <div class="news-title">
            <p v-text="newsDetail.title"></p>
            <div>
                <span>{{newsDetail.click}}次点击</span>
                <span>分类:民生经济</span>
                <span>添加时间:{{newsDetail.add_time | convertDate}}</span>  <!-- 时间转换 -->
            </div>
        </div>
        <div class="news-content" v-html="newsDetail.content"></div>  <!-- 用v-html转换成网页，因为content为主体显示网站内容的文字部分 -->
    </div>
</template>
<script>
    export default {
        data() {
            return {
                // newsDetail: {}, //(这是老师写的)用{}或[]都不重要，因为以后都会被重新赋值，所以这里就只是为了象征性的表示其数据类型
                newsDetail: {  //这是Rock写的，没有api只能固定定义了
                    id: 43,
                    title: "外交部就习近平主席出席二十国集团领导人第十四次峰会举行中外媒体吹风会",
                    click: 3,
                    add_time: "2015-04-18T04:59:13.000Z",
                    content: "<h3 style='color: #835f8c'>G20峰会最新情况</h3><br/><span style='color: #9965e2'>新华社北京6月24日电（记者杨依军、潘洁）外交部24日举行中外媒体吹风会。</br>中方二十国集团（G20）事务协调人、外交部部长助理张军，财政部副部长邹加怡，商务部副部长兼国际贸易谈判副代表王受文，人民银行副行长陈雨露介绍国家主席习近平即将出席二十国集团领导人第十四次峰会有关情况，并回答记者提问。</br>邹加怡表示，今年G20财金渠道合作总体平稳，成果丰富。不仅聚焦短期经济形势和风险，更广泛着眼应对全球中长期挑战。6月，峰会前最后一次G20财长和央行行长会议成功举行，为大阪峰会做了充分的财金政策准备，达成以下成果：一是保持政策协调，维护全球经济金融稳定。G20成员重申将继续使用一切政策工具，通过加强对话和行动增强信心，实现世界经济强劲、可持续、平衡、包容增长。二是坚持发展导向，制定高质量基础设施投资原则。会议核准《G20高质量基础设施投资原则》。这份成果文件立足发展、支持发展，反映了中方一贯倡导的“共商共建共享”、“互联互通”等重要理念，是对杭州峰会有关共识的深化和延续，也为今后各方促进高质量基础设施投资提供了有用参考。三是推动改善全球失衡，促进经济平衡、可持续增长。会议通过了《G20全球失衡综述》，就全球经常账户失衡、全球失衡结构、如何解决贸易失衡等一些重要问题形成了较为客观的共识。不少国家对中国减少外部失衡取得的重要进展及对改善全球失衡的巨大贡献给予肯定。四是着眼中长期挑战，聚焦老龄化和全民健康覆盖。会议认为，G20各方应通过包容性的劳动力市场政策、改革养老体系、提高公共收支效率等措施积极应对老龄化问题，支持代际平等，并确保公共财政可持续性。</span>"
                }
            }
        },
        created() {
            //由于Rock这里没有api接口，下面就注释了，用上面data中定义的newsDetail内容，而老师的是每次点击一次，则获取一次，显示出对应请求的内容。
            //1:获取路由参数
            //let id = this.$route.query.id;  //获取newsList.vue中的query的id，给于下面一行代码做拼接使用。
            //2:拼接路由参数成为后台请求的URL(通过拼接 getnew/43 这种url来获取不同的信息)
            //this.$ajax.get('getnew/' + id)            // this.$ajax.get('getnew/' + id)
            // .then(res=>{
            //3:响应回来渲染页面
            // this.newsDetail = res.data.message[0];  //数组，就一个数据，所以直接取[0]
            // })
            // .catch(err=>{
            //     console.log(err);
            // })

        }
    }

</script>
<style scoped>
    .news-title p {
        color: #0a87f8;
        font-size: 20px;
        font-weight: bold;
    }

    .news-title span {
        margin-right: 30px;
    }

    .news-title {
        margin-top: 5px;
        border-bottom: 1px solid rgba(0, 0, 0, 0.2);
    }


    /*主体文章的左右距离*/
    .news-content {
        padding: 10px 5px;
    }
</style>

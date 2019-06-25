<template>
    <div class="tmpl">
        <nav-bar title="图文分享"></nav-bar>  <!-- 1.7 标题栏加上 -->
        <!-- 引入返回导航 -->
        <div class="photo-header">
            <ul>
                <li v-for="category in categorys" :key="category.id">
                    <a href="javascript:;" @click="loadImg(category.id)">{{category.title}}</a>
                </li>

            </ul>
        </div>
        <div class="photo-list">
            <ul>
                <li v-for="img in imgs" :key="img.id">
                    <!--<router-link :to="{name:'photo.detail',params:{id:img.id} }">-->
                    <a>
                        <img :src="img.img_url">
                        <!-- 懒加载 -->
                        <!--<img v-lazy="img.img_url">-->
                        <p>
                            <span v-text="img.title"></span>
                            <br>
                            <span v-text="img.zhaiyao"></span>
                        </p>
                    </a>
                    <!--</router-link>-->
                </li>
            </ul>
        </div>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                categorys: [],  //1.2 分类
                imgs: [],  //图片数据
            }
        },
        created() {
            //1.1 发起请求获取导航栏数据
            // this.$ajax.get('getimgcategory')   //这是老师写的
            this.$ajax.get('photoShare/getimgcategory')
                .then(res => {
                    // this.categorys = res.data.message;  //这是老师写的
                    this.categorys = res.data;
                    // this.categorys = JSON.parse(res.data);  //正常不用JSON.parse去转换字符串为json格式的

                    //1.3 将 全部 添加到数组的第一个位置，用unshift
                    this.categorys.unshift({
                        id: 0,
                        title: '全部'
                    });
                })
                .catch(err => {
                    console.log(err);
                });

            //1.6 当页面加载默认传递0，因为0代表全部，所以刚加载的时候就是加载全部内容
            this.loadImg(0);  //该代码替换了下面的请求的代码，做了函数封装

            //1.4 将0作为参数，获取全部图片数据--下面1.5替代了此处的操作，所以此处全部注释
            // 注意: api中说的getimages/0 应该是后端做好了，把所有的数据都放在0中，所以当你请求0的时候，能获取所有信息(图片+其他内容)
            // this.$ajax.get('getimages/' + 0)
            // .then(res=>{
            //     this.imgs = res.data.message;
            // })
            // .catch(err=>{
            //     console.log(err);
            // })
        },
        methods: {
            //1.5 使用loadImg方法，替代上面1.4的操作，所以上面1.4全部注释
            loadImg(id) {
                // this.$ajax.get('getimages/' + id)
                this.$ajax.get('photoShare/getimages/' + id)
                    .then(res => {
                        this.imgs = res.data;
                        // this.imgs = JSON.parse(res.data);
                    })
                    .catch(err => {
                        console.log(err);
                    })
            }
        }
    }


</script>
<style>
    .photo-header li {
        list-style: none;
        display: inline-block;
        margin-left: 10px;
        height: 30px;
    }

    .photo-header ul {
        /*强制不换行*/
        white-space: nowrap;
        overflow-x: auto;
        padding-left: 0px;
        margin: 5px;
    }


    /*下面的图片*/

    .photo-list li {
        list-style: none;
        position: relative;
    }

    .photo-list li img {
        width: 100%;
        height: 230px;
        vertical-align: top;
    }

    .photo-list ul {
        padding-left: 0px;
        margin: 0;
    }

    .photo-list p {
        position: absolute;
        bottom: 0px;
        color: white;
        background-color: rgba(0, 0, 0, 0.3);
        margin-bottom: 0px;
    }

    .photo-list p span:nth-child(1) {
        font-weight: bold;
        font-size: 16px;
    }

    /*图片懒加载的样式*/
    image[lazy=loading] {
        width: 40px;
        height: 300px;
        margin: auto;
    }

    a:hover{
        background-color: skyblue;
    }
</style>

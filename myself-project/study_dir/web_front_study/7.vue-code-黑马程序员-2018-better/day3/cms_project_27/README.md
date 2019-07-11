# vue实战项目编写记录

+ 后期文档保存到rockwang465 github的web代码中
+ 路径 `E:\3.常用代码目录\rockwang.com\myself-project\study_dir\web_front_study\7.vue-code-黑马程序员-2018-better\day3\cms_project_27\README.md`

## 1 目录结构搭建及各基础文件编写

+ 实现效果图

![image](A7F061B5B9D444BA8A6FCEE2771A1A8B)

### 1.1 目录结构 -- 待后期补充
```
cms_project_27
    |
    |--src
        |--components
            |--home
                |--home.vue  #主页
            |--search
                |--search.vue  #搜索页面
            |--shopcart
                |--shopcart.vue  #购物车页面
            |--vip
                |--vip.vue  #会员页面
            |--news
                |--newsList.vue  #新闻列表页面
        |--static
            |--css
                |--global.css  #全局样式
            |--img
                |-- *.jpg/png  #所有图片存放位置
            |--vendor
                |--mui  #mui插件目录
                    |--dist
                    |--examples
                        |--hellow-mui  #hello-mui插件
                        |--login  #登录插件
                    |--... ...
        |--index.html
        |--app.vue
        |--main.js
    |--package.json
    |--webpack.config.js
```

### 1.2 理解当前项目所需安装的模块及安装
![image](72A3EA78056241E2A4169A963608D5F4)

```
cmd运行：
npm i mint-ui vue-preview axios vue-router moment vue -S;
npm i webpack-dev-server -D

开发依赖+css+js：
npm i webpack html-webpack-plugin css-loader style-loader less less-loader autoprefixer-loader(自动补前缀)  (js处理:)babel-loader babel-core babel-preset-es2015 babel-plugin-transform-runtime (文件:)url-loader file-loader vue-loader vue-template-compiler -D
下面不带注释的命令:
npm i webpack html-webpack-plugin css-loader style-loader less less-loader autoprefixer-loader babel-loader babel-core babel-preset-es2015 babel-plugin-transform-runtime url-loader file-loader vue-loader vue-template-compiler -D
```

+ 打开cmd进入对应目录，执行上面命令，完成依赖的建立。

### 1.3 package.json文件
```javaScript
{
  "name": "cms_project_27",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "dev": ".\\node_modules\\.bin\\webpack-dev-server --inline --hot --open --port 9999",
    "dev-bak": "E:\\web_code_node_module\\node_modules_27\\.bin\\webpack-dev-server --inline --hot --open --port 9999",
    "build": "webpack"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "autoprefixer-loader": "^3.2.0",
    "babel-core": "^6.26.3",
    "babel-loader": "^7.1.5",
    "babel-plugin-transform-runtime": "^6.23.0",
    "babel-preset-es2015": "^6.24.1",
    "css-loader": "^0.28.11",
    "file-loader": "^0.11.2",
    "html-webpack-plugin": "^2.30.1",
    "less": "^2.7.3",
    "less-loader": "^4.1.0",
    "style-loader": "^0.18.2",
    "url-loader": "^0.5.9",
    "vue-loader": "^13.7.3",
    "vue-template-compiler": "^2.6.10",
    "webpack": "^3.12.0",
    "webpack-dev-server": "^2.11.5"
  },
  "dependencies": {
    "axios": "^0.16.2",
    "mint-ui": "^2.2.13",
    "moment": "^2.24.0",
    "vue": "^2.6.10",
    "vue-preview": "^1.1.3",
    "vue-router": "^2.8.1"
  }
}
```

### 1.4 webpack.config.js文件
```javaScript
'use strict';
const path = require('path');  //用path可以拿绝对路径
const htmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {

    //入口
    entry: {
        main: './src/main.js'
    },
    output: {
        //所有产出资源路径
        path: path.join(__dirname, 'dist'),
        filename: 'build.js'
    },
    module: {
        loaders: [{  //loaders为数组
                test: /\.css$/,
                // 从右往做写
                loader: 'style-loader!css-loader!autoprefixer-loader'
            },
            {
                test: /\.less$/,
                loader: 'style-loader!css-loader!autoprefixer-loader!less-loader'
            },
            {
                test: /\.(jpg|png|svg|ttf|woff|woff2|gif)$/,
                loader: 'url-loader',
                options: {
                    limit: 4096, //4096字节以上生成文件，4096以下自己生成文件base64
                    name: '[name].[ext]'  //这里的[name]是url-loader的一种方法，可以去https://webpack.js.org上搜url-loader，看下没有找到。
                                          //这里搜file-loader里面有[name] [ext]这些用法，他们都是通用的用法。
                }
            },
            {
                test: /\.js$/,
                loader: 'babel-loader',
                exclude: /node_modules/,
                options: {
                    presets: ['es2015'], //关键字
                    plugins: ['transform-runtime'], //函数
                }
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            }
        ]
    },
    // 插件
    plugins: [
        new htmlWebpackPlugin({
            template: './src/index.html'
        })
    ]
}
```

### 1.5 src目录
#### 1.5.1 index.html文件
+ meta:vp
  - 直接生成`<meta name="viewport" content="width=device-width, initial-scale=1.0">`
  - 但我的pycharm显示的内容少了，应该是:
  - `<meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">`

+ 最终文件为
```
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>传智27期信息管理系统</title>
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
</head>
<body>
    <div id="app"></div>

</body>
</html>
```

#### 1.5.2 main.js文件

```javaScript
'use strict';  // "严格模式"是一种在JavaScript代码运行时自动实行更严格解析和错误处理的方法。这种模式使得Javascript在更严格的条件下运行。

//1.3 引入第三方包 开始
import Vue from 'vue';
//1.3.1 VueRouter:引入路由对象
import VueRouter from 'vue-router';
//VueRouter
//VueRouter:安装插件
Vue.use(VueRouter);

//1.3.2 Mint:引入mint-ui
import Mint from 'mint-ui';
//Mint:引入css
import 'mint-ui/lib/style.css';
//Mint:安装插件
Vue.use(Mint);

//1.3.3 Axios:引入axios
import Axios from 'axios';
//挂载原型
Vue.prototype.$ajax = Axios;
// Axios.prototype.$axios = Axios;  //叫$axios最好，老师习惯，叫$ajax的，所以用的上面的

//引入第三方包 结束



//1.1 引入自己的vue文件 开始
import App from './app.vue';
import Home from './components/home.vue';
//引入自己的vue文件 结束


//1.4 VueRouter: 创建对象并配置路由规则
let router = new VueRouter({
    routes: [
        //VueRouter: 配置路由规则
        //默认根目录配置重定向到home
        {path: '/', redirect: {name: 'home'}},
        {name: 'Home', path: '/home', component: Home}
    ]
});


//1.2 创建vue实例
new Vue({
    el: '#app',
    router, //如果这里不加，则会报错:Vue warn]: Error in render: "TypeError: Cannot read property 'matched' of undefined"
    render: c => c(App)
});
```

#### 1.5.3 vue.app文件
+ mint-ui的调用方法
  - 进入网站:  例如需要button，找到左侧button，复制对应的内容如: `<mt-button type="primary">测试mint-ui按钮</mt-button>`

```
<template>
    <div>
        <mt-button type="primary">测试mint-ui按钮</mt-button>
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        data() {
            return {}
        },
        created() {
            this.$ajax.get('http://10.5.1.80:6001/images/4.txt')
                .then(res => {
                    console.log(res)
                })
        }
    }
</script>

<style scoped>

</style>
```

### 1.6 自己搭建个nginx静态图片服务器
+ 自己搭建个nginx静态图片服务器，用于axios的api请求使用
```
# docker pull nginx  #镜像拉取
# docker load -i nginx.tar  #如果下不了镜像，就导入镜像
# docker run -dit --name web-api-nginx -p 6001:80 -v /root/images:/usr/share/nginx/html/images -v /root/default.conf:/etc/nginx/conf.d/default.conf -v /root/nginx.conf:/etc/nginx/nginx.conf nginx
```

+ 这里docker起的nginx.conf需要加一个跨域问题，否则页面报错Access-Control-Allow-Origin
```
# cat /etc/nginx/nginx.conf  
user  nginx;
worker_processes  1;

error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       /etc/nginx/mime.types;
    
    # rock add ==>这里加3行内容即可
    add_header 'Access-Control-Allow-Origin' '*';
    add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
    add_header 'Access-Control-Allow-Credentials' 'true';

    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;
}


```
+ 这里default.conf中的server中有location，所以在这里配置

```
# cat /etc/nginx/conf.d/default.conf
server {
    listen       80;
    server_name  localhost;

    #charset koi8-r;
    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }
    
    ## rock add images dir ==>加这里5行内容即可
    location /images/ {
        root  /usr/share/nginx/html;
        autoindex on;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}

    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    #location ~ \.php$ {
    #    root           html;
    #    fastcgi_pass   127.0.0.1:9000;
    #    fastcgi_index  index.php;
    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
    #    include        fastcgi_params;
    #}

    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    #location ~ /\.ht {
    #    deny  all;
    #}
}
```
+ 注意本地提前准备好以下几个文件
```
/root/images  目录，且里面要放几张图片
/root/default.conf
/root/nginx.conf
```

+ 测试效果
这里是公司vmware起的，ip为10.5.1.80
网页打开: 10.5.1.80:6001/images/1.jpg
![image](69CAA1FB739A44ACAA864AD0654AED9D)
![image](FD11F48DCA2F4851950305C1E21E243A)

+ 确认代码页面正常
![image](DE5F12806D4F44F4B6EE4A3062D84CE3)
至此，前期准备工作全部完成



## 2 头部和底部制作

### 2.1 分析
#### 2.1.1 实现效果
![image](9640237CE69E429BA9439643C3BA0939)
![image](4D302F16A7994FBF8B06A079E801DA24)

#### 2.1.2 分析图片
+ 这里发现头部和底部是固定的，中间是变化的
![image](8453A62C0EEE4F8D85808D7D705C0F5A)

### 2.2 头部制作
+ 进入mint-ui的头部样式: mint-ui.github.io/docs/#/zh-cn/header
  - 由于引入部分我们是全部引入的mint-ui,所以这里不用做引入操作了。
  - 拷贝代码: `<mt-header fixed title="固定在顶部"></mt-header>`
  - 注意: fixed 应该是固定头部，我们这里不需要固定头部，所以fixed 要删掉。 `<mt-header title="传智27期信息管理系统"></mt-header>`
  - 将代码放入 app.vue 的头部内容中

### 2.3 底部制作
+ 进入hellomui的底部样式: www.dcloud.io/hellomui
  - 找到tab bar(选项卡) --> 底部选项卡-div模式 --> 此时地址栏地址为: http://www.dcloud.io/hellomui/examples/tabbar.html
  - 所以在本地windows保存的hellomui中找到tabbar.html
  - 老师的项目中对应路径为: cms_project_27\src\static\vendor\mui\examples\hello-mui\examples\tabbar.html
+ 打开 src\static\vendor\mui\examples\hello-mui\examples\tabbar.html 中有这样一些内容: 
```
  		<nav class="mui-bar mui-bar-tab">
			<a class="mui-tab-item mui-active" href="#tabbar">
				<span class="mui-icon mui-icon-home"></span>
				<span class="mui-tab-label">首页</span>
			</a>
			<a class="mui-tab-item" href="#tabbar-with-chat">
				<span class="mui-icon mui-icon-email"><span class="mui-badge">9</span></span>
				<span class="mui-tab-label">消息</span>
			</a>
			<a class="mui-tab-item" href="#tabbar-with-contact">
				<span class="mui-icon mui-icon-contact"></span>
				<span class="mui-tab-label">通讯录</span>
			</a>
			<a class="mui-tab-item" href="#tabbar-with-map">
				<span class="mui-icon mui-icon-gear"></span>
				<span class="mui-tab-label">设置</span>
			</a>
		</nav>
```
  - 将代码拷贝到 app.vue 的底部内容中
  - 但是此时还没有样式，需要引入样式

  ![image](0F7065A51884453784A0F05FD48BB053)
+ 在main.js中引入mui的css样式
  - 注意: 由于我们这里static是没有对应的hellomui的example的，所以需要先copy老师的src\static\下所有目录到自己的src\static\下
  - 其实这里所需的所有mui内容都在 src\static\vendor\mui\ 下了，这个要知道的
  - main.js中引入样式: `import './static/vendor/mui/dist/css/mui.css'`
  - 此时有效果了

  ![image](BA1BB66D301C4FFFA11E5C1C74521797)
+ 优化底部内容
  - 修改底部的文字为: 首页、会员、购物车、查找
  - 将会员位置的红色小球移到购物车上显示，就是移动刚刚贴的代码中的这部分内容: `<span class="mui-badge">9</span>`
  ![image](947AE33EA73A4EB0AA9195DC1F7D41C8)

### 2.4 修改底部的图标
#### 2.4.1 查看hellomui中的icon图标
![image](97D33E74CF0A40419A90D3A1032F0C2D)
![image](626397ECF53749B083A435F3C7E3C5DE)
![image](1FF99E1B198F494F88EE07D17AA15B12)
+ 这里发现hellomui中的icon图标很少，根本不够用的。所以需要使用外部的icon图标

## 3 引用外部自定义icon图标
### 3.1 阿里巴巴矢量图标

#### 3.2 打开 https://www.iconfont.cn/
  + 选择所需的icon图标( 首页、会员、购物车、 查找 ) 并下载到本地。
  + 解压后，可以看到有个 demo_index.html 文件，打开里面有对应的引用css样式
```
    <link rel="stylesheet" href="demo.css">
    <link rel="stylesheet" href="iconfont.css">
```
  + 打开 iconfont.css
```
    #里面也是引用的ttf文件
    url('iconfont.ttf?t=1560928500325') format('truetype')
    
    #这里是对应的icon的标识代码，后面要的就是这些东西
    .icon-ai-home:before {
      content: "\e60d";
    }
    
    .icon-home:before {
      content: "\e626";
    }
    
    .icon-huiyuan:before {
      content: "\e654";
    }
    
    .icon-huiyuan1:before {
      content: "\e65d";
    }
    
    .icon-huiyuan3:before {
      content: "\e636";
    }
    
    .icon-huiyuan4:before {
      content: "\e7af";
    }
    
    .icon-gouwucheman:before {
      content: "\f0178";
    }
    
    .icon-gouwuchekong:before {
      content: "\f0179";
    }
    
    .icon-sousuo:before {
      content: "\e615";
    }
    
    .icon-sousuo1:before {
      content: "\e603";
    }
    
    .icon-wodehuiyuanqia:before {
      content: "\e619";
    }
    
    .icon-my:before {
      content: "\e660";
    }
    
    .icon-wodedangxuan:before {
      content: "\e607";
    }
    
```
#### 3.3 将下载的自定义icon图标文件及样式代码放入到项目中
  + 进入自己项目的static\vendor\mui\dist\css\mui.css
    - 你会发现，这里面有引用的对应图标代码。
    - 例如搜索: 前面F12展示图片中引用的icon代码content: '\E502' 这个E502关键字
      ![image](2F6A985FA9C3409FAD09AD0137208F90)
  + 将iconfont.ttf文件在css\mui.css中引入
    - css\mui.css中搜索: ttf，在mui.ttf引入的下面引入自定义的字体图标文件
```
    @font-face {
        font-family: Muiicons;
        font-weight: normal;
        font-style: normal;
        src: url('../fonts/mui.ttf') format('truetype');
    }
    /*引入自定义的字体图标*/
    @font-face {
        font-family: Muiicons;
        font-weight: normal;
        font-style: normal;
        src: url('../fonts/iconfont.ttf') format('truetype');
    }
```
  + 将下载下来的iconfont.ttf文件，放到上面定义的../fonts/中
    ![image](BF23A14DAD2D40EF9932E8FF08CCC842)
  + 将下载下来的iconfont.css文件中的icon样式文件内容粘贴到 css\mui.css中
```
    /*引入自定义的样式*/
    .icon-gouwucheman:before {
      content: "\f0178";
    }
    
    .icon-gouwuchekong:before {
      content: "\f0179";
    }
    
    .icon-51:before {
      content: "\e633";
    }
    ... ...
  
```
  ![image](10847580757D46B791684B406A0F6516)
#### 3.4 使用自定义的icon图标
+ 在app.vue中写入代码
  - 打开app.vue
  - 对比src\static\vendor\mui\dist\css\mui.css 中自定义的icon图标代码
  - 填入自定义的icon图标样式代码
  - 修改代码前后对比
  ![image](30FA802F181A4B62A4A0FA5E991AE11C)
  ![image](5D6BAD445FFB464FA06285A3D3AC5791)
  ![image](A8A9029FA9994B76A8967D13CC68CB02)


## 4 主体-上部分-轮播图制作
### 4.1 轮播图部分理解
+ 理解
  ![image](0B51C268C8114DFDAD24EB9BC327F534)
  - 从上图可以看出，这里轮播图部分，应该是属于home.vue的router-view部分内容
  - 这里使用mint-ui来做
### 4.2 mint-ui的引用
+ swipe轮播图代码引用
  - 对应地址: http://mint-ui.github.io/docs/#/zh-cn/swipe
  - 拷贝基本用法代码:
```
    <mt-swipe :auto="4000">
      <mt-swipe-item>1</mt-swipe-item>
      <mt-swipe-item>2</mt-swipe-item>
      <mt-swipe-item>3</mt-swipe-item>
    </mt-swipe>
```
+ 通过遍历数字，输出图片到轮播图代码中
  - 先拿到轮播图的api内容看下--老师的图
    ![image](537CF073F0904439BB660610C0C69EB8)
  - 这里要拿到轮播图的 .data.img ，而.data.url为点击图片的跳转链接地址
  - 但由于这里老师是使用api的，而我还不会flask做api，所以我的home.vue中是使用data中定义的一个images_info列表的。

+ 正确应该在home.vue中script中写created，用ajax请求轮播图，并赋值给imgs列表。
+ 然后在home.vue的template中轮播图代码中用 v-for 填入获取的imgs值。 
+ 设置轮播图的大小
  - 从图中可以看出，控制轮播图大小的div名称为mint-swipe
  ![image](2C9B97FB233543619AF3386B0E0DC3C9)
  - 所以在style中设置大小
+ 以下为home.vue中的代码
```
<template>
    <div>
        <mt-swipe :auto="4000">
            <!--1.2 通过for循环将对应的值放入轮播图代码中-->
            <mt-swipe-item v-for="(img,index) in imgs" :key="index">
                <a :href="img.url"><img :src="img.img" alt=""></a>
            </mt-swipe-item>
        </mt-swipe>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                imgs: [], //1.3 老师的轮播图列表先设为空
            }
        },
        created() {  // 用created是因为要在DOM生成前做这个操作
            // 1.1 main.js中设置Axios.defaults.baseURL="http://10.5.1.80:6001/images"
            this.$ajax.get("getlunbo")
                 .then(res => {
                    console.log(res.data.message);  //data.message为老师的api获取的图片列表信息
                    this.imgs = res.data.message;
                })
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
</style>
```
+ 轮播图的效果图

 ![image](A8C77016CB804FE4A525E579DC97335F)


## 5 主体-下部分-九宫格制作
### 5.1 hellomui的九宫格样式
  + 地址: http://www.dcloud.io/hellomui/examples/grid-default.html
  + 对应项目下的hellomui路径: \cms_project_27\src\static\vendor\mui\examples\hello-mui\examples\grid-default.html
  + 进入grid-default.html，并拷贝九宫格的代码
### 5.2 插入九宫格代码到home.vue中
  + 由于这里只需要6个格子，所以删掉后面3段九宫格代码
  + 且每个格子的名字分别改为: 新闻资讯、图文分享、商品展示、留言反馈、搜索资讯、联系我们
  + 且图文分享中有个小红球也不要(删除这段代码 `<span class="mui-badge">5</span>`)
  + 在template的轮播图下面贴上九宫格代码
```
    <div class="mui-content">
		        <ul class="mui-table-view mui-grid-view mui-grid-9">
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-home"></span>
		                    <div class="mui-media-body">新闻资讯</div></a></li>
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-email"></span>
		                    <div class="mui-media-body">图文分享</div></a></li>
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
```
### 5.3 九宫格后的背景色改为白色
+ 修改九宫格的背景颜色
  - 为了省事，这里直接在src\static\css中创建一个global.css的全局样式文件，并写入代码
```
        /*全局css样式文件*/
        body {
            background-color: white;
        }
```
  - main.js中引入global.css文件
```
    //1.3.5 引入全局样式(自定义的)
    import './static/css/global.css'
```
  - 此时背景颜色已经变白了，但是格子的线条和背景还是灰色
  
![image](EE94F6D5B85A410D8BAC4D776D7CA3F8)
+ 将九宫格的格子线及背景颜色取消
  - 找到对应的ul总的盒子
    ![image](CFB648CC92E84E31ABC655F440AE8E55)
  - home.vue的style中修改样式
```
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
```
  - 注意: 上面是连起来的，不是分开的 .mui-table-view.mui-grid-view.mui-grid-9

### 5.4 九宫格的每个图片修改
+ 先确认九宫格中每个格子的名字已经修改好
+ 然后清除九宫格原有的字体图标
  - 找到对应的文字图标名称，保存下来，下一步在home.vue中修改
  ![image](854070C26F234A749FDD65B69F19AB0C) 
  - home.vue的style中清除图标，设置高度，并放入背景图
```
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

```
  - 总结难点: 
    + :before的content为字体图标的设置。
    + 背景图缩放用background-repeat: round;
  - 效果图

![image](C3521E73CE3C4F8EA977707745FDD176)

## 6 底部-各按钮跳转及变色
### 6.1 理解
![image](982B68ECD71F48EC8682320E57A52E0F)

### 6.2 目录及文件创建
+ 创建 搜索、购物车、会员对应的目录及vue文件
```
        |--components
            |--home
                |--home.vue  #主页
            |--search
                |--search.vue  #搜索页面
            |--shopcart
                |--shopcart.vue  #购物车页面
            |--vip
                |--vip.vue  #会员页面
```
### 6.3 app.vue中修改对应按钮(去哪里)
+ app.vue中的底部内容的代码部分修改:
  - 将a标签替换为router-link
  - 将href改为to="{name:'home/search/shopcart/vip'}"
+ 修改前 

![image](687E887D1A1049C4BB161F2CF7AAA293)
+ 修改后

![image](EC3816EE99AE4678BCAE0B12FF2B9640)

### 6.4 main.js中定义(搜索、会员、购物车)导航内容
+ 引入组件文件
```
//1.1.2 引入底部按钮对应的组件文件
import Search from './components/search/search.vue';
import Vip from './components/vip/vip.vue';
import Shopcart from './components/shopcart/shopcart.vue';
```
+ 定义路由规则
```
let router = new VueRouter({
    routes: [
        ... ...
        // 1.4.2 添加底部按钮的规则
        {name: 'search', path: '/search', component: Search},
        {name: 'vip', path: '/vip', component: Vip},
        {name: 'shopcart', path: '/shopcart', component: Shopcart}
    ]
```
+ 此时点击就有效果了

![image](3FE2A5A6085E414FA5B9CC964722D907)

### 6.5 点击底部按钮变色
+ 解释router-link-active
  - 当你点击任意按钮的时候，router-link上会生成一个 router-link-active的类样式。
  - 而默认app.vue的主页上，是有个mui-active，这个样式是让当前图标变为设定的蓝色。
  - 或者应该说: 当前按钮的锚点值匹配上了当前的路由时，就会变蓝色。
  - 所以: 我们只需要在点某个按钮时，临时加上mui-active这个样式，就会让这个按钮图标变蓝了。
+ vue.js官网中对router-link中router-link-active的解释
![image](8EFD8127466D45579158FE7ED9BE1318)
![image](39E1C74335464A1AA042BC71868D82D7)

  - 官网说: 默认值可以通过路由的构造选项 linkActiveClass 来全局配置。
  - 这里，构造：表示构造函数，即new一个函数。选项：表示传入一个对象。
  - 理解为: 默认值可以通过路由(vue-router)，在构造函数中传入一个对象，这个对象要有linkActiveClass这个属性。
+ 所以，在main.js中的new VueRouter中加入定义
```
  let router = new VueRouter({
    linkActiveClass: 'mui-active',  #这里加入一行
    routes:[
        ... ...
    ]
```
+ 并删除app.vue底部内容中"首页"默认定义的 mui-active
![image](D914E0D7D97B41C48B189554D719CFB1)
+ 因为我默认是home页面，所以会自动激活给home页面添加mui-active，点击其他的时候则仅点击的按钮变蓝色。


## 7 主体-九宫格-新闻资讯页面制作
### 7.1 思路理解
![image](A2914E964882422294A148DF0DB155A6)

### 7.2 九宫格第一个新闻资讯图标定义(去哪里)
+ 编辑home.vue
```
		        <ul class="mui-table-view mui-grid-view mui-grid-9">
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
		                    <span class="mui-icon mui-icon-home"></span>
		                    <div class="mui-media-body">新闻资讯</div></a></li>
```
+ 修改: a链接改为router-link, href改为:to="{name: 'news.list'}"
```
		        <ul class="mui-table-view mui-grid-view mui-grid-9">
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><router-link :to="{name: 'news.list'}">
		                    <span class="mui-icon mui-icon-home"></span>
		                    <div class="mui-media-body">新闻资讯</div></router-link></li>
```
### 7.3 main.js中引入文件及路由组件定义(导航)
```
//1.1.3 引入新闻列表文件
import NewsList from './components/news/newsList.vue'

//1.4 VueRouter: 创建对象并配置路由规则
let router = new VueRouter({
    ... ...
    routes: [
        //VueRouter: 配置路由规则
        //name为路由的名字，path为url上显示的地址，component为上面引入文件的名字。
        ... ...
        // 1.4.3 添加新闻列表的规则
        {name: 'news.list', path: '/news/list', component: NewsList},
    ]
});
```

### 7.4 src\components\下创建news目录及newList.vue文件
### 7.5 拷贝newList.vue现成制作好的文件，然后进行部分调整
#### 7.5.1 老师的新闻列表源码来自hellomui的图文列表 
+ 地址: http://www.dcloud.io/hellomui/examples/media-list.html

![image](FF1657F6ED444E129B1AC9D29DD98A88)

#### 7.5.2 拷贝newList.vue代码
```
<template>
    <div class="tmpl">
    
        <nav-bar title="新闻列表"></nav-bar>   #这个后面 7.8.3 中要做新闻列表的头部，通过父子组件传值的。这里先粘贴过来了

    <!-- MUI 图文列表 -->
        <ul class="mui-table-view">
            <li v-for="news in newsList" :key="news.id" class="mui-table-view-cell mui-media">
                <router-link :to="{name:'news.detail',query:{id:news.id} }">   <!-- (下面8.2中的代码，这里直接拷贝过来的，本身应该为a链接的)此处用于新闻详情中 去哪里的定义，去news.detail组件，并添加了query查询字符串，给予newsDetail.vue中引用拼接url的-->
                    <img class="mui-media-object mui-pull-left" :src="news.img_url">
                    <div class="mui-media-body">
                        <span v-text="news.title"></span>
                        <div class="news-desc">
                            <p>点击数:{{news.click}}</p>
                            <p>发表时间:{{news.add_time | convertDate}}</p>  #这个是时间过滤器，下面就会讲到
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
            newsList:[], //新闻列表数据
        }
    },
    created(){
        //发起请求
        this.$ajax.get('getnewslist')
        .then(res=>{
            this.newsList = res.data.message;  //赋值给newsList
        })
        .catch(err=>{
            console.log(err);
        })
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

```

### 7.6 底部内容显示被挡住问题
+ 问题如图，底部最后一行内容被挡住部分

![image](6AC4FD3B506F4604A8901C0C535982CB)
+ 解决方法: 底部添加margin-buttom
  - 给newsList.vue的template中最外层div添加class="tmpl"
  - 在src\static\css\global.css中添加代码
```
    .tmpl {
        margin-bottom: 50px;
    }
```
![image](DBA4D55170894508A97960AE4D57DE64)

### 7.7 日期转换-过滤器组件

#### 7.7.1 问题如图，日期显示格式不好看，太长
![image](FB13B7154B6A404093F25F0ADADE569F)
#### 7.7.2 moment.js插件介绍
+ 这里推荐使用moment.js 插件
  - 插件地址: momentjs.cn
  - 里面也有安装插件的命令 : npm install moment --save
+ moment插件的一些主要用法
  - 日期格式化
```
moment().format('MMMM Do YYYY, h:mm:ss a'); // 六月 21日 2019, 10:10:59 上午
moment().format('dddd');                    // 星期五
moment().format("MMM Do YY");               // 6月 21日 19
moment().format('YYYY [escaped] YYYY');     // 2019 escaped 2019
moment().format();                          // 2019-06-21T10:10:59+08:00
```
  - 相对时间
```
moment("20111031", "YYYYMMDD").fromNow(); // 8 年前
moment("20120620", "YYYYMMDD").fromNow(); // 7 年前
moment().startOf('day').fromNow();        // 10 小时前
moment().endOf('day').fromNow();          // 14 小时内
moment().startOf('hour').fromNow();       // 11 分钟前
```
  - 多语言
```
moment.locale('en');
moment.locale('zh-cn');
```
  - 推荐用法
```
moment().format('YYYY-MM-DD, HH:mm:ss')
```

#### 7.7.3 全局过滤器定义
+ 由于日期过滤很常用，所以这里定义全局过滤器，而不是组件过滤器
+ 在main.js中定义全局过滤器
```
//1.3.6 Moment:引入moment.js插件
import Moment from 'moment';

//1.5 定义全局过滤器或全局组件，方便大家都能用  开始 +++++++++++++
Vue.filter('convertDate',function (value) {  //convertDate为过滤器名称，value为其他地方传进来的值，对这个value进行转换
    return Moment(value).format('YYYY-MM-DD HH:mm:ss a');
});

```
+ 且moment插件这里也不用命令行安装了，因为开始项目的时候就已经安装了。

#### 7.7.4 news\newsList.vue中使用过滤器
```
                            <p>发表时间:{{news.add_time | convertDate}}</p>
```
![image](FC1E5301D89146238894577327823C04)


### 7.8 新闻列表的头部及返回上一层-过滤器组件
#### 7.8.1 效果及思路理解
+ 头部新闻列表
+ 且头部新闻列表可以复用
+ 还有个返回上一层的按钮
+ 通过hello-mui制作

![image](06D5257D11964E94AF8C2BDFA79D0D1B)
#### 7.8.2 回忆父子组件的用法
+ 一句话概括: 
  - 子组件拿值，所以子组件用{{title}}拿props:['title']中title的变量。 
  - 父组件提供值，所以父组件中写上 `<nav-bar title="新闻列表"></nav-bar>` 这种子组件名称的标签来传值。

##### 下面为详细的讲解:
+ main.js: 引用子组件: `import NavBar from './components/common/navBar.vue';` 
+ main.js: 声明全局组件: `Vue.component('navBar',NavBar);`
+ 子组件引用值:  
  - navBar.vue(子组件)的script中用props拿值，template中引用`{{ title }}`
  - navBar.vue(子组件)的script中methods拿值，template的标签中引用`@click=show`
```
  props:['title'],
  methods:{
   show(){
       alert(this.title)
   }
  }
```
+ 最后父组件定义值: newsList.vue(父组件)的template的div标签里加一个 `<nav-bar title="新闻列表"></nav-bar>`标签，这里就是父组件传值的地方。

#### 7.8.3 新闻列表头部制作(父子组件传值方式)
+ 拿到hellomui的头部代码:
  - 打开并复制src\static\vendor\mui\examples\hello-mui\examples\ajax.html 中下面代码
```
	<header class="mui-bar mui-bar-nav">
		<a class="mui-action-back mui-icon mui-icon-left-nav mui-pull-left"></a>
		<h1 class="mui-title">ajax（网络请求）</h1>
	</header>
```
  - 注意: a标签为返回上一层小箭头。 h1标签为"新闻列表"的title标题部分。
+ 新建目录components\common作为常用目录，并新建navBar.vue作为后面引用
  - 目录结构: src\components\common\navBar.vue (子组件)
  - 粘贴代码到navBar.vue的template的div中
  - 修改内容:
    + 添加 @click="goBack" 点击事件，并定义 methods方法:
```
    methods: {
        goBack(){
            this.$router.go(-1);  #返回到上一层
        }
    }
```
    + 添加 {{ title }} 在标签中引用，并定义 props: ['title'] 声明接收父组件参数。
+ main.js中添加全局组件定义和组件引入
```
//1.6 引入全局组件需要的组件对象 开始 +++++++++++++++++++++++++++
import NavBar from './components/common/navBar.vue';

Vue.component('navBar',NavBar);  //(全局组件定义) 标签使用时最好用nav-bar
```

+ 父组件newsList.vue中传值
```
<template>
    <div class="tmpl">

        <nav-bar title="新闻列表"></nav-bar>   <!--3.1 父组件向子组件传值-->

    <!-- MUI 图文列表 -->
```
  - 通过`<nav-bar></nav-bar>`标签中添加`title="新闻列表"`向子组件传值title。

#### 7.8.4 头部标签栏丢失添加回来
+ 如图，发现是定位fixed固定了导致的
![image](0A09D975BC8D4AE49CCE7268AD8728C7)
![image](CC00DB1A93DB42A9B97ADBFC980B69F9)

+ 在navBar.vue中修改样式为相对定位
```
<style scoped>
    .mui-bar.mui-bar-nav{
        position: relative;
    }
</style>
```
![image](EE3626F690D8481E9E7C458335A98143)

## 8 新闻详情
### 8.1 效果图说明

![image](E2005B056BCC44FAA7A1C47DB3E27687)
![image](8D703C54614A49AEA525E3E7FFED85C4)
![image](0E494EDD24474B7A9D6928976B51E309)

+ 注意，上图中content为页面中的主要文字部分，且应该要用v-html方式去显示内容

### 8.2 配置newsList.vue中(去哪里)
+ 编辑newsList.vue
  - 图文列表中ul下li下的a标签改为router-link(上面7.5.2中由于是拷贝的，所以已经改过了),加上 :to="{}" 内容
  - 表示点击会去news.detail对应的组件；query是获取里面的id值，用来后面获取此id，进行拼接url时使用的。
```    
<router-link :to="{name:'news.detail',query:{id:news.id} }">
```

### 8.3 main.js中配置路由规则及引入组件文件(导航)
```
//1.1.4 引入新闻详细文件
import NewsDetail from './components/news/newsDetail.vue';

        // 1.4.4 添加新闻详情的规则
        {name: 'news.detail', path: '/news/detail', component: NewsDetail},
```

### 8.4 newDetail.vue文件创建及代码拷贝修改(去了干嘛)
+ 创建目录及文件 : src\components\news\newsDetail.vue
+ 拷贝代码并修改 
```
<template>
    <div class="tmpl">
        <nav-bar title="新闻详情"></nav-bar>
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
    data(){
        return {
            newsDetail:{}, //用{}或[]都不重要，因为以后都会被重新赋值，所以这里就只是为了象征性的表示其数据类型
        }
    },
    created(){
        //1:获取路由参数
        let id = this.$route.query.id;  //获取newsList.vue中的query的id，给于下面一行代码做拼接使用。
        //2:拼接路由参数成为后台请求的URL(通过拼接 getnew/43 这种url来获取不同的信息)
        this.$ajax.get('getnew/' + id)
        .then(res=>{
             //3:响应回来渲染页面
             this.newsDetail = res.data.message[0];
             //数组，就一个数据，所以直接取[0]
        })
        .catch(err=>{
            console.log(err);
        })

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
```

+ 效果图

![image](05FF0127F7E04F418C53F34071B4589B)
![image](D086E332CD7E4DD4BBB0913A2904F99C)


## 9 主体-九宫格-图文分享页面制作
### 9.1 额外内容: 课程第4天开始，同学反馈部分说下
![image](3563B2826E8248D3BCFA9B1727677746)

### 9.2 api文档及请求介绍
#### 9.2.1 getimgcategory请求后展示的分类id与title
+ 使用 /api/getimgcategory 接口，显示所有的分类信息。
  - 例如: 分类id为14，代表家具生活；分类id为15代表摄影设计。
  - api文档介绍

![image](D2AE3B8BF28A4783972D020CA23C5D1E)
![image](72FFF4B69DB64AB6898CE20784359562)
  - 实际请求效果

![image](679652F5F9E64BB5AE5D6D71E40B55DC)
#### 9.2.2 getimages+分类id，显示对应分类的信息
+ 使用 /api/getimages/:分类id 接口，显示对应分类的信息及对应的分类的图片
  - 当id为0时，获取所有图片数据信息，0代表对应的"全部"这个title。
  - 注意: api中说的getimages/0 应该是后端做好了，把所有的数据都放在0中，所以你请求0的时候，能获取所有信息(图片+其他内容)
  - api文档介绍

![image](81213ECCEC13406BA068B71E6D8F30EB)

  - 例如: 请求分类id为17，则显示对应 img_url和title的内容为 "空间设计" 相关的内容，这里的空间设计，就是上面分类id对应的title 。

![image](E3E0BE22DCCA4B39A65B1FF150892AC1)

### 9.3 思路理解
![image](C4B5608EC7CC4F96A71D165F12918DD3)

### 9.4 开始编写代码
#### 9.4.1 home.vue九宫格中设置去哪里
```
		            <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><router-link :to="{name: 'photo.share'}">
		                    <span class="mui-icon mui-icon-email"></span>
		                    <div class="mui-media-body">图文分享</div></router-link></li>
```
#### 9.4.2 main.js中引入组件文件及定义路由规则
```
//1.1.5 引入图文分享文件
import PhotoShare from './components/photo/photoShare.vue';

        // 1.4.5 添加图文分享的规则
        {name: 'photo.share', path: '/photo/share', component: PhotoShare},
```
#### 9.4.3 创建对应目录及文件
+ 创建目录及文件: src\components\photo\photoShare.vue

#### 9.4.4 粘贴模板代码到photoShare.vue中

#### 9.4.5 自己nginx准备好对应请求内容
+ nginx的目录结构
```
.
|-- 1.jpg
|-- 2.png
|-- 3.jpg
|-- 4.txt
|-- 5.info
|-- 6.info
|-- photoShare
|   |-- getimages  #按不同id获取对应的分类信息
|   |   |-- 0   #此文件处包含所有分类的所有信息，这个应该后端做的自动更新0为所有分类的所有信息的集合文件
|   |   |-- 14
|   |   |-- 15
|   |   |-- 16  #此文件为明星美女分类的信息
|   |   |-- 17
|   |   |-- 18
|   |   |-- 19
|   |   `-- 20
|   `-- getimgcategory  #获取所有分类信息
|-- star01.jpg  #这里是id为16的明星美女分类的3张图片，没弄多的
|-- star02.jpg  #这里是id为16的明星美女分类的3张图片，没弄多的
`-- star03.jpg  #这里是id为16的明星美女分类的3张图片，没弄多的
```
+ id为16(明星美女)分类的json内容
```
[root@rock ~] cd /root/images/photoShare/getimages
[root@rock getimages]# cat 16
[
  {
    "id": 37,
    "title": "张若昀唐艺昕婚礼伴手礼曝光 牵手拍婚纱照甜笑恩爱",
    "img_url": "http://10.5.1.80:6001/images/star01.jpg",
    "zhaiyao": "新浪娱乐讯 6月25日，小浪一早收到张若昀唐艺昕新人伴手礼，除了情侣香水，化妆品和永生花朵，还有两张甜蜜婚纱照，二人亲密牵手，恩爱甜笑，超美的！卡片上印有叶芝的诗《当你老了》和“Z&T 2019 06 27”，温馨浪漫。"
  },
  {
    "id": 38,
    "title": "2019香港小姐竞选面试美女如云 众佳丽齐穿华服美裙亮相",
    "img_url": "http://10.5.1.80:6001/images/star02.jpg",
    "zhaiyao": "新浪娱乐讯 6月24日，众佳丽前往参加2019香港小姐竞选面试。当天，众佳丽齐穿华服美裙盛装亮相，对镜凹造型星范儿满满，现场美女如云，你看谁有冠军相？"
  },
  {
    "id": 39,
    "title": "杨幂穿粉色包臀裙吸睛力满分 走路带风撩发秀美腿",
    "img_url": "http://10.5.1.80:6001/images/star03.jpg",
    "zhaiyao": "新浪娱乐讯 6月24日，杨幂北京现身机场。她头戴白色鸭舌帽显低调，身穿白色T恤下搭玫粉色包臀裙，与荧光粉运动鞋相呼应十分抢镜，走路带风大秀美腿，撩发尽显女人味。"
  },
  ... ...
]
```
+ 全部分类的信息概括文件
```
[root@rock photoShare]# cat getimgcategory
[
{
  "title": "家居生活",
  "id": 14
},
{
  "title": "摄影设计",
  "id": 15
},
{
  "title": "明星美女",
  "id": 16
},
{
  "title": "空间设计",
  "id": 17
},
{
  "title": "户型装饰",
  "id": 18
},
{
  "title": "广告摄影",
  "id": 19
},
{
  "title": "摄影学习",
  "id": 20
}
]
```

+ 我这里自己出现的问题--以后一定要注意
  - 由于我没有api请求的后端，自己写了一个文件，然后想着转为json做数据就行。
  - 问题出现在:
    + 我定义的文件 17 ，在nginx中，由于json格式没有写好，少了逗号，导致一直以为是字符串json转json对象的问题。
    + 还有个问题，就是id,title等key上没有加双引号。
  - 其实Axios默认是会帮你转成json对象的，不用在用JSON.parse()这个方法了。
  - 可以写好json内容后，去网上json格式化校验一下，如果成功则表示json写的没错。
  - 这里简单的看下Axios获取nginx的数据的写法
```
            loadImg(id) {
                // this.$ajax.get('getimages/' + id)
                // this.$ajax.get('photoShare/getimages/' + id)
                this.$ajax.get('photoShare/getimages/' + 17)
                    .then(res => {
                        //console.log(typeof(res.data));  //#正常都是string字符串格式的，但没有影响
                        this.imgs = res.data;
                        // this.imgs = JSON.parse(res.data);  //#这里不用JSON.parse转换也可以，虽然是字符串，但是Axios是可以自动转换的
```

#### 9.4.6 代码理解
```
<template>
    <div class="tmpl">
        <nav-bar title="图文分享"></nav-bar>  <!-- 1.8 标题栏加上 -->
        <!-- 引入返回导航 -->
        <div class="photo-header">
            <ul>
                <li v-for="category in categorys" :key="category.id">
                    <a href="javascript:;" @click="loadImg(category.id)">{{category.title}}</a>   <!-- 1.5 @click点击事件，传入点击内容的id -->
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
            // this.$ajax.get('getimgcategory')   //老师的代码
            this.$ajax.get('photoShare/getimgcategory')
                .then(res => {
                    // this.categorys = res.data.message;  //老师的代码
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

            //1.7 当页面加载默认传递0，因为0代表全部，所以刚加载的时候就是加载全部内容
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
            //1.6 使用loadImg方法，替代上面1.4的操作，所以上面1.4全部注释
            loadImg(id) {
                // this.$ajax.get('getimages/' + id)  //老师的代码
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

    .photo-header a:hover{
        background-color: skyblue;
    }
</style>

```

### 9.5 懒加载
#### 9.5.1 懒加载是使用mint-ui插件
+ 地址: http://mint-ui.github.io/docs/#/zh-cn2/lazyload
+ 懒加载的代码
```
<ul>
  <li v-for="item in list">
    <img v-lazy="item">
  </li>
</ul>
```
+ 图片懒加载的样式代码
```
image[lazy=loading] {
  width: 40px;
  height: 300px;
  margin: auto;
}
```

#### 9.5.2 懒加载使用
+ 在photoShare.vue中，替换 `:src="img.img_url"` 为 `v-lazy="img.img_url"`
```
                        <!--<img :src="img.img_url">-->
                        <!-- 2.2 懒加载 -->
                        <img v-lazy="img.img_url">
```
+ 在photoShare.vue中，加上css样式
```
    /*2.1 图片懒加载的样式*/
    image[lazy=loading] {
        width: 40px;
        height: 300px;
        margin: auto;
    }
```

## 10 图文详情
### 10.1 需要实现效果

![image](0563ACB78A574FC5883288A9DEF85EF6)
![image](1289883BA8F24EFFAB8142E4D53931D5)

### 10.2 思路理解
![image](626F78518B1A4EA4AC93ACDBC710D0E1)

### 10.3 开始编写代码(基础布局)
#### 10.3.1 photoShare.vue中定义router-link去哪里
```
                    <router-link :to="{name:'photo.detail',params:{id:img.id} }">  <!-- 3.1 图文详情: 去哪里路由定义 -->
                        <!--<img :src="img.img_url">-->
                        <!-- 2.2 懒加载 -->
                        <img v-lazy="img.img_url">
                        ... ...
                    </router-link>
```

==注意:== 
+ 这里的img.id，来自于photoShare.vue的created()函数中通过$ajax请求的`this.$ajax.get('photoShare/getimages/' + id)`的值(res.data.message)中的id。
+ 而这里的getimages+id的内容是获取图片的信息内容，里面有id,title,zhaiyao,img_url。
+ 这里之所以说下，就是怕下面"图文详细"部分引用此id的时候，Rock自己测试报错，找了半天才知道id的来源。

#### 10.3.2 main.js中引入组件文件及路由定义
```
//1.1.6 引入图文详细文件
import PhotoDetail from './components/photo/photoDetail.vue';

        // 1.4.6 添加图文详细的规则(这里的:id 和 photoShare.vue中的params:{id:img.id}这个id对应)
        {name: 'photo.detail', path: '/photo/detail/:id', component: PhotoDetail},
```

==**注意**==: 路由规则中的 :id 和 photoShare.vue中的params:{id:img.id}这个id是相互对应的。

#### 10.3.3 粘贴模板代码到photoDetail.vue中
+ 贴入代码
==注意:== 这里是老师的代码，我的代码在10.5.5中改了，因为Rock的vue-preview是新版本的，代码和老师的旧版本不同的。
```
<template>
    <div class="tmpl">
        <!--  组件名navBar -->
        <nav-bar title="图片详情"></nav-bar>
        <!-- 组件名:navbar -->
        <!--  使用：navbar-->
        <div class="photo-title">
            <p v-text="imgInfo.title"></p>
            <span>发起日期：{{imgInfo.add_time | convertDate}}</span>
            <span>{{imgInfo.click}}次浏览</span>
            <span>分类：民生经济</span>
        </div>
        <ul class="mui-table-view mui-grid-view mui-grid-9">
            <li v-for="(img,index) in imgs"  :key="index"  class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
                <!-- <img :src="img.src"> -->
                 <img class="preview-img" :src="img.src" height="100" @click="$preview.open(index, imgs)">
            </li>
        </ul>
        <div class="photo-desc">
            <p v-html="imgInfo.content"></p>
        </div>

        <!-- 使用评论子组件 -->
        <comment :cid="pid"></comment>
    </div>
</template>
<script>
export default {
    data(){
        return {
            imgs:[],//存放缩略图
            imgInfo:{},//详情数据对象
            pid:this.$route.params.id, //记录当前图片id

        }
    },
    created(){
        //1:获取photoShare.vue中用户点击页面内容的路由id参数
        // let pid = this.$route.params.id;
        //2:发起请求2个
        //2.1:获取图片对应的详情
        this.$ajax.get('getimageInfo/' + this.pid)
        .then(res=>{
            //一个id对应一个详情对象,所以用message[0]
            this.imgInfo = res.data.message[0];
        })
        .catch(err=>{
            console.log(err)
        });
        //2.2:获取缩略图(多个缩略图)的地址
        this.$ajax.get('getthumimages/' + this.pid)
        .then(res=>{
            this.imgs = res.data.message;

            //forEach设置每个图片的高、宽
            this.imgs.forEach((ele)=>{
                ele.w=300;
                ele.h=200; //设置缩略图显示的高
            })
        })
        .catch(err=>{
            console.log(err)
        })
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
</style>
```

+ script中加上created()函数
```
    created(){
        //1:获取photoShare.vue中用户点击页面内容的路由id参数
        // let pid = this.$route.params.id;
        //2:发起请求2个
        //2.1:获取图片对应的详情
        this.$ajax.get('getimageInfo/' + this.pid)
        .then(res=>{
            //一个id对应一个详情对象,所以用message[0]
            this.imgInfo = res.data.message[0];
        })
        .catch(err=>{
            console.log(err)
        });
        //2.2:获取缩略图(多个缩略图)的地址
        this.$ajax.get('getthumimages/' + this.pid)
        .then(res=>{
            this.imgs = res.data.message;

            //forEach设置每个图片的高、宽
            this.imgs.forEach((ele)=>{
                ele.w=300;
                ele.h=200; //设置缩略图显示的高
            })
        })
        .catch(err=>{
            console.log(err)
        })
    }
```

+ photoDetail.vue 的template中填入对应的数据
```
        <div class="photo-title">
            <p v-text="imgInfo.title"></p>
            <span>发起日期：{{imgInfo.add_time | convertDate}}</span>
            <span>{{imgInfo.click}}次浏览</span>
            <span>分类：民生经济</span>
        </div>
        <ul class="mui-table-view mui-grid-view mui-grid-9">
            <li v-for="(img,index) in imgs"  :key="index"  class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
                <!-- <img :src="img.src"> -->
                 <img class="preview-img" :src="img.src" height="100" @click="$preview.open(index, imgs)">
            </li>
        </ul>
```

### 10.4 图文详情页面--图片预览(小图+预览图)
#### 10.4.1 图片预览插件(vue-preview)使用及注意事项
==注意:== 此处(10.5.1中)只是介绍和注意，不用进行操作，下面10.5.2中会有操作部分
+ 插件地址: github.com/Ls1231/vue-preview
+ 效果图

![image](AC4DDA26E0344840B1FC95CBE8E34EB4)

![image](5D5290ABC63043C7B8A1A44974B11821)
+ ==注意:== 上图中提示: 如果使用的是 vue-cli 生成的项目，需要修改webpack.base.conf.js文件中loaders，添加一个loader。
  - 原因: 插件编写中使用了es6语法，需要进行代码编译。
  ```
  {
      test: /vue-preview.src.*?js$/,
      loader: 'babel'
  }
  ```
+ 安装vue-preview命令: `npm i vue-preview -S`
+ 安装插件(介绍)
```
import VuePreview from 'vue-preview'
Vue.use(VuePreview)
```
+ 使用插件的样例
```
<template>
    <!--这里的height="100"是设置页面中每个小图片的高度，不是点开后的高度-->
    <img class="preview-img" v-for="(item,index) in list" :src="item.src" height="100" @click="$preview.open(index, list)">
</template>

<script>
    export default {
        data(){
            return {
                list: [{  //在总的list中定义每个图片的宽和高，当然你也可以考虑用forEach这种循环设置每个图片相同大小。
                          //下面老师的代码就是用forEach设定每个图片的大小相同的。就不用这里每个都定义了。
                    src: 'https://placekitten.com/600/400',
                    // 下面是设置每个图片点开后的大小
                    w: 600,  //定义图片的宽
                    h: 400   //定义图片的高
                },{
                    src: 'http://placekitten.com/1200/900',
                    w: 1200,
                    h: 900
                }]
            }
        }
    }
</script>
```
==注意:== 
老师在官网(github.com/Ls1231/vue-preview)看的vue-preview 和 Rock看的版本已经不同了(视频是2017的，我是2019.6看的)，所以里面的样例模板也已经不同。

#### 10.4.2 根据上面10.4.1中提到修改webpack.base.conf.js问题
==注意:== 从此处开始操作
+ 虽然老师这里用的是webpack，不是用vue-cli脚手架生成的项目，但是也还是需要做解析es6代码的
+ vue-preview开始就安装过了，这里就不用安装了。
+ 编辑webpack.base.conf.js 文件
  - 下面内容中 :
    + 将`test: /\.js$/,`中的options对应内容注释。
    + 新增vue-preview的es代码解析部分(`test: /vue-preview.src.*?js$/,`)，且也不使用options。
```
            {
                test: /\.js$/,
                loader: 'babel-loader',
                exclude: /node_modules/,  //需要排除，否则报错nodemodule\\lodash\\lodash.js as it exceeds the max of '500KB'
                // options: {  //如果多次使用babel-loader就需要多次使用这个options，所以这里注释掉，
                //               建议使用 .babelrc 文件，再当前根目录就可以了
                //     presets: ['es2015'], //关键字
                //     plugins: ['transform-runtime'], //函数
                // }
            },
            // 解析 vue-preview 的es代码
            {
                test: /vue-preview.src.*?js$/,
                loader: 'babel-loader',
                // options: {  //如果多次使用babel-loader就需要多次使用这个options，所以这里注释掉，
                //               建议使用 .babelrc 文件，再当前根目录就可以了
                //     presets: ['es2015'], //关键字
                //     plugins: ['transform-runtime'], //函数
                // }                
            },
```
+ 然后根目录创建 .babelrc 文件
  - 和src同级目录，创建 .babelrc
  - `cms_project_27\.babelrc` 和 `cms_project_27\src` 同级文件
+ 编写 .babelrc 文件(将原来options中内容拷贝过来即可)
```
{
    "presets": ["es2015"],
    "plugins": ["transform-runtime"]
}
```

#### 10.4.3 main.js中引入vue-preview插件
```
//1.3.7 VuePreview:引入vue-preview
import VuePreview from 'vue-preview';
//挂载使用
Vue.use(VuePreview);
```
==注意:== 
这里挂载了VuePreview，则表示后面vue文件中可以调用this.$preview(script中使用)和$preview(template中使用)这种方法去使用。

#### 10.4.4 photoDetail.vue中使用vue-preview
```
            <li v-for="(img,index) in imgs"  :key="index"  class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
                 <!--<img :src="img.src">-->  <!--下面使用vue-preview，所以这一行就注释了-->
                 <!-- 这里使用 v-for中的img.src， $preview.open是它的用法，里面放入index 和 总的imgs的信息 -->
                 <img class="preview-img" :src="img.src" height="100" @click="$preview.open(index, imgs)">
            </li>
```

#### 10.4.5 新版vue-preview用法
==注意:== 
由于Rock这里用老师的提示报错，因为我npm install vue-preview是新版本的了，所以必须更新为新代码了。下面为新版本的vue-preview。
+ main.js中引入vue-preview和使用
```
import VuePreview from 'vue-preview';

Vue.use(VuePreview, {
  mainClass: 'pswp--minimal--dark',
  barsSize: {top: 0, bottom: 0},
  captionEl: false,
  fullscreenEl: false,
  shareEl: false,
  bgOpacity: 0.85,
  tapToClose: true,
  tapToToggleControls: false
});

```

+ photoDetail.vue中的代码
```
<template>
    <div class="tmpl">
        <!--  组件名navBar -->
        <nav-bar title="图片详情"></nav-bar>
        <!-- 组件名:navbar -->
        <!--  使用：navbar-->
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
                <!--这里的height="100"是设置页面中每个小图片的高度，不是点开后的高度-->
                <!--<img class="preview-img" :src="img.src" height="100" @click="$preview.open(index, imgs)">-->
                <!--下面是新版本vue-preview用法，上面为vue-preview老版本用法-->
            </li>
            <!--imgs为所有的图片信息，handleClose不用管，必须加的，且在methos中添加此方法-->
            <!--注意: 这里是imgs，而按老师的来，则会在for中套imgs，会导致按倍数显示图片，解决方法就是离开上面的for循环即可-->
            <vue-preview :slides="imgs" @close="handleClose"></vue-preview>
        </ul>
        <div class="photo-desc">
            <p v-html="imgInfo.content"></p>
        </div>

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
                pid: this.$route.params.id, //记录当前图片id
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
                })
        },
        methods: {  
            handleClose() {  //新版本vue-preview要加这个方法
                console.log('close event')
            }
        }
    }
</script>
<style scoped>
    ... ...
</style>
```

+ photoDetail.vue 中vue-preview要注意的问题:
  - 1. template中的vue-preview
`<vue-preview :slides="imgs" @close="handleClose"></vue-preview>`
  - 2. created()中的设置w(宽)、h(高)、msrc=src(小图=大图)
```
                      //forEach（es6的语法）设置每个图片的高、宽
                    this.imgs.forEach((ele) => {
                        // 下面是设置每个图片点开后的大小(即缩略图)
                        ele.w = 400;
                        ele.h = 500; //设置缩略图显示的高
                        ele.msrc = ele.src;  //新版本vue-preview要加这一行，msrc为小图片，src为点击后的大图片。
                    })
```
  - 3. handleClose()关闭事件
```
            handleClose() {  //新版本vue-preview要加这个方法
                console.log('close event')
            }
```


+ static\css\global.css中添加样式来固定页面中每个小图片的高度
```
.my-gallery figure a img {
    height: 100px; /*每个图片的高度*/
    margin: 10px 0 10px 0; /*图片上下之间的间距*/
    box-shadow: 0 0 10px #ccc; /*图片边框的阴影*/
}

.my-gallery figure {
    display: inline-block;  /*让图片按行显示，否则默认是一行只有一个图片*/
}
```

+ 遗留一个问题
  - 发现在nginx中配置的 /root/images/photoDetail/getthumimages/37 中定义的src为4个图片，页面中则显示每个图片x4，共16张图片。
  - 如果设置2张，则显示2x2=4张图片。总是成倍的显示。
  - 暂时怀疑是forEach导致的。
![image](6C824EC828894201B953F04769FD48F2)


### 10.5 图文详情--评论部分
#### 10.5.1 评论部分理解

![image](760B496337C84D55892910B66344F2E0)
![image](7B0C97DEE8A44030A6FB0C8879751DF1)

#### 10.5.2 photoDetail.vue中添加评论部分代码
+ ==注意:== 后面会把下面的代码放入公共模板中，因为其他地方也要使用评论模块
+ template中添加模板代码
```
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
                <li v-for="(comment,index) in comments" :key="index">
                    {{comment.user_name}}：{{comment.content}} {{comment.add_time|convertDate}}
                </li>
               
            </ul>
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
```
+ style中添加样式代码
```
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
```












# heros_detail_info
## 1.Project setup
```
npm install
```
### Compiles and hot-reloads for development
```
npm run serve
```
### Compiles and minifies for production
```
npm run build
```
### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


## 2.安装插件
### 安装axios插件
```
npm install axios --save
npm install vue-router --save
```

## 3.解决问题
### 解决跨域问题
+ 根目录下创建`vue.config.js`
```
const path = require('path');

const resolve = dir => {
    return path.join(__dirname, dir)
};

module.exports = {
    publicPath: '/',
    // baseUrl: '/',
    // outputDir: 'dist', //构建输出目录，默认就是dist
    // assetsDir: 'assets', //静态资源目录(js,css,img,fonts)
    // 如果你不需要使用eslint，把lintOnSave设为false即可
    lintOnSave: false,
    chainWebpack: config => {
        config.resolve.alias
            .set('@', resolve('src')) // key,value自行定义，比如.set('@@', resolve('src/components'))
        //     .set('_c', resolve('src/components'))
    },
    // 打包时不生成.map文件
    productionSourceMap: false,
    // 这里写你调用接口的基础路径，来解决跨域，如果设置了代理，那你本地开发环境的axios的baseUrl要写为 '' ，即空字符串
    devServer: {
        open: true, // run项目后自动打开网页
        // host: '127.0.0.1', // 真实项目建议0.0.0.0
        port: 8080, // 项目端口
        // https: false, //https功能
        //跨域配置
        proxy: {
            // v1表示api的请求地址是: http://localhost:5000/v1/ 这里url第一个斜杠/ 后代表的内容，即下面target也写了，是v1
            '/v1': {
                // 跨域就是vue服务的端口，和请求其他(如flask)端口不同，则表示跨域了
                target: ' http://10.151.106.247:5001',
                ws: true,
                changeOrigin: true,  // 是否跨域
                //rewrite 表示以v1开头的
                // pathRewrite: {
                //     '^/v1': ''
                // }
            }
        }
    }
};
```
+ main.js中最好配置下baseURL，方便后面访问
  `axios.defaults.baseURL = "http://10.151.106.247:5000/v1/";`
+ 重要说明
  - 先尝试访问`http://10.151.106.247:5000`对应的api，如果浏览器可以访问则没事。
  - 如果不能访问，应该是windows问题，可以考虑更换为5001端口，这个是flask的运行端口的修改，需要确认flask和vue的配置一致。

### 解决锚点问题
+ main.js 的router定义中加入
+ 解决 `http://127.0.0.1:8080/#/xxxx`中的 `#/` 卡在中间的问题，这个是锚点
```
let router = new VueRouter({
    mode: 'history',  // 用于解决 url后有个 "#/" 存在的问题，加了这个，则不显示此锚点了
```
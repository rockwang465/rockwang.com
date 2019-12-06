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

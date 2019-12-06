const path = require('path');

const resolve = dir => {
    return path.join(__dirname, dir)
};

module.exports = {
    publicPath: '/',
    lintOnSave: true,
    chainWebpack: config => {
        config.resolve.alias
            .set('@', resolve('src')) // key,value自行定义，比如.set('@@', resolve('src/components'))
            .set('_c', resolve('src/components'))
    },
    productionSourceMap: false,
    devServer: {
        port: 8080,
        // host: '10.151.106.247',
        proxy: {
            '/v1': {
                target: ' http://10.151.106.247:5001',
                ws: true,
                changeOrigin: true
            }
        }
    }
};
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
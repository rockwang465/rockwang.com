# translate 翻译页面制作
## 1.Project setup
+ create my_project
```
vue create my_project
```
+ install modules
```
npm install
```
+Compiles and hot-reloads for development
```
npm run serve
```
+ Compiles and minifies for production
```
npm run build
```
+ Lints and fixes files
```
npm run lint
```
+ Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

## 2. 操作部分
### A.`Hello World.vue`删除
 + 注意很多地方引用，需要通过搜索方式去删除最干净。
### B. 报错问题
+ 双引号换成单引号
+ 缩进报错问题
  - 需要在package.json的rules中设置
  ```
    "rules": {
      "indent": "off"
    },
  ```
  - 不加则有类似报错:
  `error: Expected indentation of 2 spaces but found 4 (indent) at src\components\translateForm.vue:18:1:`
+ 关闭eslint功能
  - eslint用于检测语法，非常烦人，建议关闭
  - 删除`package.json`中此部分内容即可
  ```
  "eslintConfig": {
    "root": true,
    "env": {
      "node": true
    },
    "extends": [
      "plugin:vue/essential",
      "@vue/standard"
    ],
    "rules": {
      "indent": "off"
    },
    "parserOptions": {
      "parser": "babel-eslint"
    }
  },
  ```
## 3.翻译api工具
### A.获取api工具key
+ 打开`https://tech.yandex.com/`
+ 最下方找到 `Translate API`，点击进入
+ 点击下方: Get a free [API key](https://translate.yandex.com/developers/keys)
+ 另外: 点击之前，还有个: Read the [documentation](https://tech.yandex.com/translate/doc/dg/concepts/About-docpage), where ...,这个是使用文档。
+ 登录google用户，并点击创建新的key`create a new key`
+ 下面是rock创建的key
```
trnsl.1.1.20191202T081747Z.69aacfde7fc5e5c3.cd884f584fcd2b26954a68bbaeeb6ebf938008f1
```
### B.插件安装
+ 需要调用上面的api，所以这里老师用了`vue-resource`
```
npm install vue-resource 或 axios --save
```
+ `main.js`中引用`vue-resource`
```
cd 项目路径
import VueResource from 'vue-resource'
Vue.use(VueResource)
```
+ `main.js`中引用`axios`
```
// 添加axios用于http请求的，rock另外学习测试下
import Axios from 'axios'

//挂载原型(vue-resource和axios挂载方式不同)
Vue.prototype.$axios = Axios;
```

### C.api使用
+ 在A中说了，有个[documentation](https://tech.yandex.com/translate/doc/dg/concepts/About-docpage)
+ 进去，点击`Translate text`
+ 在`Reqqquest syntax`下面有个url，此为请求的url:
  - `https://translate.yandex.net/api/v1.5/tr.json/translate`
  - `? key=<API key>`
  - `& text=<text to translate>`
  - `& lang=<translation direction>`
  - `& [format=<text format>]`
  - `& [options=<translation options>]`
  - `& [callback=<name of the callback function>]`
  - `key=<上面创建的key>`
  - 这里正常只需要前面4个即可。
+ 代码写法:
```
methoes:{
  translateText: function(text){
  // lang=en 表示翻译为英文
  // text='翻译内容' 表示传入需要翻译的内容
  this.$http.get('https://translate.yandex.net/api/v1.5/tr.json/translate?key=trnsl.1.1.20191202T081747Z.69aacfde7fc5e5c3.cd884f584fcd2b26954a68bbaeeb6ebf938008f1&lang=en&text='+text)
    .then((response)=>{
      console.log(response.body.text[0])  //拿到翻译的值
      
    })
  }
}
```
+ 可以考虑这么写
```
    export default {
        ... ...
        data: function () {
            return {
                translate_body: 'App default value',
                translate_url: 'https://translate.yandex.net/api/v1.5/tr.json/translate',
                translate_key: '?key=trnsl.1.1.20191202T081747Z.69aacfde7fc5e5c3.cd884f584fcd2b26954a68bbaeeb6ebf938008f1',
                translate_lang: '&lang=en',

            }
        },
        methods: {
            // 拿到子组件传来的用户输入的内容，通过api调用，进行翻译成英文，并发送给另一个显示翻译后内容的组件(translateOutput)
            translate_text(value) {  /*定义方法，获取子组件传过来的值*/
                this.translate_body = value;
                // console.log(this.translate_url + this.translate_key + '&text=' + this.translate_body + this.translate_lang);
                this.$http.get(this.translate_url + this.translate_key + '&text=' + this.translate_body + this.translate_lang)
                    .then(function (response){
                        console.log(response.body.text[0]); //拿到请求结果的第一个内容
                    })
            }
        }
    }
```

### D.多语言翻译
+ 这里有每个语言对应的ui code : `https://tech.yandex.com/translate/doc/dg/concepts/api-overview-docpage/#api-overview__languages`
+ 即传入的`lang=`后的code，例如:en为英语、ko为韩语、ja为日语...
+ 然后使用v-model绑定语言选项select标签，知道用户选择哪个语言，并放入子组件向父组件传参的函数中。
+ 如下，上面定义了v-model在select中，并传入了2个参数。
```
                <select v-model='lang'>  // 绑定了语言事件，用于传参选择语言
                    <option value='en'>英语</option>
                    <option value='ko'>韩语</option>
                    <option value='ja'>日语</option>
                    <option value='th'>泰语</option>
                    <option value='fr'>法语</option>
                    <option value='ru'>俄语</option>
                </select>
                <button class="translate_button" @click="translate_submit">翻译</button>
                ... ...
        methods: {
            translate_submit() {
                this.$emit('get_text', this.translate_text, this.lang); //给父组件传需要翻译的内容和语言，这里传了2个值哦

            }
        }
```

+ ok,至此已全部完成基本框架，具体见代码。
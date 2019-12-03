<template>
    <div id='app'>
        <h1>在线翻译</h1>
        <p>简单 / 易用 / 便捷</p>
        <TranslateForm @get_text='translate_text'></TranslateForm>
        <TranslateOutput :translate_result=translate_result></TranslateOutput>
    </div>
</template>

<script>
    import TranslateForm from './components/translateForm'
    import TranslateOutput from './components/translateOutput'

    export default {
        name: 'app',
        components: {
            TranslateForm,
            TranslateOutput
        },
        data: function () {
            return {
                translate_body: 'App default value',
                translate_url: 'https://translate.yandex.net/api/v1.5/tr.json/translate',
                translate_key: '?key=trnsl.1.1.20191202T081747Z.69aacfde7fc5e5c3.cd884f584fcd2b26954a68bbaeeb6ebf938008f1',
                translate_lang: '&lang=',
                translate_result: '欢迎使用在线翻译'
            }
        },
        methods: {
            // 拿到子组件传来的用户输入的内容，通过api调用，进行翻译成对应语言的内容，并发送给另一个显示翻译后内容的组件(translateOutput)
            translate_text(text,lang) {  /*定义方法，获取子组件传过来的值*/
                // console.log(this.translate_body)
                this.translate_body = text;  //传入需要翻译的字符串
                // this.translate_lang = this.translate_lang + lang;  // 传入需要翻译的语言

                // // A.vue-resoure请求
                // // console.log(this.translate_url + this.translate_key + '&text=' + this.translate_body + this.translate_lang);
                // this.$http.get(this.translate_url + this.translate_key + '&text=' + this.translate_body + this.translate_lang)
                //     .then(function (response) {
                //         this.translate_result = response.body.text[0];
                //         console.log(response.body.text[0]); //拿到翻译结果的第一个内容
                //     })

                // B.axios请求
                this.$axios.get(this.translate_url + this.translate_key + '&text=' + this.translate_body + this.translate_lang + lang)
                    .then(res=>{
                        // console.log(res.data.text[0]); //拿到翻译结果的第一个内容
                        this.translate_result = res.data.text[0]
                    })
                    .catch(err=>{
                        console.log("报错----------------------->");
                        console.log(err);
                    })

            }
        }
    }
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    h1, p {
        text-align: center; /*h1标签居中*/
    }

    p {
        color: darkgray;
    }
</style>

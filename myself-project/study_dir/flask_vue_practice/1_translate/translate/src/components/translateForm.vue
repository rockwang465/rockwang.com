<template>
    <div class='translateForm_main'>
        <div class='outside_info'>
            <div class='inner_info'>
                <input type='text' v-model="translate_text" placeholder='请输入中文'>
                <select v-model='lang'>
                    <option value='en'>英语</option>
                    <option value='ko'>韩语</option>
                    <option value='ja'>日语</option>
                    <option value='th'>泰语</option>
                    <option value='fr'>法语</option>
                    <option value='ru'>俄语</option>
                </select>
                <button class="translate_button" @click="translate_submit">翻译</button>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'translateForm',
        data: function () {
            return {
                translate_text: '',
                lang: 'en' //设置默认语言
            }
        },
        methods: {
            translate_submit() {
                // alert(this.translate_text);
                // e.preventDefault()  // 取消默认事件触发
                this.$emit('get_text', this.translate_text, this.lang); //给父组件传需要翻译的内容和语言，这里传了2个值哦
                // this.translate_text = ''
            }
        }
    }
</script>

<style scoped>
    .translateForm_main {
        width: 600px;
        height: 100px;
        margin: 0 auto;
        /*border: 1px solid blue;*/
    }

    .outside_info {
        width: 500px;
        height: 80px;
        border: 1px solid dimgray;
        border-radius: 10px;
        background-color: whitesmoke;
        margin: 0 auto;
    }

    .inner_info {
        width: 350px;
        height: 40px;
        /*background-color: red;*/
        margin: 25px auto;
    }

    .translate_button {
        background-color: dodgerblue;
        margin-left: 5px;
        border-radius: 3px;
    }
</style>

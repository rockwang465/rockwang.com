<template>
    <div class="translateOutput_main">
        <div class="outside_info">
            <div class="inner_info">
                <!--<h3>展示翻译结果</h3>-->
                <h3>{{translate_result}}</h3>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'translateForm',
        props: {
            translate_result:{
                type: String,
                default: '默认为空哦'
            }
        }
    }
</script>

<style scoped>
    .translateOutput_main {
        width: 600px;
        height: 100px;
        margin: 0 auto;
        /*border: 1px solid mediumvioletred;*/
    }

    .outside_info {
        width: 500px;
        height: 80px;
        border: 1px solid white;
        border-radius: 10px;
        /*background-color: whitesmoke;*/
        margin: 0 auto;
    }
</style>

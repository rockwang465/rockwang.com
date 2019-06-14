<template>
    <!-- 只能有1个根节点 -->
    <div>   
        <pre>
                * v-text 是元素的innerText只能在双标签中使用
                * v-html 是元素的innerHTML，不能包含<!-- {{xxx}} -->
                * v-if 元素是否移除或者插入
                * v-show 元素是否显示或隐藏
                * v-model 
        </pre>

        v-text :
        <span v-text="text"></span>
        <input type="text" name="" v-text="text">
        <hr/>
        v-html :
        <span v-html="html"></span>
         <hr/>
         v-if :
         <div v-if="isShow" style="height:100px;background-color:red;"></div>
          <hr/>
          v-show:
        <div v-show="isShow" style="height:100px;background-color:green;"></div>
          <hr/>
          v-model:
          <input type="text" name="" v-model="mTest">
          {{mTest}}
          <br/>给下面的input的value赋值用v-bind:value="mTest"
          <input type="text" name="" v-bind:value="mTest">
        

    </div>
</template>
<script>
    export default {
        data(){
            return {
                text:'我是v-text玩的东西',
                html:`
                    <ul>
                        <li>哈哈</li>
                        <li>呵呵</li>
                    </ul>
                `,
                isShow:false,
                mTest:'haha'
            }
        }
    }
</script>
<style>
    
</style>
//commonjs引入cal对象
// var cal = require('./cal.js');

//es6的模块引入
import cal from './cal.js';

//获取按钮
document.getElementById('btn').onclick = function(){
    var n1 = document.getElementById('n1').value - 0;
    var n2 = document.getElementById('n2').value - 0;
    var sum = cal.add(n1,n2);
    document.getElementById('result').value = sum;
}
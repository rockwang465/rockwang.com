var temp1 = '我是默认导出的结果';
export default temp1;
// 导入方式是  import xxx from './cal.js'

//声明式导出
export var obj1 = '我是声明式导出1';
export var obj2 = '我是声明式导出2';
export var obj3 = '我是声明式导出3';


// 导入方式是
// import {obj1,obj2} from './cal.js';

//另一种方式声明导出
var obj4 = '我是声明式导出4'
export {obj4};
// 导入方式是
// import {obj1,obj2,obj4} from './cal.js';



package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rockwang.com/rock-second-demo/demo/response"
)

// 返回panic的报错给用户界面，而不是直接被panic了。需要在路由中Use。
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Response(ctx, 500, 500001, fmt.Sprint(err), nil)
				// 以下为顺义的写法，判断逻辑详细
				//    switch e := err.(type) {
				//        case *utils.ConsoleError:
				//                logger.Error(e)
				//                ctx.JSON(e.HttpCode, gin.H{"error": e.Error(),"error_code": e.ErrCode})
				//                ctx.Abort()
				//                return
				//        case *drone.DroneError:
				//                logger.Error(e)
				//                ctx.JSON(e.HttpCode, gin.H{"error": e.Error(),"error_code": e.ErrCode})
				//                ctx.Abort()
				//                return
				//        case *mysql.MySQLError:
				//                logger.Errorf("Mysql error with num %v and message is: %v", e.Number, e.Message)
				//                ctx.JSON(500, gin.H{"error": fmt.Sprintf("Mysql error with num %v and message is: %v", e.Number, e.Message),"error_code": 50000005})
				//                ctx.Abort()
				//                return
				//        case *k8sErr.StatusError:
				//                logger.Errorf("K8s error with code %v and message is: %v", e.ErrStatus.Code, e.ErrStatus.Message)
				//                ctx.JSON(int(e.ErrStatus.Code), gin.H{"error": e.ErrStatus.Message,"error_code": 50000006})
				//                ctx.Abort()
				//                return
				//        case validator.ValidationErrors:
				//                msg := "Params aren't valid because: "
				//
				//                for _, error := range e {
				//                        msg += fmt.Sprintf("%v 's value %v is not valid on field(%v); ",
				//                                error.Namespace(), error.Value(), error.Tag())
				//                }
				//                logger.Errorf("%v", msg)
				//                ctx.JSON(http.StatusBadRequest, gin.H{"error": msg,"error_code": 40000001})
				//                ctx.Abort()
				//                return
				//        default:
				//                logger.Error(err)
				//                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err,"error_code": 50000001})
				//                ctx.Abort()
				//                return
				//        }

			}
		}()
		ctx.Next()
	}
}
// 这里判断太简单了，详细的要看顺义的console中的写法
// 具体写法文件是 infra-console-service中的error_handler.go (/gitlab.sz.sensetime.com/galaxias/infra-console-service/console/server/middleware)
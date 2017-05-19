package goio

const (
	MQTT_CONNECT     = 1  //请求连接
	MQTT_CONNACK     = 2  //请求应答
	MQTT_PUBLISH     = 3  //发布消息
	MQTT_PUBACK      = 4  //发布应答
	MQTT_PUBREC      = 5  //发布已接收，保证传递1
	MQTT_PUBREL      = 6  //发布释放，保证传递2
	MQTT_PUBCOMP     = 7  //发布完成，保证传递3
	MQTT_SUBSCRIBE   = 8  //订阅请求
	MQTT_SUBACK      = 9  //订阅应答
	MQTT_UNSUBSCRIBE = 10 //取消订阅
	MQTT_UNSUBACK    = 11 //取消订阅应答
	MQTT_PINGREQ     = 12 //ping请求
	MQTT_PINGRESP    = 13 //ping响应
	MQTT_DISCONNECT  = 14 //断开连接
)

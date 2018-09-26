package goomem

var isUp bool = false

//启动程序
func StartUp() bool {
	if isUp {
		return false
	}
	//已经启动了。
	isUp = true

	return true
}

//获取值
func Get() {

}

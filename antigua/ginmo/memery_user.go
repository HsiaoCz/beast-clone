package main

var MemMap map[string]UserInfo

func InsertMap(id string, userInfo UserInfo) {
	MemMap = make(map[string]UserInfo)
	MemMap[id] = userInfo
}

func GetMemery(id string) (UserInfo, bool) {
	userInfo, ok := MemMap[id]
	return userInfo, ok
}

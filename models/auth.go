package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	if isPresent(auth) {
		return true
	}
	return false
}

func isPresent(obj interface{}) bool {
	//if a, ok := obj.(Auth); ok {
	//	return a.ID > 0
	//}
	//
	//flag := false
	//if flag == false {
	//
	//}
	//
	//if a, ok := obj.([]Auth); ok{
	//	return len(a) > 0
	//}
	//
	//return false

	switch a := obj.(type) {
	case Auth:
		return a.ID > 0
	}
	return false
}
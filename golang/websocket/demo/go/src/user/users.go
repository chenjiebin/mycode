package user

//用户群对象
var UserList Users

//用户群
type Users struct {
	Users map[string]*User
}

//增加用户
func (this *Users) Add(u *User) {
	if this.IsExist(u.Name) {
		return
	}
	this.Users[u.Name] = u
}

//检测用户是否存在
func (this *Users) IsExist(Name string) bool {
	_, ok := this.Users[Name]
	return ok
}

//移除用户
func (this *Users) RemoveUser(u *User) {
	delete(this.Users, u.Name)
}

func init() {
	UserList = Users{
		Users: make(map[string]*User),
	}
}

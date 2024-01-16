package datatypes

type SVCS map[string]string

type FileInfo struct {
	FileNames []string
}
type User struct {
	UserName string
	FileInfo
}

func CreateUser() *User {
	return new(User)
}
func (u *User) ConfigAction(userName string) *User {
	if u.UserName == userName {
		//  create a temp User object with only UserName field
		return &User{UserName: userName}
	}
	u.UserName = userName
	return u
}

func (u *User) AddAction(fileName string) {
	u.FileNames = append(u.FileNames, fileName)
}

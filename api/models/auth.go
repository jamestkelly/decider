package models

/*
User
Interface to represent basic user authentication objects.
*/
type User struct {
	ID        string   `json:"id"`
	Name 	  string   `json:"name"`
	Email     string   `json:"email"`
	AvatarURL string   `json:"avatarURL"`
	Password  string   `json:"-"`
	ServerIDs []string `json:"serverIDs"`
}

func (u *User) editName(newName string) {
	u.Name = newName
}

func (u *User) editEmail(newEmail string) {
	u.Email = newEmail
}

func (u *User) editAvatar(newAvatarURL string) {
	u.AvatarURL = newAvatarURL
}

func (u *User) editPassword(newPassword string) {
	u.Password = newPassword
}

func (u *User) joinServer(serverID string) {
	u.ServerIDs = append(u.ServerIDs, serverID)
}

func (u *User) leaveServer(serverID string) {
	for i, id := range u.ServerIDs {
		if id == serverID {
			u.ServerIDs = append(u.ServerIDs[:i], u.ServerIDs[i+1:]...)
			break
		}
	}
}

/*
RegisterResponse
Interface to represent user registration response structure.
*/
type RegisterResponse struct {
	Message string `json:"message"`
}

/*
LoginResponse
Interface to represent user login response structure.
*/
type LoginResponse struct {
	Token string `json:"token"`
}

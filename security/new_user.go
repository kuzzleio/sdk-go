package security

func (s *Security) NewUser() *User {
	u := &User{
		Security: s,
	}
	return u
}

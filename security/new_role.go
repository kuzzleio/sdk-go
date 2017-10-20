package security

func (s *Security) NewRole() *Role {
	r := &Role{
		Security: s,
	}

	return r
}

package security

func (s *Security) NewRole() *Role {
	r := &Role{
		Kuzzle: s.Kuzzle,
	}

	return r
}

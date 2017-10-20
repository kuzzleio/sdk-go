package security

func (s *Security) NewProfile () *Profile {
	return &Profile{
		Security: s,
	}
}

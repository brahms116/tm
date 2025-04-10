package tm

type UserErr string

func (e UserErr) Error() string {
	return string(e)
}

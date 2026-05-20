package user

func ValidateRole(role string) bool {
	switch role {
	case "owner", "tenant":
		return true
	default:
		return false
	}
}

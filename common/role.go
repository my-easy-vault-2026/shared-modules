package common

type Role int // 角色
const (
	ROLE_GUEST          Role = 1
	ROLE_USER           Role = 2
	ROLE_ADMIN          Role = 3
	ROLE_MERCHANT       Role = 4
	ROLE_MERCHANT_ADMIN Role = 5
	ROLE_MERCHANT_USER  Role = 6
	ROLE_SYSTEM         Role = 9
)

func (r Role) FromString(s string) Role {
	switch s {
	case "guest":
		return ROLE_GUEST
	case "user":
		return ROLE_USER
	case "admin":
		return ROLE_ADMIN
	case "merchant":
		return ROLE_MERCHANT
	case "merchant_admin":
		return ROLE_MERCHANT_ADMIN
	case "merchant_user":
		return ROLE_MERCHANT_USER
	case "system":
		return ROLE_SYSTEM
	}
	return 0
}

func (r Role) String() string {
	switch r {
	case ROLE_GUEST:
		return "guest"
	case ROLE_USER:
		return "user"
	case ROLE_ADMIN:
		return "admin"
	case ROLE_MERCHANT:
		return "merchant"
	case ROLE_MERCHANT_ADMIN:
		return "merchant_admin"
	case ROLE_MERCHANT_USER:
		return "merchant_user"
	case ROLE_SYSTEM:
		return "system"
	}
	return ""
}

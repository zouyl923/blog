package cKey

var (
	AdminTokenTtl = 2 * 60 * 60
	AdminTokenKey = "admin:token:"

	AdminRefreshTokenTtl = 30 * 60 * 60
	AdminRefreshTokenKey = "admin:refresh_token:"

	AdminPermissionTtl = 24 * 60 * 60
	AdminPermissionKey = "admin:role:permission:"
)

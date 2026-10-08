package usecases

type AuthUseCase interface {
	Login(email, password string) (accessToken, refreshToken string, err error)
	RefreshToken(token string) (accessToken, refreshToken string, err error)
}

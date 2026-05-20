package usecases

type AuthUseCase interface {
	Login(email, password string) (accessToken, refreshToken string, err error)
	RefreshToken(refreshToken string) (accessToken string, err error)
}

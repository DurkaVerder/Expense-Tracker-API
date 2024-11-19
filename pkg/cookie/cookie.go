package cookie

import (
	"Expense-Tracker-API/pkg/jwt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type Cookie interface {
	SaveJWTInCookie(ctx echo.Context, userId int) error
	GetUserIdByCookie(ctx echo.Context) (int, error)
}

type CookieManager struct {
	store *sessions.CookieStore
}

func NewCookie() Cookie {
	return CookieManager{sessions.NewCookieStore([]byte("key"))}
}

func (c CookieManager) SaveJWTInCookie(ctx echo.Context, userId int) error {
	token, err := jwt.GenerateJWT(userId)
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt-token"
	cookie.Value = token
	cookie.HttpOnly = true                          // Запрет доступа к cookie через JavaScript
	cookie.Expires = time.Now().Add(72 * time.Hour) // Устанавливаем время истечения cookie
	cookie.Path = "/"                               // Cookie будет доступно для всех путей сайта
	cookie.SameSite = http.SameSiteLaxMode          // Устанавливаем политику SameSite для предотвращения CSRF
	cookie.Secure = false                           // Используется только через HTTPS

	ctx.SetCookie(cookie)

	return nil
}

func (c CookieManager) getJWTFromCookie(ctx echo.Context) (string, error) {
	Cookie, err := ctx.Cookie("jwt-token")
	if err != nil {
		return "", err
	}

	return Cookie.Value, nil
}

func (c CookieManager) GetUserIdByCookie(ctx echo.Context) (int, error) {
	token, err := c.getJWTFromCookie(ctx)
	if err != nil {
		return -1, err
	}

	claims, err := jwt.ValidateJWT(token)
	if err != nil {
		return -1, err
	}
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return -1, err
	}
	userId := int(userIdFloat)

	return userId, nil
}

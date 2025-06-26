package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"re-crm/dtos"
	"re-crm/errs"
	"re-crm/utils"
)

func (app *App) HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from /"))
}

func (app *App) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from /api/dashboard"))
}
func (app *App) HandleLogin(w http.ResponseWriter, r *http.Request) {
	dto := dtos.LoginDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("🚀 got login request with nick %s and pass %s\n", dto.Nick, dto.Pass)

	token, role, err := app.AuthSvc.Login(r.Context(), dto)
	if err != nil {
		switch err {
		case errs.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case errs.InvalidCreds:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	cookie := &http.Cookie{
		Name:     "authJWT",
		Value:    token,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure:   false, //true in staging | prod
	}
	http.SetCookie(w, cookie)
	app.writers.http.Json(w, http.StatusOK, utils.Object{
		"role":       role,
		"isLoggedIn": true,
	})
	fmt.Println("the set cookie has value: ", token)

}
func (app *App) HandleChat(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from /api/dashboard"))
}
func (app *App) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "authJWT",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false, //prod = true
	})
}

func (app *App) HandleMe(w http.ResponseWriter, r *http.Request) {

	auth, ok := r.Context().Value(AuthCtxKey).(authCtx)
	if !ok {
		fmt.Println("not ok.. auth is: ", auth)
		app.writers.http.Error(w, http.StatusInternalServerError, "auth was not in r.Context")
		return
	}
	obj := utils.Object{
		"role":       auth.Role,
		"isLoggedIn": true,
	}
	fmt.Printf("☠️ sent to user obj %+v. and role is: %s\n", obj, auth.Role)
	app.writers.http.Json(w, http.StatusOK, obj)
}
func (app *App) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from /api/dashboard"))
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alejandroas97/online-streaming/models"
	"github.com/alejandroas97/online-streaming/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret_key")

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users *[]models.User = repository.GetAllUsers()
	json.NewEncoder(w).Encode(&users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser, err := repository.CreateUser(&user)

	if createdUser == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	fmt.Println("LOGIN")
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	user := loginUser
	userLogged := repository.FindUser(&user)
	if userLogged == nil {
		// Si el usuario no existe, las credenciales no son válidas
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogged.Password), []byte(loginUser.Password))
	if err != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver el token JWT como respuesta
	w.Write([]byte(tokenString))

}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Se requiere un token de autenticación", http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[len("Bearer "):]

		// Parsear el token JWT
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Token de autenticación inválido", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Token de autenticación inválido", http.StatusBadRequest)
			return
		}
		if !token.Valid {
			http.Error(w, "Token de autenticación inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

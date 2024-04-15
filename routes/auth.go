package routes

// import (
//     "encoding/json"
//     "net/http"
//     "time"

//     "github.com/dgrijalva/jwt-go"
// )

// var jwtKey = []byte("your-secret-key") // Cambia esto por tu propia clave secreta

// type Claims struct {
//     Username string `json:"username"`
//     jwt.StandardClaims
// }

// func Login(username, password string) (string, error) {
//     // Aquí deberías validar el nombre de usuario y la contraseña
//     // y verificar si son válidos en tu sistema de autenticación

//     // Simplemente generamos un token JWT de ejemplo por ahora
//     expirationTime := time.Now().Add(24 * time.Hour)
//     claims := &Claims{
//         Username: username,
//         StandardClaims: jwt.StandardClaims{
//             ExpiresAt: expirationTime.Unix(),
//         },
//     }
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     tokenString, err := token.SignedString(jwtKey)
//     if err != nil {
//         return "", err
//     }
//     return tokenString, nil
// }

// func Middleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         tokenString := r.Header.Get("Authorization")
//         if tokenString == "" {
//             http.Error(w, "Authorization token missing", http.StatusUnauthorized)
//             return
//         }

//         claims := &Claims{}
//         token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//             return jwtKey, nil
//         })
//         if err != nil {
//             http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
//             return
//         }
//         if !token.Valid {
//             http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
//             return
//         }

//         // Si el token es válido, continuar con la siguiente manejador
//         next.ServeHTTP(w, r)
//     })
// }

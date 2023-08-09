package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	port = 8082
	// Test: pass1234
	exampleHashValue = "b66dd5a7a689f88e302ab2ae4a9567f9c7572c18e520b3bf712bb2630b3931a503d647baedf48df470006312d07984216578b60526e5ee6137ef1fd215190a0c"
)

var (
	secretKey []byte
)

// Estrutura do usuário
type DataLogin struct {
	UserId string `json:"userId"`
	Pwd    string `json:"pass"`
}

// Middleware para autenticar o token JWT
func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}

func applyCORS(w http.ResponseWriter, r *http.Request, m string) bool {
	fmt.Println("CORS")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		log.Println("OPTIONS")
		w.Header().Set("Access-Control-Allow-Methods", m)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return true
	}
	return false
}

// Rota para autenticação de login
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	if applyCORS(w, r, "POST") {
		return
	}
	var user DataLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("nameid:", user.UserId)
	fmt.Println("pwdHashed:", user.Pwd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.UserId != "Junio" || user.Pwd != exampleHashValue {
		http.Error(w, "Invalid name or password", http.StatusUnauthorized)
		return
	}

	// Gera o token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserId,
		"pass":   user.Pwd,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"token": tokenString}
	fmt.Println("token:", tokenString)
	json.NewEncoder(w).Encode(resp)
}

// Rota protegida com autenticação JWT
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protected")
	if applyCORS(w, r, "GET") {
		return
	}
	w.Write([]byte("protected route.\nIt works!"))
}

func main() {
	// Cria um slice de bytes para armazenar a chave secreta
	secretKey = make([]byte, 32) // Gera uma chave secreta de 32 bytes (256 bits)

	// Preenche o slice de bytes com valores aleatórios
	_, err := rand.Read(secretKey)
	if err != nil {
		fmt.Println("Erro ao gerar a chave secreta:", err)
		return
	}

	fmt.Println("initializing go server")

	// Inicialização do roteador do Gorilla Mux
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST", "OPTIONS")
	router.HandleFunc("/protected", authenticate(protectedEndpoint)).Methods("GET", "OPTIONS")

	fmt.Printf("listening port %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/big"
	"net/http"
)

// PasswordRequest representa a requisição para gerar senha
type PasswordRequest struct {
	Length     int  `json:"length"`
	Uppercase  bool `json:"uppercase"`
	Lowercase  bool `json:"lowercase"`
	Numbers    bool `json:"numbers"`
	Symbols    bool `json:"symbols"`
}

// PasswordResponse representa a resposta com a senha gerada
type PasswordResponse struct {
	Password string `json:"password"`
	Error    string `json:"error,omitempty"`
}

const (
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	numberChars    = "0123456789"
	symbolChars    = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func main() {
	// Servir arquivos estáticos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Servir favicon.ico da raiz
	http.HandleFunc("/favicon.ico", faviconHandler)
	
	// Rota principal
	http.HandleFunc("/", homeHandler)
	
	// API para gerar senha
	http.HandleFunc("/generate", generatePasswordHandler)
	
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Criador Strings/Senhas</title>
	<link rel="icon" type="image/x-icon" href="/favicon.ico">
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <div class="container">
        <h1>Gerador de Senhas</h1>
        
        <form id="passwordForm">
            <div class="form-group">
                <label for="length">Comprimento da senha:</label>
                <select id="length" name="length">
                    <option value="10">10 caracteres</option>
                    <option value="15" selected>15 caracteres</option>
                    <option value="20">20 caracteres</option>
                </select>
            </div>
            
            <div class="form-group">
                <label>Tipos de caracteres:</label>
                <div class="checkbox-group">
                    <label>
                        <input type="checkbox" id="uppercase" name="uppercase" checked>
                        Maiúsculas (A-Z)
                    </label>
                    <label>
                        <input type="checkbox" id="lowercase" name="lowercase" checked>
                        Minúsculas (a-z)
                    </label>
                    <label>
                        <input type="checkbox" id="numbers" name="numbers" checked>
                        Números (0-9)
                    </label>
                    <label>
                        <input type="checkbox" id="symbols" name="symbols">
                        Símbolos (!@#$%^&*)
                    </label>
                </div>
            </div>
            
            <button type="submit">Gerar Senha</button>
        </form>
        
        <div id="result" class="result" style="display: none;">
            <h3>Senha Gerada:</h3>
            <div class="password-display">
                <input type="text" id="generatedPassword" readonly>
                <button type="button" id="copyBtn">Copiar</button>
            </div>
        </div>
        
        <div id="error" class="error" style="display: none;"></div>
    </div>
    
    <script src="/static/script.js"></script>
</body>
</html>`
	
	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	t.Execute(w, nil)
}

func generatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	var req PasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	
	// Validar requisição
	if req.Length != 10 && req.Length != 15 && req.Length != 20 {
		sendErrorResponse(w, "Comprimento deve ser 10, 15 ou 20 caracteres")
		return
	}
	
	if !req.Uppercase && !req.Lowercase && !req.Numbers && !req.Symbols {
		sendErrorResponse(w, "Pelo menos um tipo de caractere deve ser selecionado")
		return
	}
	
	// Gerar senha
	password, err := generatePassword(req)
	if err != nil {
		sendErrorResponse(w, "Erro ao gerar senha: "+err.Error())
		return
	}
	
	// Enviar resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PasswordResponse{Password: password})
}

func generatePassword(req PasswordRequest) (string, error) {
	var charset string
	
	if req.Uppercase {
		charset += uppercaseChars
	}
	if req.Lowercase {
		charset += lowercaseChars
	}
	if req.Numbers {
		charset += numberChars
	}
	if req.Symbols {
		charset += symbolChars
	}
	
	if len(charset) == 0 {
		return "", fmt.Errorf("nenhum tipo de caractere selecionado")
	}
	
	password := make([]byte, req.Length)
	
	for i := 0; i < req.Length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}
	
	return string(password), nil
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(PasswordResponse{Error: message})
}
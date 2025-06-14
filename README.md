### Gerador de Senhas em Go

Uma aplicação web simples em Go para gerar senhas seguras com diferentes opções de configuração.

#### Funcionalidades

- Geração de senhas com 10, 15 ou 20 caracteres
- Opções de caracteres:
  - Maiúsculas (A-Z)
  - Minúsculas (a-z)
  - Números (0-9)
  - Símbolos (!@#$%^&*()_+-=[]{}|;:,.<>?)
- Interface web simples e responsiva
- Botão para copiar senha gerada
- Validação de entrada
- Geração criptograficamente segura usando `crypto/rand`

#### Estrutura do Projeto

```
password-generator/
├── main.go           # Servidor principal
├── go.mod           # Módulo Go
├── README.md        # Documentação
└── static/          # Arquivos estáticos
    ├── style.css    # Estilos CSS
    └── script.js    # JavaScript frontend
```

### Como Executar

1. **Pré-requisitos:**
   - Go 1.21+ instalado

2. **Clonar/criar o projeto:**
   ```bash
   mkdir password-generator
   cd password-generator
   ```

3. **Criar a estrutura de pastas:**
   ```bash
   mkdir static
   ```

4. **Adicionar os arquivos** (main.go, go.mod, static/style.css, static/script.js)

5. **Executar a aplicação:**
   ```bash
   go run main.go
   ```

6. **Acessar no browser:**
   ```
   http://localhost:8080
   ```

### API Endpoints

#### `GET /`
Retorna a página principal da aplicação.

#### `POST /generate`
Gera uma nova senha baseada nos parâmetros fornecidos.

**Request Body:**
```json
{
  "length": 15,
  "uppercase": true,
  "lowercase": true,
  "numbers": true,
  "symbols": false
}
```

**Response:**
```json
{
  "password": "Kj8mN2pL9qR4tY7"
}
```

**Error Response:**
```json
{
  "error": "Pelo menos um tipo de caractere deve ser selecionado"
}
```

### Segurança

- Utiliza `crypto/rand` para geração criptograficamente segura
- Não armazena senhas geradas
- Validação tanto no frontend quanto no backend

### Tecnologias Utilizadas

- **Backend:** Go (net/http, crypto/rand, html/template)
- **Frontend:** HTML5, CSS3, JavaScript (ES6+)
- **Arquitetura:** SPA com API REST
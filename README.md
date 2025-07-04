## Pass Generator

![Pipeline](https://github.com/dellabeneta/pass-generator/actions/workflows/main.yaml/badge.svg)
![GHCR version](https://img.shields.io/badge/ContainerVersion-v11-blue)

Um gerador de senhas moderno, seguro e fácil de usar, desenvolvido em Go. Gere senhas fortes com uma interface web intuitiva e responsiva.

### Funcionalidades

- Geração de senhas com comprimentos personalizáveis (10, 15 ou 20 caracteres)
- Opções flexíveis de caracteres:
  - Letras maiúsculas (A-Z)
  - Letras minúsculas (a-z)
  - Números (0-9)
  - Símbolos especiais (!@#$%^&*)
- Interface responsiva para desktop e mobile
- Cópia rápida para área de transferência
- Geração criptograficamente segura
- Alta disponibilidade com Kubernetes

### Começando

**Pré-requisitos**
- Go 1.24.3 ou superior
- Docker (opcional)
- Kubernetes/k3s (opcional)

**Instalação Local**
```bash
# Clone o repositório
git clone https://github.com/seu-usuario/pass-generator.git
cd pass-generator

# Execute a aplicação
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`

**Usando Docker**
```bash
# Construa a imagem
docker build -t pass-generator .

# Execute o container
docker run -p 8080:8080 pass-generator
```

**Deploy no Kubernetes**
```bash
# Crie o namespace
kubectl apply -f k3s/namespace.yaml

# Deploy da aplicação
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
```

### Tecnologias

- **Backend**: Go 1.24.3
- **Frontend**: HTML5, CSS3, JavaScript
- **Container**: Docker
- **Orquestração**: Kubernetes/k3s
- **Segurança**: crypto/rand para geração segura

### Configuração

O serviço pode ser configurado através das seguintes variáveis de ambiente:

| Variável | Descrição | Padrão |
|----------|-----------|---------|
| PORT | Porta do servidor | 8080 |

### Estrutura do Projeto
```
della@ubuntu:~/projetos/pass-generator$ tree
.
├── Dockerfile
├── go.mod
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── LICENSE
├── main.go
├── nuke.sh
├── pass.png
├── README.md
└── static
    ├── favicon.ico
    ├── script.js
    └── style.css

3 directories, 13 files
della@ubuntu:~/projetos/pass-generator$
```

### Contribuindo

1. Faça um Fork do projeto
2. Crie sua Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a Branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('passwordForm');
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error');
    const generatedPasswordInput = document.getElementById('generatedPassword');
    const copyBtn = document.getElementById('copyBtn');

    form.addEventListener('submit', function(e) {
        e.preventDefault();
        generatePassword();
    });

    copyBtn.addEventListener('click', function() {
        copyToClipboard();
    });

    function generatePassword() {
        // Ocultar resultados anteriores
        hideResults();

        // Coletar dados do formulário
        const length = parseInt(document.getElementById('length').value);
        const uppercase = document.getElementById('uppercase').checked;
        const lowercase = document.getElementById('lowercase').checked;
        const numbers = document.getElementById('numbers').checked;
        const symbols = document.getElementById('symbols').checked;

        // Validar se pelo menos uma opção foi selecionada
        if (!uppercase && !lowercase && !numbers && !symbols) {
            showError('Selecione pelo menos um tipo de caractere!');
            return;
        }

        // Preparar dados para envio
        const requestData = {
            length: length,
            uppercase: uppercase,
            lowercase: lowercase,
            numbers: numbers,
            symbols: symbols
        };

        // Fazer requisição para o servidor
        fetch('/generate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                showError(data.error);
            } else {
                showResult(data.password);
            }
        })
        .catch(error => {
            showError('Erro ao gerar senha: ' + error.message);
        });
    }

    function showResult(password) {
        hideError();
        generatedPasswordInput.value = password;
        resultDiv.style.display = 'block';
    }

    function showError(message) {
        hideResults();
        errorDiv.textContent = message;
        errorDiv.style.display = 'block';
    }

    function hideResults() {
        resultDiv.style.display = 'none';
    }

    function hideError() {
        errorDiv.style.display = 'none';
    }

    function copyToClipboard() {
        generatedPasswordInput.select();
        generatedPasswordInput.setSelectionRange(0, 99999); // Para dispositivos móveis

        try {
            document.execCommand('copy');
            
            // Feedback visual
            const originalText = copyBtn.textContent;
            copyBtn.textContent = 'Copiado!';
            copyBtn.style.backgroundColor = '#218838';
            
            setTimeout(() => {
                copyBtn.textContent = originalText;
                copyBtn.style.backgroundColor = '#28a745';
            }, 2000);
            
        } catch (err) {
            console.error('Erro ao copiar: ', err);
            alert('Erro ao copiar a senha');
        }
    }
});
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Estrutura para fazer parse do JSON das credenciais AWS
type AWSCredentials struct {
	Credentials struct {
		AccessKeyId     string `json:"AccessKeyId"`
		SecretAccessKey string `json:"SecretAccessKey"`
		SessionToken    string `json:"SessionToken"`
		Expiration      string `json:"Expiration"`
	} `json:"Credentials"`
	AssumedRoleUser struct {
		AssumedRoleId string `json:"AssumedRoleId"`
		Arn           string `json:"Arn"`
	} `json:"AssumedRoleUser"`
}

func main() {
	var awsSessionToken, awsSecretAccessKey, awsAccessKeyID, awsRegion string

	// Verifica se foi passado um JSON como argumento
	if len(os.Args) > 1 {
		// Parse do JSON das credenciais
		var creds AWSCredentials
		err := json.Unmarshal([]byte(os.Args[1]), &creds)
		if err != nil {
			fmt.Printf("Erro ao fazer parse do JSON: %v\n", err)
			return
		}

		// Usa as credenciais do JSON
		awsSessionToken = creds.Credentials.SessionToken
		awsSecretAccessKey = creds.Credentials.SecretAccessKey
		awsAccessKeyID = creds.Credentials.AccessKeyId
		awsRegion = os.Getenv("AWS_REGION") // Região ainda vem da variável de ambiente
		if awsRegion == "" {
			awsRegion = "us-east-1" // Valor padrão
		}

		fmt.Println("Usando credenciais do JSON fornecido")
	} else {
		// Obtém as variáveis de ambiente (comportamento original)
		awsSessionToken = os.Getenv("AWS_SESSION_TOKEN")
		awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		awsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
		awsRegion = os.Getenv("AWS_REGION")

		fmt.Println("Usando credenciais das variáveis de ambiente")
	}

	// Validação das credenciais
	if awsAccessKeyID == "" || awsSecretAccessKey == "" || awsSessionToken == "" {
		fmt.Println("Erro: Credenciais AWS incompletas")
		fmt.Println("Certifique-se de fornecer um JSON válido ou definir as variáveis de ambiente:")
		fmt.Println("- AWS_ACCESS_KEY_ID")
		fmt.Println("- AWS_SECRET_ACCESS_KEY")
		fmt.Println("- AWS_SESSION_TOKEN")
		return
	}

	// Caminho do arquivo de credenciais
	credentialsPath := fmt.Sprintf("%s/.aws/credentials", os.Getenv("HOME"))

	// Cria ou sobrescreve o arquivo de credenciais
	file, err := os.Create(credentialsPath)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo de credenciais: %v\n", err)
		return
	}
	defer file.Close()

	// Escreve as credenciais no arquivo
	_, err = file.WriteString(fmt.Sprintf(`[default]
aws_session_token=%s
aws_secret_access_key=%s
aws_access_key_id=%s
`, awsSessionToken, awsSecretAccessKey, awsAccessKeyID))
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo de credenciais: %v\n", err)
		return
	}

	// Exporta as variáveis de ambiente
	os.Setenv("AWS_REGION", awsRegion)
	os.Setenv("AWS_SESSION_TOKEN", awsSessionToken)
	os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey)
	os.Setenv("AWS_ACCESS_KEY_ID", awsAccessKeyID)
	os.Setenv("AWS_DEFAULT_REGION", awsRegion)

	fmt.Println("Credenciais configuradas com sucesso!")
}

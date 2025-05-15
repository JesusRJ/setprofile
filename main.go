package main

import (
    "fmt"
    "os"
)

func main() {
    // Obtém as variáveis de ambiente
    awsSessionToken := os.Getenv("AWS_SESSION_TOKEN")
    awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
    awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    awsRegion := os.Getenv("AWS_REGION")

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
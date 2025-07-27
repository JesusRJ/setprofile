# setprofile

AWS Set Profile write AWS credentials values to ~/.aws/credentials file and export the right environment variables.

## Uso

O programa pode ser usado de duas formas:

### 1. Com JSON das credenciais AWS (Recomendado)

Execute o programa passando o JSON das credenciais como argumento:

```bash
./setprofile '{"Credentials":{"AccessKeyId":"YOUR_ACCESS_KEY","SecretAccessKey":"YOUR_SECRET_KEY","SessionToken":"YOUR_SESSION_TOKEN","Expiration":"2025-07-27T19:57:14+00:00"},"AssumedRoleUser":{"AssumedRoleId":"YOUR_ROLE_ID","Arn":"YOUR_ARN"}}'
```

Ou usando um arquivo JSON:

```bash
./setprofile "$(cat example.json)"
```

### 2. Com variáveis de ambiente (Comportamento original)

Configure as variáveis de ambiente e execute o programa:

```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
export AWS_SESSION_TOKEN="your-session-token"
export AWS_REGION="us-east-1"
./setprofile
```

## Formato do JSON

O JSON deve seguir o formato retornado pelo comando `aws sts assume-role`:

```json
{
  "Credentials": {
    "AccessKeyId": "ASIATKADFAFAR",
    "SecretAccessKey": "Ui7/JONW69ADFAFAFDg",
    "SessionToken": "IQoJb3JpZ2luX2VjEFMa...",
    "Expiration": "2025-07-27T19:57:14+00:00"
  },
  "AssumedRoleUser": {
    "AssumedRoleId": "AROA6ETFJ7MA5XWJY6EBC:AWSCLI-Session",
    "Arn": "arn:aws:sts::971948817153:assumed-role/terraform-admin/AWSCLI-Session"
  }
}
```

## Compilação

```bash
go build -o setprofile main.go
```
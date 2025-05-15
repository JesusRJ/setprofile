#!/usr/bin/env bash

# Cria ou sobrescreve o arquivo de credenciais AWS
cat > "$HOME/.aws/credentials" <<EOF
[default]
aws_session_token=$AWS_SESSION_TOKEN
aws_secret_access_key=$AWS_SECRET_ACCESS_KEY
aws_access_key_id=$AWS_ACCESS_KEY_ID
EOF

# Exporta as variáveis de ambiente necessárias
export AWS_REGION=$AWS_REGION
export AWS_SESSION_TOKEN=$AWS_SESSION_TOKEN
export AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
export AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
export AWS_DEFAULT_REGION=$AWS_REGION

# Opcional: Configura a região padrão no AWS CLI
# aws configure set default.region $AWS_REGION
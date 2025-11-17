#!/usr/bin/env bash

# Load credentials
if [ ! -f .env.ec2-credentials ]; then
    echo "Error: .env.ec2-credentials file is missing"
    exit 1
fi

SSH_KEY=$(grep SSH_KEY .env.ec2-credentials | cut -d '=' -f 2)
AWS_USER=$(grep AWS_USER .env.ec2-credentials | cut -d '=' -f 2)
AWS_HOST=$(grep AWS_HOST .env.ec2-credentials | cut -d '=' -f 2)

if [ -z "$SSH_KEY" ] || [ -z "$AWS_USER" ] || [ -z "$AWS_HOST" ]; then
    echo "Error: Missing required variables in .env.ec2-credentials [SSH_KEY, AWS_USER, AWS_HOST]"
    exit 1
fi

echo "Initializing deployment on ${AWS_USER}@${AWS_HOST}..."

# SSH and run initialization commands on EC2
ssh -i "$SSH_KEY" "${AWS_USER}@${AWS_HOST}" << EOF
    sudo apt update -y && sudo apt upgrade -y
    sudo apt install -y golang-go git mysql-client
    sudo mkdir -p /apps/core
    sudo chown -R ${AWS_USER}:${AWS_USER} /apps/core
    sudo chmod -R 755 /apps/core
    echo "Initialization complete on EC2."
EOF

if [ $? -ne 0 ]; then
    echo "Error: Initialization failed"
    exit 1
fi

echo "Done!"
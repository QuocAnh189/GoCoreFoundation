#!/usr/bin/env bash

# Check if the .env.ec2-credentials file exists
if [ ! -f .env.ec2-credentials ]; then
    echo "Error: .env.ec2-credentials file is missing"
    exit 1
fi


SSH_KEY=$(cat .env.ec2-credentials | grep SSH_KEY | cut -d '=' -f 2)
USER=$(cat .env.ec2-credentials | grep AWS_USER | cut -d '=' -f 2)
HOST=$(cat .env.ec2-credentials | grep AWS_HOST | cut -d '=' -f 2)

if [ -z "$SSH_KEY" ] || [ -z "$USER" ] || [ -z "$HOST" ]; then
    echo "Error: .env.ec2-credentials file is missing one or more required variables [SSH_KEY, USER, HOST]"
    exit 1
fi

# 0. Define constants

SVR_DIR="/var/www/core_foundation_svr"

# 1. Create Go server web directory:

ssh -i "$SSH_KEY" "$USER@$HOST" \
"
cd /var/www/ && \
sudo mkdir core_foundation_svr && \
sudo chgrp -R monex core_foundation_svr/ && \
sudo chmod 775 -R core_foundation_svr/
"

# 2. Upload JWT handshake keys:

rsync -e "ssh -i $SSH_KEY" resources/certs/*.pem "$USER@$HOST":"$SVR_DIR/keys"

# 3. Create .env file:

# First, upload the file
scp -i "$SSH_KEY" ./.env "$USER@$HOST":"$SVR_DIR/.env"
# Then, open it in vim if you need to make further edits
ssh -t -i "$SSH_KEY" "$USER@$HOST" "sudo vim $SVR_DIR/.env"

# 4. Create 'ez' utility command

cat << 'EOF' | ssh -i "$SSH_KEY" "$USER@$HOST" "sudo tee /usr/local/bin/ez >/dev/null && sudo chmod +x /usr/local/bin/ez"
eval "$(ssh-agent)"
# echo | ssh-add "./scripts/aws-bot-vps.pem"
ssh -tt -i "scripts/aws-bot-vps.pem" -o "StrictHostKeyChecking=no" ubuntu@ec2-13-48-30-65.eu-north-1.compute.amazonaws.com << EOF

docker stop play
docker rmi sochisic/play-with-lefthook
docker run -d -p 3000:3000 --name play --rm sochisic/play-with-lefthook
EOF

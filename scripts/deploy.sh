# eval "$(ssh-agent)"
# ssh -tt -i "scripts/aws-bot-vps.pem" -o "StrictHostKeyChecking=no" ubuntu@ec2-13-48-30-65.eu-north-1.compute.amazonaws.com << EOF
# docker stop play
# docker rm play
# docker rmi --force sochisic/play-with-lefthook
# docker run -d -p 3000:3000 --name play --restart always sochisic/play-with-lefthook
# EOF
eval "$(ssh-agent)"
yes '' | ssh -tt -i "scripts/aws-bot-vps.pem" -o "StrictHostKeyChecking=no" ubuntu@ec2-13-48-30-65.eu-north-1.compute.amazonaws.com "
docker stop play &&
docker rm play &&
docker rmi --force sochisic/play-with-lefthook &&
docker run -d -p 3000:3000 --name play --restart always sochisic/play-with-lefthook
"

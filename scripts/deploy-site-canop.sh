
# déploiement du site web
rsync -avz --del --stats --exclude="deploy.sh" /home/dys/dev/Chrall/web/* dys@canop.org:/var/www/canop/chrall

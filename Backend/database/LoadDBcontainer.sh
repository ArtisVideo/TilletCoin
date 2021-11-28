if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi
docker rm db_container
gunzip -c db_container.tar.gz | docker load
docker run -d --name=db_container db_container
docker stop db_container
echo "Contaier added / updatded from image please start using docker start db_container"
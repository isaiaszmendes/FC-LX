# Docker

"A imagem é como se fosse a planta da casa
O container é como se fosse a casa
E eu posso ter várias casas iguais baseados em UMA planta (imagem)"


#### build image
docker build -t image-nginx .    

#### run container from image
docker run -p 3000:80 image-nginx
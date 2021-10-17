# SteganographyChallenge

Resolução do desafio da criação de uma REST API que consiste em:

/upload: Uma requisição POST que recebe um multipart/form-data com uma
imagem bitmap e retorna o nome do arquivo armazenado em um diretório
temporário do servidor.

● /write-message-on-image: Uma requisição POST que recebe um
application/json com o caminho da imagem (resposta do /upload), aplica um
algoritmo de Esteganografia (será descrito melhor abaixo) retornando um JSON
informando o nome do novo arquivo.

● /get-image: Uma requisição GET que recebe na query o nome da imagem a ser
acessada e retorna o arquivo para download.

● /decode-message-from-image: Uma requisição GET que recebe na query o nome
da imagem a ser decodificada e retorna a mensagem escondida na imagem.

Lista de curl:
Favor alterar onde está descrito "(Inserir)"

/upload:
curl --location --request POST localhost:8080/post --form newFile=@"C:/imagens/(Inserir).bmp"

/write-message-on-image:
curl -i -X POST -H "Content-Type: application/json" -d  "{\"Path\":\"upload-(Inserir).bmp\",\"Message\":\"(Inserir)\"}" http://localhost:8080/write-message-on-image

/get-image:	
curl --location --request GET localhost:8080/get-image?FileName=Modified-(Inserir).bmp -o Modified-(Inserir).bmp

/decode-message-from-image:
curl --location --request GET localhost:8080/decode-message-from-image?FileName=Modified-(Inserir).bmp



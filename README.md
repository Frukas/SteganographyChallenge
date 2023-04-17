# SteganographyChallenge

Resolução do desafio da criação de uma REST API que recebe uma imagem bmp, insere uma mensagem secreta nos bits da imagem, e retorna a mensagem secreta escrita na imagem.
O desafio é separado em 4 chamadas:

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

_______________________________________________________________________________________________________________
#English description

Resolution of the challenge of creating a REST API that receives a bmp image, inserts a secret message in the bits of the image, and returns the secret message written in the image. The challenge is divided into 4 calls:

/upload: A POST request that receives a multipart/form-data with a bitmap image and returns the name of the file stored in a temporary directory on the server.

/write-message-on-image: A POST request that receives an application/json with the path of the image (response from /upload), applies a steganography algorithm (described below) returning a JSON informing the name of the new file.

/get-image: A GET request that receives in the query the name of the image to be accessed and returns the file for download.

/decode-message-from-image: A GET request that receives in the query the name of the image to be decoded and returns the hidden message in the image.

Curl list: Please replace "(Insert)" where described.

/upload: curl --location --request POST localhost:8080/post --form newFile=@"C:/images/(Insert).bmp"

/write-message-on-image: curl -i -X POST -H "Content-Type: application/json" -d "{"Path":"upload-(Insert).bmp","Message":"(Insert)"}" http://localhost:8080/write-message-on-image

/get-image: curl --location --request GET localhost:8080/get-image?FileName=Modified-(Insert).bmp -o Modified-(Insert).bmp

/decode-message-from-image: curl --location --request GET localhost:8080/decode-message-from-image?FileName=Modified-(Insert).bmp




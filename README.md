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

var  CloudmersiveVirusApiClient = require ( 'cloudmersive-virus-api-client' ); 
var defaultClient = CloudmersiveVirusApiClient . ApiClient . instance ; 

// Configurar autorização de chave de API: Apikey 
var  Apikey = defaultClient. authentications [ 'Apikey' ]; 
Apikey . apiKey = 'SUA CHAVE DE API' ; 



var apiInstance = new  CloudmersiveVirusApiClient . ScanApi (); 

var inputFile = Buffer . from (fs. readFileSync ( "C:\\temp\\inputfile" ). buffer ); // Arquivo | Arquivo de entrada para executar a operação. 

var opts = { 
  'allowExecutables' : true , // Boolean | Defina como false para bloquear arquivos executáveis ​​(código de programa) de serem permitidos no arquivo de entrada. O padrão é false (recomendado). 
  'allowInvalidFiles' : true , // Boolean | Defina como falso para bloquear arquivos inválidos, como um arquivo PDF que não é realmente um arquivo PDF válido ou um documento do Word que não é um documento do Word válido. O padrão é falso (recomendado). 
  'allowScripts' : true , // Boolean | Defina como falso para bloquear arquivos de script, como arquivos PHP, scripts Python e outros conteúdos maliciosos ou ameaças de segurança que podem ser incorporados no arquivo. Defina como verdadeiro para permitir esses tipos de arquivo. O padrão é falso (recomendado). 
  'allowPasswordProtectedFiles' : true , // Boolean | Defina como falso para bloquear arquivos protegidos por senha e criptografados, como arquivos zip e rar criptografados e outros arquivos que buscam contornar a varredura por senhas. Defina como verdadeiro para permitir esses tipos de arquivo. O padrão é falso (recomendado). 
  'allowMacros' : true , // Boolean | Defina como falso para bloquear macros e outras ameaças incorporadas em arquivos de documento, como macros incorporadas do Word, Excel e PowerPoint e outros arquivos que contêm ameaças de conteúdo incorporado. Defina como verdadeiro para permitir esses tipos de arquivo. O padrão é false (recomendado). 
  'allowXmlExternalEntities' : true , // Boolean | Defina como false para bloquear XML External Entities e outras ameaças incorporadas em arquivos XML, e outros arquivos que contenham ameaças de conteúdo incorporado. Defina como true para permitir esses tipos de arquivo. O padrão é false (recomendado).
  'allowInsecureDeserialization' : true , // Boolean | Defina como false para bloquear a Desserialização Insegura e outras ameaças incorporadas em JSON e outros arquivos de serialização de objetos, e outros arquivos que contêm ameaças de conteúdo incorporado. Defina como true para permitir esses tipos de arquivo. O padrão é false (recomendado). 
  'allowHtml' : true , // Boolean | Defina como false para bloquear a entrada HTML no arquivo de nível superior; HTML pode conter XSS, scripts, acessos a arquivos locais e outras ameaças. Defina como true para permitir esses tipos de arquivo. O padrão é false (recomendado) [para chaves de API criadas antes do lançamento deste recurso, o padrão é true para compatibilidade com versões anteriores]. 
  'restrictFileTypes' : "restrictFileTypes_example"  // String | Especifique um conjunto restrito de formatos de arquivo para permitir tão limpo quanto uma lista separada por vírgulas de formatos de arquivo, como .pdf, .docx, .png permitiria apenas arquivos de documentos PDF, PNG e Word. Todos os arquivos devem passar pela verificação de conteúdo em relação a esta lista de formatos de arquivo, caso contrário, o resultado será retornado como CleanResult=false. Defina o parâmetro restrictedFileTypes como nulo ou string vazia para desabilitar; o padrão é desabilitado.
 }; 

var callback = function ( error, data, response ) { 
  if (error) { 
    console . error (error); 
  } else { 
    console . log ( 'API chamada com sucesso. Dados retornados: ' + data); 
  } 
}; 
apiInstance. scanFileAdvanced (inputFile, opts, callback);
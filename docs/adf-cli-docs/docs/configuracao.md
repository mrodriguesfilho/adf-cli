# Configuração

Esta ferramenta utiliza os arquivos de configuração references.json e preferences.json para definir valores para variáveis de configuração. O local default é seguinte para cada plataforma:

- GNU/Linux e macOS: $HOME/.adf
- Windows: %USERPROFILE%/.adf

Preferences.json
- O arquivo contem uma coleção de urls para download das ferramentas necessárias para o adf. Separadas e indicas por versão

References.json
- O arquivo contem uma coleção de cada pacote instalado na maquina do usuário, separado por versão. Permite instalar as versões
do adf em diferentes locais no disco.

## Opções de configuração:

Version: versão do pacote com as ferramentas necessarias para usar o adf

DirectoryPath: Caminho no disco em que a versão especifica do adf está instalado

## Exemplo de arquivo de configuração

```JSON
{
  "InstalledBundles": [
    {
      "Version": "0.0.1", 
      "DirectoryPath": "/Users/mrf/.adf/0.0.1"
    }
  ]
}

```
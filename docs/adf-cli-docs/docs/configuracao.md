# Configuração

Esta ferramenta utiliza o arquivo de configuração .adf.toml para definir valores para variáveis de configuração. O local default é seguinte para cada plataforma:

- GNU/Linux e macOS: $HOME
- Windows: %USERPROFILE%

## Opções de configuração:

- web.port: int. Número de porta TCP a ser utilizada pelo ADF Web. Default: 8051
- web.workdir: string. Diretório de trabalho do ADF Web

- services.terminology.address: string. Endereço do servidor de dados. Default: nenhum
- services.terminology.port: int. Número de porta do servidor de dados. Default: nenhum
- services.terminology.certificate: string. Nome de arquivo de certificado digital para o serviço de terminologias. Default: nenhum

## Exemplo de arquivo de configuração

```toml
[web]
port = 8051
workdir = "~/adf"

[services]

[services.terminology]
address = "https://www.fhir-terminology.com/data/"
port = 433
certificate = "~/adf/terminologies-certificate.pem
```
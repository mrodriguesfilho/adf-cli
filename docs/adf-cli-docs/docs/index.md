# Bem vindo à documentação da ferramenta de CLI do Ambiente de Design FHIR (ADF)

A ferramenta de CLI do ADF é aplicação responsável por implementar as funcionalidades administrativas do Ambiente de Design FIHR (ADF) utilizando comandos de texto via shell.

Tais funcionalidades administrativas são as seguintes:

- Instalar, verificar (se a instalação está "correta") e remover.
- Iniciar e parar o ADF.
- Consultar o status da execução do ADF (monitorar registros de log).
- Configurar as opções de uso do ADF. A saber: servidores onde dados estão disponíveis (versões dos dados); (b) certificado digital (empregado para acesso ao servidor FHIR); (c) portas empregadas pelos serviços; (d) diretório de trabalho.
- Configurar as terminologias disponíveis para uso.
- Configurar os perfis disponíveis para uso.
- Configurar soluções disponíveis para uso.

## Plataformas suportadas

Até o momento, há suporte disponível para:

- [x] GNU/Linux + X86-64.
- [ ] macOS + X86-64.
- [ ] macOS + ARM64.
- [ ] Windows + X86-64.

O objetivo é que todas as combinações de plataformas citadas acima sejam suportadas no futuro.

## Exemplos

```shell-session

	$ adf install --version 0.0.1
	Instalando versão 0.0.1 do ADF Web...
	Versão 0.0.1 do ADF Web instalada com sucesso

	$ adf list
	Versões instaladas:
	ADF Web 0.0.1

	$ adf use 0.0.1
	Definida a versão 0.0.1 do ADF Web a ser utilizada

	$ adf list
	ADF Web 0.0.1 - em uso
```

## Opções globais

```
      --config string   config file (default is $HOME/.adf.toml)
  -h, --help            help for adf
  -t, --toggle          Help message for toggle
      --webPort int     Número de porta TCP do ADF Web (default 8050)
```
## Código-fonte

[https://github.com/karlosdaniel451/adf-cli](https://github.com/karlosdaniel451/adf-cli)

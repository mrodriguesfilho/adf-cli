## Instalação

Veja as plataformas suportadas [aqui](/#plataformas-suportadas).


## GNU/Linux

### A partir do código-fonte

1.  Requisitos

    1.  [Git](https://git-scm.com/download/linux)
    1.  [Go](https://go.dev/doc/install)
    2.  [Taskfile](https://taskfile.dev/installation/)

2.  Clone o repositório do código-fonte com o Git:
```shell
git clone https://github.com/karlosdaniel451/adf-cli
```

3.   Vá até o novo diretório criado:
```shell
cd adf-cli
```

4. Utilize o task para instalar a partir do código fonte:
```shell
task install_linux
```

Caso queira desinstalar, digite o seguinte no diretório do código-fonte do ADF:
```shell
task uninstall_linux
```
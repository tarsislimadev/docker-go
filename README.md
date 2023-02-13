# Docker Go Lang

[![github/actions/workflow/status](https://img.shields.io/github/actions/workflow/status/brtmvdl/docker-go/docker-push.yml)](https://img.shields.io/github/actions/workflow/status/brtmvdl/docker-go/docker-push.yml) [![github/license](https://img.shields.io/github/license/brtmvdl/docker-go)](https://img.shields.io/github/license/brtmvdl/docker-go) [![github/stars](https://img.shields.io/github/stars/brtmvdl/docker-go?style=social)](https://img.shields.io/github/stars/brtmvdl/antify?style=social)

Para compilaçao e entrega de projetos escritos em Go Lang

Veja mais em [hub.docker.com/r/tmvdl/go](https://hub.docker.com/r/tmvdl/go)

## Como usar

Instalar o [Docker](https://docs.docker.com/engine/install/).

### Em ambiente de desenvolvimento

Criar um arquivo `docker-compose.yaml` na raiz do projeto com a imagem [tmvdl/go](https://hub.docker.com/r/tmvdl/go).

```yaml
version: '3'

services:
  app:
    image: tmvdl/go
    volumes:
      - .:/app
```

Subir o container para a construção do build

```bash
docker-compose up --build
```

### Em ambiente de produção

Executar como container do Docker

```sh
docker run tmvdl/go
```

## License

[MIT](./LICENSE) 

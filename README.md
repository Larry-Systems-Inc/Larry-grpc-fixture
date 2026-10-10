# Larry gRPC Fixture

Servidor gRPC local para testar o Larry API Client.

Este repositório fornecerá cenários determinísticos para validar:

- RPC unary
- server streaming
- client streaming
- bidirectional streaming
- gRPC Reflection
- metadata
- erros
- timeouts
- trailers
- cancelamento

Este projeto é uma fixture de desenvolvimento. Ele não faz parte da aplicação de produção do Larry.

## Status

Configuração inicial do projeto. A implementação do servidor será adicionada em etapas.

## Relação com o Larry

A fixture roda localmente e permite que o Larry teste comportamentos gRPC sem depender de APIs externas ou serviços remotos.

## Cenários planejados

- Resposta unary normal
- Múltiplas respostas em server streaming
- Agregação de mensagens em client streaming
- Troca de mensagens em bidirectional streaming
- Descoberta por Reflection
- Metadata inválida
- Erros de autenticação
- Respostas com atraso
- Encerramento de stream pelo servidor
- Cancelamento pelo cliente

## Desenvolvimento

O servidor será construído incrementalmente para que cada etapa tenha um objetivo de teste claro no Larry.

## Licença

A definir junto com a primeira implementação.

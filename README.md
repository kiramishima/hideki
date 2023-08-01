# Hideki

Proyecto prueba para mejorar mi skill en Go despues de haber sido desarrollado por un mock interview.
A diferencia de Cleara que era para probar la arquitectura hexagonal y poder aplicarla a otros proyectos,
Hideki es una versión avanzada de Cleara con Dependency Injection, para ello se usa Uber FX, pensaba usar Google Wire, 
pero desde 2021 no han lanzado una versión nueva y pues Go ha cambiado.

## Librerias externas usadas en este proyecto:

- Chi Router
- Uber FX
- Uber Gomock
- Zap
- pq
- envconfig
- go-sqlmock
- testify

## Estructura

Se sigue la arquitectura hexagonal con clean architecture :-D

- cmd
- bootstrap
- config
- migrations
- internal
    - core
      - domain
      - ports
        - repository
        - service
        - handler
      - services
    - database
      - repositories
    - handlers
    - mocks
    - server

## Version de Go

- 1.20.6

## TODO

- Integrar SQLX con PGX
- Cambiar en Database/repositories a Adapter/Database/PostgreSQL/repositories
  -  Agregar un Adapter/Database/MongoDB/repositories
  -  Agregar un Adapter/Database/Neo4J/repositories
- Integrar Redis o Elasticsearch
- Agregar Prometheus pa' las metricas.
- Ir a comer un ramen con los amigos(as)
- Pasar mi certificación de AWS, Azure y Snowflake.
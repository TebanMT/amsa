# Scaffolding
```
amsa/
│
├── app/                      # Capa de Aplicación (Casos de uso y coordinación de flujos)
│   ├── use_cases/            # Casos de uso específicos de la aplicación
│   │   ├── create_client.py  # Caso de uso para 'crear cliente'
│   │   └── ...
│   ├── controllers/          # Controladores para manejar solicitudes HTTP o eventos
│   └── serializers/          # Para la serialización de datos entre casos de uso y entrega
│
├── domain/                   # Capa de Dominio (Lógica de negocio y entidades)
│   ├── entities/             # Entidades del dominio
│   │   ├── client.py         # Entidad 'Client'
│   │   └── ...
│   └── repositories/         # Interfaces para la capa de persistencia
│
├── infrastructure/           # Capa de Infraestructura (Frameworks, DB, adaptadores)
│   ├── db/                   # Integración con la base de datos
│   │   ├── models/           # Modelos ORM
│   │   └── repository_impl/  # Implementaciones concretas de repositorios
│   ├── api/                  # Adaptadores para API (puede ser API Gateway o REST API)
│   └── config/               # Configuración de la aplicación y la infraestructura
│
├── interfaces/               # Interfaces para adaptadores externos
│   ├── rest/                 # Adaptadores para la API REST
│   └── cli/                  # Interfaz de línea de comandos, si es necesaria
│
├── tests/                    # Pruebas automatizadas
│   ├── unit/                 # Pruebas unitarias
│   └── ...
│
├── docker/                   # Archivos Docker para contenedorización
│   ├── Dockerfile            # Dockerfile para construir la imagen del servidor
│   └── docker-compose.yml    # Docker Compose para orquestar contenedores
|
├── terraform/                # Carpeta dedicada a Terraform
│   ├── db/                   # Configuraciones de Terraform para la base de datos
│   │   ├── main.tf           # Archivo principal de Terraform para la base de datos
│   │   ├── variables.tf      # Definiciones de variables para Terraform
│   │   └── outputs.tf        # Salidas de configuración de Terraform
│   └──...
|
├── ci-cd/                    # Configuración para integración y despliegue continuo
│
├── scripts/                  # Scripts de utilidad, como despliegue o inicialización
│
└── README.md                 # Documentación del proyecto
```

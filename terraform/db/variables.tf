variable "aws_region" {
  description = "La región de AWS donde se desplegará la base de datos"
  default     = "us-east-1"
}

variable "db_allocated_storage" {
  description = "El tamaño de almacenamiento para la base de datos (en GB)"
  default     = 20
}

variable "db_instance_class" {
  description = "El tipo de instancia de la base de datos"
  default     = "db.t3.micro"
}

variable "db_name" {
  description = "El nombre de la base de datos"
  default     = "miBaseDeDatos"
}

variable "db_username" {
  description = "El nombre de usuario para la base de datos"
}

variable "db_password" {
  description = "La contraseña para la base de datos"
}
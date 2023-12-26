output "db_instance_endpoint" {
  value       = aws_db_instance.mi_db_postgres.endpoint
  description = "El endpoint de la instancia de la base de datos"
}
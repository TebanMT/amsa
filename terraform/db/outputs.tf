output "db_instance_endpoint" {
  value       = aws_db_instance.amsa_erp.endpoint
  description = "El endpoint de la instancia de la base de datos"
}
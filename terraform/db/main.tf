provider "aws" {
    profile = "default"
    region = var.aws_region
}

resource "aws_db_instance" "amsa_erp" {
  allocated_storage    = var.db_allocated_storage
  db_name              = "amsa"
  engine               = "mysql"
  engine_version       = "8.0.36"
  instance_class       = var.db_instance_class
  username             = var.db_username
  password             = var.db_password
  parameter_group_name = "default.mysql8.0.36"
  skip_final_snapshot  = true
}
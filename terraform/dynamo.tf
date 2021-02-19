resource "aws_dynamodb_table" "web_socket_table" {
  name           = "web_socket_connections"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "connectionId"

  attribute {
    name = "connectionId"
    type = "S"
  }

  tags = {
    Name = "web_socket_connections"
  }
}
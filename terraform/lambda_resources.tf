data "archive_file" "OnConnectZip" {
  type        = "zip"
  source_dir = "dist/connect_lambda"
  output_path = "dist/connect_lambda.zip"
}

data "archive_file" "OnDisconnectZip" {
  type        = "zip"
  source_dir = "dist/disconnect_lambda"
  output_path = "dist/disconnect_lambda.zip"
}

data "archive_file" "MessageZip" {
  type        = "zip"
  source_dir = "dist/message_lambda"
  output_path = "dist/message_lambda.zip"
}
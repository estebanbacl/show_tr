resource "aws_apigatewayv2_api" "web_socket_api" {
  name                       = "web_socket_api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}
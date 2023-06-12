provider "aws" {
  region = "us-east-1"
}

resource "aws_ecs_cluster" "example_cluster" {
  name = "example-cluster"
}

resource "aws_ecs_task_definition" "example_task_definition" {
  family                   = "example-task-family"
  container_definitions    = <<DEFINITION
[
  {
    "name": "example-container",
    "image": "<your-docker-image-url>",
    "cpu": 256,
    "memory": 512,
    "essential": true,
    "portMappings": [
      {
        "containerPort": 8080,
        "hostPort": 8080
      }
    ]
  }
]
DEFINITION
}

resource "aws_ecs_service" "example_service" {
  name            = "example-service"
  cluster         = aws_ecs_cluster.example_cluster.id
  task_definition = aws_ecs_task_definition.example_task_definition.arn
  desired_count   = 1

  network_configuration {
    subnets         = ["<your-subnet-ids>"]
    security_groups = ["<your-security-group-ids>"]
  }

  load_balancer {
    target_group_arn = "<your-target-group-arn>"
    container_name   = "example-container"
    container_port   = 8080
  }
}

#!/bin/sh

eval $(aws ecr get-login --no-include-email --region ap-northeast-1)
docker push $AWS_ACCOUNT_ID.dkr.ecr.ap-northeast-1.amazonaws.com/go-ecs-deploy-app:latest
./scripts/ecs-deploy --cluster go-ecs-deploy-cluster --service-name go-ecs-deploy-service2 --image $AWS_ACCOUNT_ID.dkr.ecr.ap-northeast-1.amazonaws.com/go-ecs-deploy-app:latest --timeout 600
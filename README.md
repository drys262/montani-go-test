# README

## Overview

This repository demonstrates deploying a containerized application on AWS using Amazon ECS (Elastic Container Service) and ECR (Elastic Container Registry), with infrastructure defined through AWS CloudFormation, and CI/CD pipelines managed by GitHub Actions.

## Prerequisites

- AWS Account
- AWS CLI configured with appropriate IAM user credentials
- Docker installed locally
- AWS CloudFormation stack creation permissions
- GitHub repository with appropriate secrets configured

## Files Overview

### `ecs-ecr-setup.yml`
This CloudFormation template sets up the infrastructure required to run the application on ECS:

- **VPC**: Creates a new Virtual Private Cloud with subnets and Internet Gateway.
- **Security Group**: Allows HTTP access on port 8080.
- **Network Load Balancer**: Manages the inbound and outbound traffic.
- **ECS Cluster**: The cluster where the ECS service will run.
- **IAM Role**: Defines permissions for the ECS tasks.
- **Task Definition**: Specifies the Docker container settings.
- **ECR Repository**: Stores the Docker images.
- **ECS Service**: Manages the Fargate tasks in the ECS cluster.

### `.github/workflows/dev.yml`
This GitHub Actions workflow automates the deployment of the application:

- **Login to ECR**: Authenticates GitHub Actions with AWS ECR.
- **Build and Push Docker Image**: Builds and pushes the Docker image to ECR.
- **Render Task Definition**: Updates the ECS task definition with the new Docker image.
- **Deploy ECS Task Definition**: Deploys the updated task definition to the ECS service.

## Step-by-Step Instructions

### 1. Deploy CloudFormation Stack

1. Navigate to the AWS Management Console.
2. Go to the CloudFormation service.
3. Create a new stack by uploading the `ecs-ecr-setup.yml` file.
4. Provide the necessary parameters such as:
   - **ECSClusterName**
   - **ECSServiceName**
   - **ContainerName**
   - **ECRRepositoryName**
   - **Environment**
   - **FeatureBranch**
   - **BranchSha**
5. Wait for the stack to be created successfully. This will set up all the required resources.

### 2. Set Up GitHub Repository

1. Push your application code to a GitHub repository.
2. In your repository, navigate to **Settings** > **Secrets and variables** > **Actions**.
3. Add the following secrets:
   - `AWS_ACCESS_KEY_ID`
   - `AWS_SECRET_ACCESS_KEY`

### 3. Configure GitHub Actions Workflow

The `.github/workflows/dev.yml` is already configured to trigger on pushes to the `main` branch.

## Git Flow

The repository adheres to the Git Flow methodology:

### Branches

- **main**: The main branch is used for development and integrating feature branches. It is not deployed directly to production.
- **release**: The release branch is used for production deployments.
- **feature/***: Feature branches are used for developing new features. These branches are based on the `main` branch.

### 4. Trigger Deployment

1. Push a new commit to the `main` branch.
2. The GitHub Actions workflow will be triggered automatically, performing the following steps:
   - Checkout the code.
   - Configures AWS credentials.
   - Logs in to Amazon ECR.
   - Builds and pushes the Docker image to ECR.
   - Renders the new ECS task definition.
   - Deploys the update to ECS, ensuring service stability.

3. You can monitor the workflow's progress in the **Actions** tab of your GitHub repository.

### 5. Access the Application

After successful deployment, you can access the application via the DNS name of the Network Load Balancer. This is output by the CloudFormation stack and can be found under the **Outputs** section:

- **LoadBalancerDNSName**

Visit the DNS endpoint in your browser to access the application.

## Cleanup

To clean up the resources and avoid incurring unnecessary costs:

1. Delete the CloudFormation stack from the AWS Management Console.
2. Remove the GitHub repository or disable the GitHub Actions workflow.

## Conclusion

This setup demonstrates a fully automated pipeline for deploying containerized applications to AWS ECS using GitHub Actions and AWS CloudFormation.

For any issues or questions, please refer to the AWS documentation or raise an issue in this repository.

NAMESPACE="default"
SERVICE_ACCOUNT="aws-sa"
ROLE_ARN="<your-role-arn>"
CLUSTER_NAME="<your-cluster-name>"

aws eks create-pod-identity-association \
    --namespace "$NAMESPACE" \
    --service-account "$SERVICE_ACCOUNT" \
    --role-arn "$ROLE_ARN" \
    --cluster-name "$CLUSTER_NAME"


DEPLOYMENT_TYPE=""
if [ "$CIRCLE_BRANCH" = "master" ];
then
    echo "Starting Production Deployment"
    DEPLOYMENT_TYPE="Production"
    PROJECT='heph-dev-core'
elif [ "$CIRCLE_BRANCH" = "development" ];
    echo "Starting Development Deployment"
    DEPLOYMENT_TYPE="Development"
    PROJECT='uplifted-elixir-203119'
else
    echo "Starting NOT_CONFIGURED Deployment"
    echo "Completed"
    DEPLOYMENT_TYPE="NOT_CONFIGURED"
    exit 0
fi
# Comment Next Two Lines once configuration is completed
echo "Implementing configurations for CI/D"
exit 0


gcloud app deploy \
--project $PROJECT \

echo "$DEPLOYMENT_TYPE deployment has finished"
exit 0
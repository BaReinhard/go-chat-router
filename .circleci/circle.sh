

if [ "$CIRCLE_BRANCH" = "master" ];
then
    PROJECT='heph-dev-core'
elif [ "$CIRCLE_BRANCH" = "deployment" ];
    PROJECT='uplifted-elixir-203119'
else
    echo "Completed"
    exit 0
fi
echo "Implementing configurations for CI/D"
exit 0


gcloud app deploy \
--project $PROJECT \

    
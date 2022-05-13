PROJECT_ID := $(shell gcloud config get project)

all:
	gcloud beta functions deploy cp-task-2 \
	--gen2 \
	--runtime go116 \
	--trigger-http \
	--region=europe-west4 \
	--entry-point CSSaverServer \
	--source ./cloud_functions/cs-saver-server/ \
	--set-env-vars PROJECT_ID=$(PROJECT_ID)

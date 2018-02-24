# Google App Engine Demo App

This repository contains a demo of what it takes to deploy a simple Hello World Web application into Google App Engine's [Flexible Environment](https://cloud.google.com/appengine/docs/flexible/) using a [Custom Runtime](https://cloud.google.com/appengine/docs/flexible/custom-runtimes/). This is Google's way of saying "BYO Docker Container".

## Pre-requisites
  - Download and install:
    - [Docker Community Edition](https://www.docker.com/community-edition)
    - [Google Cloud SDK](https://cloud.google.com/sdk/)
  - (Optional) Sign up for a new gmail account, Google will provide you with $300 worth of free credits to use over one year so this is definitely worthwhile

## Instructions
### Set up your project
  - Open up your favourite terminal window, and run the following commands in order:
    - `gcloud auth login`
      - A browser window will open up and you will be directed to sign in. Once you've logged in you should be able to close the browser window
    - `gcloud projects create <project-id> --name=<project-name>`
      - A "project" is somewhat analogous to a Microsoft Azure subscription or an AWS account
      - Note that the `<project-id>` needs to be globally unique. For the purpose of this demo it's probably a good idea to have the same `<project-name>` as well
     - `gcloud config set project <project-id>`
       - This sets the context of the terminal session so that all `gcloud` commands are executed against this project
  - Link the project with your billing account by browsing to `https://console.developers.google.com/project/<project-id>/settings`. This is a mandatory step before you can make any deployments

### Build and deploy your application container (optional)
  - This step is optional, as I've already provided a public container image `williamtsoi/hello-world-go` in Docker Hub
  - The source code and `Dockerfile` for this Docker image is in the `./app` directory of this repo
  - If you would like to build your own container, note the requirements from [this guide](https://cloud.google.com/appengine/docs/flexible/custom-runtimes/build). This will be invaluable in letting you know what is required of your container in order for it to be hosted by Google App Engine

### Deploy your application into Google App Engine
  - Go into the `./infrastructure` directory of this repo
  - Open `app.yaml`
    - This contains the configuration required for the app to run
    - The full documentation of `app.yaml` can be found [here](https://cloud.google.com/appengine/docs/flexible/custom-runtimes/configuring-your-app-with-app-yaml)
  - Open `Dockerfile`
    - This should be a single line Dockerfile which references the container image in the step above. Replace the image with one of your own if you wish to deploy your own container image
  - Run `gcloud app deploy`
    - If this is your first deployment you will be prompted for the region for your app. Choose whatever is appropriate
    - Please also ensure that you have linked this project to your billing account, else this step will fail
  - Run `gcloud app browse`
    - If all goes well, a new browser window will open and you will be greeted with your application!
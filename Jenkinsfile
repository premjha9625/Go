pipeline {
    agent any

    environment {
        // Define Docker Hub credentials
        DOCKER_HUB_CREDENTIALS = 'dockerhub'
        // Define the Docker image name and tag
        DOCKER_IMAGE_NAME = '885185/go-api'
        DOCKER_IMAGE_TAG = 'latest'
        GITHUB_CREDENTIALS = 'github'
    }

    stages {
        stage('Cloning our Git') { 
            scripts{
        steps { 
        //         git credentialsId: 'f87a34a8-0e09-45e7-b9cf-6dc68feac670', 
        //         url: 'https://github.com/premjha9625/Go.git',
        //         branch: 'main'
        //         git 'https://github.com/premjha9625/Go.git' 
        //     }
            withCredentials([usernamePassword(credentialsId: GITHUB_CREDENTIALS, passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
                        sh ``` git 'https://github.com/premjha9625/Go.git'  ```                       
                    }
            }
         }
        }
        stage('Checkout'){
           steps {
                git credentialsId: 'github', 
                url: 'https://github.com/premjha9625/Go.git',
                branch: 'dev'
           }
        }
        stage('Build') {
            steps {
                scripts{
                // Checkout the source code from your Git repository
                //git 'https://github.com/premjha9625/Go'

                // Build the Go application
                //sh 'go build -o myapp'

                // Build the Docker image
                sh 'docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} .'
                }
            }
        }

        stage('Push') {
            steps {
                scripts{
                // Authenticate with Docker Hub
                withCredentials([usernamePassword(credentialsId: DOCKER_HUB_CREDENTIALS, usernameVariable: 'DOCKER_HUB_USERNAME', passwordVariable: 'DOCKER_HUB_PASSWORD')]) {
                    sh "docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_PASSWORD}"
                }

                // Push the Docker image to Docker Hub
                sh "docker push ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
            }
        }
        }
    }
}

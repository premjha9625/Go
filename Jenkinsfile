pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = 'dockerhub'
        DOCKER_IMAGE_NAME = '885185/go-api'
        DOCKER_IMAGE_TAG = 'latest'
        GITHUB_CREDENTIALS = 'github'
    }

    stages {
        stage('Clone repository') {
            steps {
                git credentialsId: 'github', 
                url: 'https://github.com/premjha9625/Go',
                branch: 'dev'
            }
        }

        stage('Build image') {
            steps {
                script {
                    sh "docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} ."
                }
            }
        }

        stage('Push image') {
            steps {
                script {
                    withDockerRegistry([credentialsId: "dockerhub", url: ""]) {
                        sh "docker push ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                    }
                }
            }
        }

        stage('Cleaning up') {
            steps {
                sh "docker rmi ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
            }
        } 
    }
}

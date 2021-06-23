pipeline {
    agent any

    environment {
        BINARY_NAME="a3v1"
    }

    stages {
        stage('Build') {
            steps {
                sh 'go version'

                sh "go build ${BINARY_NAME}"
            }
        }

        stage('Execute') {
            steps {
                sh "./${BINARY_NAME}"
            }
        }
    }

    options {
        buildDiscarder(
            logRotator(numToKeepStr: '7')
        )
    }
}
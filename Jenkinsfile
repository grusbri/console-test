pipeline {
    agent any

    environment {
        BINARY_NAME="a3v1"
    }

    options {
        buildDiscarder(
            logRotator(numToKeepStr: '7')
        )
    }

    tools {
        go 'Go 1.16.5'
    }

    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh "go build -o ${BINARY_NAME}"
            }
        }

        stage('Execute') {
            steps {
                sh "./${BINARY_NAME}"
            }
        }
    }
}
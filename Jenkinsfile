pipeline {
    agent any

    environment {
        BINARY_NAME="a3v1"
    }

    options {
        ansiColor('xterm')
        buildDiscarder(
            logRotator(numToKeepStr: '7')
        )
    }

    tools {
        go 'Go 1.16.5'
    }

    triggers {
        cron('0 0/15 * 1/1 * ? *')
    }

    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh "go build -o ${BINARY_NAME}"
            }
        }

        stage('Validate') {
            steps {
                sh "./${BINARY_NAME}"
            }
        }

        stage('Execute') {
            steps {
                sh "./${BINARY_NAME}"
            }
        }
    }
}
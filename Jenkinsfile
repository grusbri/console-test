pipeline {
    agent any

    stages {
        stage('Build') {
            agent {
                docker {
                    image `golang:1.16.5`
                    reuseNode true
                }
            }

            steps {
                sh 'go version'
            }
        }
    }
}
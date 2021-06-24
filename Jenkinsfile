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
        cron('*/15 6-18 * * 1-5')
    }

    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh "go build -o ${BINARY_NAME}"
            }
        }

        stage('Confirm') {
            when {
                not { branch 'prod' }
            }
            steps {
                sh "./${BINARY_NAME}"
                script {
                    env.PROCEED = input message: "LGTM. Proceed?"
                        parameters: [choice(name: "Proceed?", choices:['YES', 'NO'])]
                }
            }
        }

        stage('Execute') {
            when {
                anyOf {
                    environment name: 'PROCEED', value: 'YES';
                    triggeredBy 'TimerTrigger';
                }
            }
            steps {
                sh "./${BINARY_NAME}"
            }
        }
    }
}
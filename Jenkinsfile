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
        cron('*/15 * * * 1-5')
    }

    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh "go build -o ${BINARY_NAME}"
            }
        }

        stage('Dry Run') {
            steps {
                sh "./${BINARY_NAME}" // -dryrun
            }
        }

        stage('Confirm') {
            when {
                not { branch 'master' }
            }
            steps {
                script {
                    env.PROCEED = input message: "Proceed or Abort?",
                    parameters: [choice(name: "proceed", choices:['YES'])]
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
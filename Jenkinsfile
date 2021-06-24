pipeline {
    agent any

    environment {
        BINARY_NAME="a3v1"
    }

    options {
        ansiColor('xterm')
        disableConcurrentBuilds()
        buildDiscarder(logRotator(numToKeepStr: '7'))
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
            steps {
                script {
                    if(env.GIT_BRANCH != 'master') {
                        env.PROCEED = input message: "Proceed?", ok: 'OK',
                        parameters: [choice(name: "proceed", choices:['YES', 'NO'])]
                    }
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
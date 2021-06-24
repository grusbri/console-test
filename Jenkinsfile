pipeline {
    agent any

    environment {
        BINARY_NAME="a3"
    }

    options {
        ansiColor('xterm')
        disableConcurrentBuilds()
        buildDiscarder(logRotator(numToKeepStr: '7'))
    }

    triggers {
        cron('*/15 * * * 1-5')
    }

    stages {

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.16.5'
                    reuseNode true
                }
            }
            steps {
                sh 'go version'
                sh "go build -o ${BINARY_NAME}"
            }
        }

        stage('Plan Changes') {
            steps {
                sh "./${BINARY_NAME} plan"
            }
        }

        stage('Confirm') {
            options {
                timeout(time: 2, unit: 'MINUTES')
            }

            steps {
                script {
                    def buildCause = currentBuild.rawBuild.getCauses()
                    echo "Current build was caused by: ${buildCause}\n"

                    // Force users to commit to a non-master branch first
                    if(env.GIT_BRANCH != 'origin/master' || "${buildCause}".contains("UserIdCause")) {
                        env.PROCEED = input message: "Proceed?", ok: 'OK',
                        parameters: [choice(name: "proceed", choices:['YES', 'NO'])]
                    }
                }
            }
        }

        stage('Apply Changes') {
            when { 
                anyOf {
                    environment name: 'PROCEED', value: 'YES';
                    triggeredBy 'TimerTrigger';
                }
            }
            steps {
                sh "./${BINARY_NAME} apply"
            }
        }
    }
}
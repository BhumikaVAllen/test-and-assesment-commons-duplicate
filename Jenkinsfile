// @Library('Allen_Shared_Libraries') _
pipeline {
    agent {
        docker {
            alwaysPull true
            image '537984406465.dkr.ecr.ap-south-1.amazonaws.com/allen-jenkins-agent:latest'
            registryUrl 'https://537984406465.dkr.ecr.ap-south-1.amazonaws.com'
            registryCredentialsId 'ecr:ap-south-1:AWSKey'
            args '-v /var/run/docker.sock:/var/run/docker.sock -u 0:0'
        }
    }

    stages {
        stage('Checkout Code') {
            steps {
                // Get the latest code from your version control system
                checkout scm
            }
        }

        stage('Sensitive Information Check - Git Secret') {
            steps {
                // Perform sensitive information check
                echo 'Scanning...'
                sensitiveInformationCheck()
            }
        }

//         stage('Lint Checks') {
//             steps {
//             sh '''
//             echo "Running golang lint"
//             make lint
//             '''
//             }
//         }

        stage('Run Tests') {
            steps {
                withCredentials([string(credentialsId: 'git_pat', variable: 'GITHUB_PAT')]) {
                    sh '''
                        # Configure Git with the provided GitHub Personal Access Token (GITHUB_PAT)
                        echo 'Configuring Git...'
                        git config --global --add safe.directory '*'
                        git config --global url.https://$GITHUB_PAT@github.com/.insteadOf https://github.com/

                        # Initialize the project (if necessary)
                        echo 'Initializing the project...'
                        make init

                        # Configure project settings
                        echo 'Configuring project settings...'
                        make config

                        go mod tidy
                    '''
                }

                echo "Verifying mocks..."
                 sh '''
                    ./verify_mocks.sh
                '''

                // Run tests on your code
                echo 'Testing...'
                sh '''
                     echo 'Running unit tests...'
                    # Run unit tests and generate code coverage report
                    go test ./... -covermode=count -coverprofile=coverage.out
                    echo 'Unit tests completed.'

                    echo 'Running tests for SonarQube...'
                    # Run additional tests for SonarQube and generate code coverage report
                    cp coverage.out sonar_coverage.out
                    echo 'Tests for SonarQube completed.'
                '''

            }
        }

        stage('Integration Testing') {
            steps {
                // Run integration tests
                echo 'Running Integration tests...'
            }
        }

        stage('SonarQube Scan') {
            steps {
                // Perform SonarQube scan
                echo 'Running Sonarqube Scan...'
                sonarQubeScan('test-and-assessment-commons') // Pass the projectKey value as an argument
            }
        }

        stage("Quality Gate Check") {
            steps {
                // Check quality gate status
                echo 'Running Sonarqube Quality Gate check...'
                // qualityGateCheck()
            }
        }

        stage("Create new release") {
            when {
                branch 'main'
            }
            steps {
                withCredentials([string(credentialsId: 'git_pat', variable: 'GITHUB_PAT')]) {
                    echo "Running generate release script..."
                    sh '''
                        ./auto_release.sh $GITHUB_PAT
                    '''
                }
            }
        }
    }

    post {
        always {
            echo 'Cleaning up resources...'
            // Add any clean-up actions here
            cleanupResources()
        }
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
            // Handle failures, for example, notify stakeholders
        }
    }
}

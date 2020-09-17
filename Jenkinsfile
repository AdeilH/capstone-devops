pipeline{
    agent any
    stages{
        stage('lint'){
            steps{
            sh 'hadolint Dockerfile'
            sh 'golint main.go'
               }
           }
        stage('build docker container'){
            steps{
            sh 'docker build -t gowebapp .'   
            }
        }
        stage('Push Docker Image') {
              steps {
                  withDockerRegistry([url: "", credentialsId: "adeel"]) {
                      sh "docker tag gowebapp adeelhussain13/gowebapp"
                      sh 'docker push adeelhussain13/gowebapp'
                  }
              }
         }
         stage('Deploying to AWS') {
              steps{
                  withAWS(credentials: 'adeel-aws', region: 'us-west-2') {
                      sh "aws eks --region us-west-2 update-kubeconfig --name my-capstone-cluster"
                      sh "kubectl config use-context arn:aws:eks:us-west-2:939300009960:cluster/my-capstone-cluster"
                      sh "kubectl apply -f deployment.yml"
                      sh "kubectl get nodes"
                      sh "kubectl get deployment"
                      sh "kubectl get pod -o wide"
                      sh "kubectl get gowebapp"
                  }
              }
        }
        stage("Cleaning up") {
              steps{
                    sh "docker system prune"
              }
        }
    }
}
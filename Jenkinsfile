pipeline{
    agent any
    stages{
        stage('lint'){
            steps{
            sh 'tidy -q -e *.html'
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
                      sh "docker tag gowebapp gowebapp/adeelhussain13"
                      sh 'docker push adeelhussain13/gowebapp'
                  }
              }
         }
         stage('Deploying to AWS') {
              steps{
                  withAWS(credentials: 'aws', region: 'us-west-2') {
                      sh "aws eks --region us-west-2 update-kubeconfig --name capstonecluster"
                      sh "kubectl config use-context arn:aws:eks:us-west-2:988212813982:cluster/capstonecluster"
                      sh "kubectl set image deployments/gowebapp gowebapp=adeelhussain13/gowebapp:latest"
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
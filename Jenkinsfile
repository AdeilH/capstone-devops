pipeline{
    agent any
    stages{
        stage('lint'){
            steps{
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
                      sh 'kubectl set image deployments/gowebapp capstone-project=adeelhussain13/gowebapp'
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
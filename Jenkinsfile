  node("TestMachine-ut") {
    stage("Connecting To Azure Artifacts...."){
      withCredentials([usernamePassword(credentialsId: 'Azure-Cred', passwordVariable: 'AZPASS', usernameVariable: 'AZUSER')]) {
          sh 'docker run mcr.microsoft.com/azure-cli /bin/bash -c "export AZURE_DEVOPS_EXT_PAT=$AZPASS ; az extension add --name azure-devopsaz extension add --name azure-devops ; az devops login --organization $AZUSER"'
      }
    }
    stage("Preapring Environment"){
      sh "rm -rf /datavolume1/* ; mkdir /tmp/${BUILD_NUMBER} ; git clone https://github.com/mms-cv/cd-demo.git /tmp/${BUILD_NUMBER}/ ; mv /tmp/${BUILD_NUMBER}/* /datavolume1/"
    }
    stage("Unit Test") {
      sh "cd /datavolume1 ; docker run --rm -v DataVolume1:/go/src/cd-demo golang go test cd-demo -v --run Unit"
    }
    stage("Integration Test") { 
      try {
        sh "cd /datavolume1 ; docker build -t cd-demo ."
        sh "docker rm -f cd-demo || true"
        sh "docker run -d -p 8080:8080 --name=cd-demo cd-demo"
        // env variable is used to set the server where go test will connect to run the test 
        sh "docker run --rm -v DataVolume1:/go/src/cd-demo --link=cd-demo -e SERVER=cd-demo golang go test cd-demo -v --run Integration"
      }
      catch(e) {
        error "Integration Test failed"
      }finally {
        sh "docker rm -f cd-demo || true"
        sh "docker ps -aq | xargs docker rm || true"
        sh "docker images -aq -f dangling=true | xargs docker rmi || true"
      }
    }
    stage('Logging Into Harbor'){
          withCredentials([usernamePassword(credentialsId: 'harbor-sec', passwordVariable: 'HPASS', usernameVariable: 'HUSER')]) {
              sh 'echo "10.0.0.145    harbor.this" >> /etc/hosts'
              sh 'echo $HPASS > ~/pass.txt'
              sh 'docker login -u $HUSER harbor.this --password-stdin < ~/pass.txt'
          }
    }
    stage("Build") {
      sh "cd /datavolume1 ; docker build -t harbor.this/codevalue/cd-demo:${BUILD_NUMBER} ."
    }
    stage("Publish") {
        sh "docker push harbor.this/codevalue/cd-demo:${BUILD_NUMBER}"
    }
  }

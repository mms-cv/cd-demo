  node("TestMachine-ut") {
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
   /* stage('Logging Into Harbor'){
          withCredentials([usernamePassword(credentialsId: 'harbor-sec', passwordVariable: 'HPASS', usernameVariable: 'HUSER')]) {
              sh 'echo "10.0.0.145    harbor.this" >> /etc/hosts'
              sh 'echo $HPASS > ~/pass.txt'
              sh 'docker login -u $HUSER harbor.this --password-stdin < ~/pass.txt'
          }
    }*/
    stage('Uploading Artifact To Jfrog'){
      withCredentials([usernamePassword(credentialsId: 'JfrogArtifacte', passwordVariable: 'JPASSWORD', usernameVariable: 'JUSER')]) {
          sh "tar -czvf go-test-project_${BUILD_NUMBER}.tar.gz /datavolume1/ ; curl -u$JUSER:$JPASSWORD -T go-test-project_${BUILD_NUMBER}.tar.gz 'https://golan.jfrog.io/golan/go/go-test-project_${BUILD_NUMBER}.tar.gz'"
      }
    }
    stage("Build") {
      sh "cd /datavolume1 ; docker build -t mms2020/cd-demo:${BUILD_NUMBER} ."
    }
    stage("Publish") {
      withCredentials([usernamePassword(credentialsId: 'DockerHub', passwordVariable: 'DPASSWORD', usernameVariable: 'DUSER')]) {
        sh "echo $DPASSWORD | docker login --username $DUSER --password-stdin "
        sh "docker push mms2020/cd-demo:${BUILD_NUMBER}"
      }
    }
  }

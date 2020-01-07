  node("TestMachine-ut") {
    stage("Unit Test") {
      try{
        sh 'cd /datavolume1 ; git clone https://github.com/mms-cv/cd-demo.git . ;  ls ; pwd'
      }catch(e){
        sh 'cd /datavolume1 ; git pull https://github.com/mms-cv/cd-demo.git ;  ls ; pwd'
      }
      sh "docker run --rm -v DataVolume1:/go/src/cd-demo golang go test cd-demo -v --run Unit"
      sh "ls ; pwd"
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
    stage('logging into harbor'){
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

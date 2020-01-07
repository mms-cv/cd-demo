  env.DOCKERHUB_USERNAME = 'mms-cv'

  node("TestMachine-ut") {
    checkout scm
    stage("Unit Test") {
<<<<<<< HEAD
      sh "docker run --rm -v cd-demo:/go/src/cd-demo golang go test cd-demo -v --run Unit"
=======
      sh 'cd /datavolume1 ; git clone https://github.com/mms-cv/cd-demo.git . ;  ls ; pwd'
      sh "docker run --rm -v DataVolume1:/go/src/cd-demo golang go test cd-demo -v --run Unit"
>>>>>>> c06e9e20e73d9617852076829dfd46020687e143
    }
    stage("Integration Test") { 
      try {
        sh "docker build -t cd-demo ."
        sh "docker rm -f cd-demo || true"
        sh "docker run -d -p 8080:8080 --name=cd-demo cd-demo"
        // env variable is used to set the server where go test will connect to run the test
<<<<<<< HEAD
        sh "docker run --rm -v cd-demo:/go/src/cd-demo --link=cd-demo -e SERVER=cd-demo golang go test cd-demo -v --run Integration"
      } 
=======
        sh "docker run --rm -v DataVolume1:/go/src/cd-demo --link=cd-demo -e SERVER=cd-demo golang go test cd-demo -v --run Integration"
      }
>>>>>>> c06e9e20e73d9617852076829dfd46020687e143
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
      sh "docker build -t harbor.this/codevalue/cd-demo:${BUILD_NUMBER} ."
    }
    stage("Publish") {
        sh "docker push harbor.this/codevalue/cd-demo:${BUILD_NUMBER}"
    }
  }

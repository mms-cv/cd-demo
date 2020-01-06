  env.DOCKERHUB_USERNAME = 'mms-cv'

  node("TestMachine-ut") {
    checkout scm
    stage("Unit Test") {
      //sh 'cd /datavolume1 ; git clone https://github.com/mms-cv/cd-demo.git ;  ls ; pwd'
      sh "docker run --rm -v DataVolume1/cd-demo:/go/src/cd-demo golang go test cd-demo -v --run Unit"
    }
    stage("Integration Test") { 
      try {
        sh "docker build -t cd-demo ."
        sh "docker rm -f cd-demo || true"
        sh "docker run -d -p 8080:8080 --name=cd-demo cd-demo"
        // env variable is used to set the server where go test will connect to run the test
        sh "docker run --rm -v DataVolume1/cd-demo:/go/src/cd-demo --link=cd-demo -e SERVER=cd-demo golang go test cd-demo -v --run Integration"
      }
      catch(e) {
        error "Integration Test failed"
      }finally {
        sh "docker rm -f cd-demo || true"
        sh "docker ps -aq | xargs docker rm || true"
        sh "docker images -aq -f dangling=true | xargs docker rmi || true"
      }
    }
    stage("Build") {
      sh "docker build -t ${DOCKERHUB_USERNAME}/cd-demo:${BUILD_NUMBER} ."
    }
    stage("Publish") {
     // withDockerRegistry([credentialsId: 'DockerHub']) {
     //   sh "docker push ${DOCKERHUB_USERNAME}/cd-demo:${BUILD_NUMBER}"
     // }
    }
  }

 /* node("docker-stage") {
    checkout scm

    stage("Staging") {
      try {
        sh "docker rm -f cd-demo || true"
        sh "docker run -d -p 8080:8080 --name=cd-demo ${DOCKERHUB_USERNAME}/cd-demo:${BUILD_NUMBER}"
        sh "docker run --rm -v ${WORKSPACE}:/go/src/cd-demo --link=cd-demo -e SERVER=cd-demo golang go test cd-demo -v"

      } catch(e) {
        error "Staging failed"
      } finally {
        sh "docker rm -f cd-demo || true"
        sh "docker ps -aq | xargs docker rm || true"
        sh "docker images -aq -f dangling=true | xargs docker rmi || true"
      }
    }
  }*/
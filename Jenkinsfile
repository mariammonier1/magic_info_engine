def jobnameparts = JOB_NAME.tokenize('/') as String[]
def jobconsolename = jobnameparts.length < 2 ? jobnameparts[0] : jobnameparts[jobnameparts.length - 2]

pipeline {
  agent any
  stages {
    stage('Build-kuido') {
      when {
        branch 'kuido'
      }
      steps {
         script {
           docker.withRegistry("https://$registry", 'nexus') {
               def image = docker.build("$repo/$jobconsolename")
               image.push("latest-$BRANCH_NAME")
               image.push(env.BUILD_TIMESTAMP)
           }
         }
      }
    }
  }
  environment {
    registry = 'hub.iotblue.io'
    repo = 'cervello'
    registryCredential = 'nexus'
    release = 'patch'
  }
}

node {
	stage('Checkout') {
		checkout scm
	}
	stage('Build') {
		sh 'bazel build //...'
	}
	stage('Test') {
		sh 'bazel test //...'
	}
}

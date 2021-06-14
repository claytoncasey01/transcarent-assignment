# Transcarent Assigment

###  Usage

 - Run `git clone
   https://github.com/claytoncasey01/transcarent-assignment.git`
  - CD into the directory where the project was cloned.
- CD into server
	 - #### With Docker
		  - run `make container` to build the docker container
		  - run `make run-container`
		  - Navigate to `http://localhost:8080/v1/user-posts/{id}` where id is a number between 1 and 10
	- #### Without Docker
		- run `make build`
		-  Execute the generated file
		- OR run `make run` to run without building
		- Navigate to `http://localhost:8080/v1/user-posts/{id}` where id is a number between 1 and 10

	- #### Running Tests
		- Simply run `make test` **(go must be installed)**

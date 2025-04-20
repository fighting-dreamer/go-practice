## Basic App : 

1. setup config.
2. setup logger 
3. setup dependencies
	1. setup external clients
		1. database-clients
		2. http client/sdks
		3. message queue system
	2. setup internal logical services/constructs : that are related to runtime and OS
		1. some transactional lock manager
		2. some in-memory store
		3. some pool manager
		4. some executor
		5. some file reader/writer etc...
		6. logger, or some other middleware can be setup.
	3. setup services
		1. setup repo and add the relevant db client, logger etc...
		2. setup app services
		3. setup http routes/server
		4. any service that was to depend on earlier instances, is to be setup here.

4. you need to start the services and server



#!/bin/bash

migrate -verbose -path "/migrations" -database $DATABASE_URL up
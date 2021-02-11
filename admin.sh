#!/bin/bash
create_email(){
   curl -X POST -H "Content-Type: application/json" \
       -d @.email.js \
    http://localhost:1323/admin/email
}

create_job(){
   curl -X POST -H "Content-Type: application/json" \
       -d @.job.json \
    http://localhost:1323/admin/job

}

case $1 in
    email) create_email ;;
    job) create_job;;
    *)
esac

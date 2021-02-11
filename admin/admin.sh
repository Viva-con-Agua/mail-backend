#!/bin/bash
create_email(){
   curl -X POST -H "Content-Type: application/json" \
       -d @.email.json \
    http://localhost:1337/admin/email/email
}

create_job(){
   curl -X POST -H "Content-Type: application/json" \
       -d @${1} \
    http://localhost:1323/admin/email/job

}

case $1 in
    email) create_email ;;
    job) create_job $2;;
    *)
esac

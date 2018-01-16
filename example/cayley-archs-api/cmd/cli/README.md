# Admin CLI

The CLI is used for create and manage administrators, clinics, and clinics employees.

The following commands are available:
```
./cli list-quads [-config /tmp/config.json]
./cli list-admins
./cli list-clinics
./cli list-employees
./cli get-clinic -id 111
./cli add-admin -email vera@gmail.com -password 112233 -name vera
./cli login-admin -email vera@gmail.com -password 112233
./cli add-clinic -clinic-file test-data/clinic-all.json -jwt 111
./cli add-employee -fname jonh -lname doe -email john@gmail.com -clinic-id 111 -jwt 111
./cli delete-clinic -id 111 -jwt 111
```

## Endpoints

### Admin Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"vera@gmail.com","password":"112233"}' "http://localhost:3000/adminlogin"

### Admin create clinic

    curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"groupName":"Group Name","clinicFullName":"Clinic Full Name","clinicShortName":"Clinic Short Name (displayed in search results)","branchName":"Branch Name","address1":"Address 1","address2":"Address 2","addressCity":"Address City","addressZipCode":"Address Zip Code","addressCountry":"Address Country","officeTel":"Office Tel","officeFax":"Office Fax","website":"Website","clinicEmail":"Clinic Email","clinicAdminFullName":"Clinic Admin Full Name","clinicAdminMobile":"Clinic Admin Mobile","clinicAdminEmail":"Clinic Admin Email","clinicAdminPassword":"Clinic Admin Password","clinicAdminConfirmPassword":"Clinic Admin Confirm Password","autoBidToggle":"Auto bid toggle","priority":"Priority","maxNbOfPatientsPerHour":"Max Nb of Patients per Hour","notificationsMobile":"Notifications Mobile","notificationsEmail":"Notifications Email","togglePushNotifications":"Toggle Push Notifications","pocFullName":"POC Full Name","pocMobile":"POC Mobile","pocEmail":"POC Email"}' "http://localhost:3000/clinics"

### Admin create employee

    curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"fname":"John", "lname":"Doe", "email":"johndoe@gmail.com", "clinic_id":"111"}' "http://localhost:3000/employees"

### Admin get all clinics

    curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" "http://localhost:3000/clinics"

### Admin get a specific clinics

    curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" "http://localhost:3000/clinics/1"

### Patient Signup

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/signup"

### Patient Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"

### Get User Settings

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"

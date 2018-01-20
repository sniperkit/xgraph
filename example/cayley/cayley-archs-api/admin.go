package admin

// todo: use context to keep my db connection
// or use a service structure

import (
	"errors"
	"log"

	"github.com/cayleygraph/cayley"
	_ "github.com/cayleygraph/cayley/graph/kv/bolt"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// var dbPath = "C:/Users/Alan/Projects/data/db.boltdb"
var MySigningKey = []byte("secret")

func init() {
	schema.RegisterType("Admin", Admin{})
	schema.RegisterType("Clinic", Clinic{})
	schema.RegisterType("Employee", Employee{})
	schema.GenerateID = func(_ interface{}) quad.Value {
		return quad.IRI(uuid.NewV1().String())
	}
}

const (
	Monday = quad.IRI("http://schema.org/Monday")
)

type Admin struct {
	Name           string `json:"name" quad:"name"`
	Email          string `json:"email" quad:"email"`
	HashedPassword string `json:"hashedPassword"  quad:"hashed_password"`
	LoggedIn       bool
	Store          *cayley.Handle
}

type Clinic struct {
	ID                         quad.IRI       `json:"id" quad:"@id"`
	CreatedBy                  quad.IRI       `quad:"createdBy"`
	Employees                  []Employee     `json:"employees"`
	GroupName                  string         `json:"groupName" quad:"groupName"`
	ClinicFullName             string         `json:"clinicFullName" quad:"clinicFullName"`
	ClinicShortName            string         `json:"clinicShortName" quad:"clinicShortName"`
	BranchName                 string         `json:"branchName" quad:"branchName"`
	Address1                   string         `json:"address1" quad:"address1"`
	Address2                   string         `json:"address2" quad:"address2"`
	AddressCity                string         `json:"addressCity" quad:"addressCity"`
	AddressZipCode             string         `json:"addressZipCode" quad:"addressZipCode"`
	AddressCountry             string         `json:"addressCountry" quad:"addressCountry"`
	OfficeTel                  string         `json:"officeTel" quad:"officeTel"`
	OfficeFax                  string         `json:"officeFax" quad:"officeFax"`
	Website                    string         `json:"website" quad:"website"`
	ClinicEmail                string         `json:"clinicEmail" quad:"clinicEmail"`
	ClinicAdminFullName        string         `json:"clinicAdminFullName" quad:"clinicAdminFullName"`
	ClinicAdminMobile          string         `json:"clinicAdminMobile" quad:"clinicAdminMobile"`
	ClinicAdminEmail           string         `json:"clinicAdminEmail" quad:"clinicAdminEmail"`
	ClinicAdminPassword        string         `json:"clinicAdminPassword" quad:"clinicAdminPassword"`
	ClinicAdminConfirmPassword string         `json:"clinicAdminConfirmPassword" quad:"clinicAdminConfirmPassword"`
	AutoBidToggle              bool           `json:"autoBidToggle" quad:"autoBidToggle"`
	Priority                   bool           `json:"priority" quad:"priority"`
	MaxNbOfPatientsPerHour     int            `json:"maxNbOfPatientsPerHour" quad:"maxNbOfPatientsPerHour"`
	NotificationsMobile        string         `json:"notificationsMobile" quad:"notificationsMobile"`
	NotificationsEmail         string         `json:"notificationsEmail" quad:"notificationsEmail"`
	TogglePushNotifications    bool           `json:"togglePushNotifications" quad:"togglePushNotifications"`
	PocFullName                string         `json:"pocFullName" quad:"pocFullName"`
	PocMobile                  string         `json:"pocMobile" quad:"pocMobile"`
	PocEmail                   string         `json:"pocEmail" quad:"pocEmail"`
	Hours                      []OpeningHours `json:"hours" quads:"schema:openingHoursSpecification"`
}

func (c Clinic) Validate() error {
	// TODO: add more validations
	if c.ClinicFullName == "" {
		return errors.New("Name cannot be empty")
	}

	return nil

}

// if c.GroupName == "" || c.ClinicFullName == "" || c.ClinicShortName == "" || c.BranchName == "" || c.Address1 == "" || c.Address2 == "" || c.AddressCity == "" || c.AddressZipCode == "" || c.AddressCountry == "" || c.OfficeTel == "" || c.OfficeFax == "" || c.Website == "" || c.ClinicEmail == "" || c.ClinicAdminFullName == "" || c.ClinicAdminMobile == "" || c.ClinicAdminEmail == "" || c.ClinicAdminPassword == "" || c.ClinicAdminConfirmPassword == "" || c.AutoBidToggle == "" || c.Priority == "" || c.MaxNbOfPatientsPerHour == "" || c.NotificationsMobile == "" || c.NotificationsEmail == "" || c.TogglePushNotifications == "" || c.PocFullName == "" || c.PocMobile == "" || c.PocEmail == "" {

type OpeningHours struct {
	DayOfWeek quad.IRI `json:"day" quad:"schema:dayOfWeek"` // set to one of consts like the one above
	Slot      int      `json:"slot" quad:"slot"`
	Opens     string   `json:"opens" quad:"schema:opens"` // ex: 12:00 or 12:00:00
	Closes    string   `json:"closes" quad:"schema:closes"`
}

type NewClinic struct {
	GroupName                  string         `json:"groupName" quad:"groupName"`
	ClinicFullName             string         `json:"clinicFullName" quad:"clinicFullName"`
	ClinicShortName            string         `json:"clinicShortName" quad:"clinicShortName"`
	BranchName                 string         `json:"branchName" quad:"branchName"`
	Address1                   string         `json:"address1" quad:"address1"`
	Address2                   string         `json:"address2" quad:"address2"`
	AddressCity                string         `json:"addressCity" quad:"addressCity"`
	AddressZipCode             string         `json:"addressZipCode" quad:"addressZipCode"`
	AddressCountry             string         `json:"addressCountry" quad:"addressCountry"`
	OfficeTel                  string         `json:"officeTel" quad:"officeTel"`
	OfficeFax                  string         `json:"officeFax" quad:"officeFax"`
	Website                    string         `json:"website" quad:"website"`
	ClinicEmail                string         `json:"clinicEmail" quad:"clinicEmail"`
	ClinicAdminFullName        string         `json:"clinicAdminFullName" quad:"clinicAdminFullName"`
	ClinicAdminMobile          string         `json:"clinicAdminMobile" quad:"clinicAdminMobile"`
	ClinicAdminEmail           string         `json:"clinicAdminEmail" quad:"clinicAdminEmail"`
	ClinicAdminPassword        string         `json:"clinicAdminPassword" quad:"clinicAdminPassword"`
	ClinicAdminConfirmPassword string         `json:"clinicAdminConfirmPassword" quad:"clinicAdminConfirmPassword"`
	AutoBidToggle              bool           `json:"autoBidToggle" quad:"autoBidToggle"`
	Priority                   bool           `json:"priority" quad:"priority"`
	MaxNbOfPatientsPerHour     int            `json:"maxNbOfPatientsPerHour" quad:"maxNbOfPatientsPerHour"`
	NotificationsMobile        string         `json:"notificationsMobile" quad:"notificationsMobile"`
	NotificationsEmail         string         `json:"notificationsEmail" quad:"notificationsEmail"`
	TogglePushNotifications    bool           `json:"togglePushNotifications" quad:"togglePushNotifications"`
	PocFullName                string         `json:"pocFullName" quad:"pocFullName"`
	PocMobile                  string         `json:"pocMobile" quad:"pocMobile"`
	PocEmail                   string         `json:"pocEmail" quad:"pocEmail"`
	Hours                      []OpeningHours `json:"hours" quads:"schema:openingHoursSpecification"`
}

type Employee struct {
	ID        quad.IRI `json:"id" quad:"@id"`
	FName     string   `json:"fname" quad:"fname"`
	LName     string   `json:"lname" quad:"lname"`
	Email     string   `json:"email" quad:"email"`
	WorkFor   quad.IRI `quad:"workFor"`
	CreatedBy quad.IRI `quad:"createdBy"`
}

type NewEmployee struct {
	FName    string `json:"fname"`
	LName    string `json:"lname"`
	Email    string `json:"email"`
	ClinicId string `json:"clinic_id"`
}

type EmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// used for returning the login endpoint
type User struct {
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}

type AdminService interface {
	CreateAdmin(a *Admin, password string) error
	Login(password string, email string) (string, error)
	Authenticate(jwt string) (*MyCustomClaims, error)
	AddClinic(c *Clinic, email string) error
	All() ([]Admin, error)
	AllClinics() ([]Clinic, error)
	GetClinic(clinicId string) (Clinic, error)
	AddEmployee(e *NewEmployee, email string) error
	AllEmployees() ([]Employee, error)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func New(store *cayley.Handle) (*Admin, error) {
	a := &Admin{}
	a.Store = store
	a.LoggedIn = false

	return a, nil
}

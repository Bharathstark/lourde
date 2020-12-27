package database

import (
	"time"

	"github.com/Bharathstark/lourde/back-end/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

//Initialize used to Initialize the DB Pool
func Initialize() error {
	logger.Log.Printf("Intializing DB")
	var err error
	dsn := "root:lmcmysql123@tcp(localhost:3306)/lourde?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Printf("Error Intializing DB")
		return err
	}
	sqlDB, err := db.DB()

	if err != nil {
		logger.Log.Printf("Error Intializing DB")
		return err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	initializeTables()

	return nil

}

var p [7]interface{}

func initializeTables() {

	logger.Log.Print(db)

	migrate := db.Debug().Migrator()

	logger.Log.Printf("Loading Tables")

	p[0] = Notification{}
	p[1] = Event{}
	p[2] = EventImages{}
	p[3] = Student{}
	p[4] = User{}
	p[5] = Marks{}
	p[6] = Calendar{}

	for _, element := range p {
		logger.Log.Printf(" Table is beign created for %T", element)
		if !migrate.HasTable(element) {
			error := migrate.CreateTable(element)
			if error != nil {
				logger.Log.Printf("Cannot create table Schema Database Error- %T", element)
			}
		}

	}
}

//InitializeDefaultData is used to
func InitializeDefaultData() {
	//initialize default users
	master := User{
		Name: "master",
		Hash: "lourde_master",
	}
	admin := User{
		Name: "administrator",
		Hash: "lourde_admin",
	}
	master.Hash, _ = HashUserPassword(master.Hash)
	admin.Hash, _ = HashUserPassword(admin.Hash)
	db.Where(&User{Name: "master"}).FirstOrCreate(&master)
	db.Where(&User{Name: "administrator"}).FirstOrCreate(&admin)

	banner := Notification{

		Text:        "This is a Banner Text",
		Link:        "This is a Banner Link",
		Entity:      "banner",
		Description: "inactive",
		Date:        0,
	}
	db.Where(&Notification{Entity: "banner"}).FirstOrCreate(&banner)

	principle := Event{
		Name:        "Principal",
		Description: "Empty",
		Path:        "Empty",
		Entity:      "principal",
		Date:        0,
	}

	db.Where(&Event{Entity: "principal"}).FirstOrCreate(&principle)

}

//Notification struct model
type Notification struct {
	ID          uint      `json:"id"`
	Text        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"url"`
	Entity      string    `gorm:"not null" json:"entity"`
	Date        int       `json:"date"`
	CreatedAt   time.Time `json:"created_time"`
}

//Event is a bean used to store Event Details
type Event struct {
	ID          uint          `json:"id"`
	Name        string        `gorm:"not null" json:"name"`
	Description string        `json:"description"`
	Path        string        `gorm:"not null" json:"thumbnail"`
	Entity      string        `gorm:"not null" json:"entity"`
	Date        int           `json:"date"`
	EventImages []EventImages `json:"images,omitempty"`
	CreatedAt   time.Time     `json:"created_time"`
}

//Student - struct used to store Student Info
type Student struct {
	ID          uint      `json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Rank        int       `json:"order"`
	Path        string    `gorm:"not null" json:"thumbnail"`
	Entity      string    `gorm:"not null" json:"entity"`
	Marks       []Marks   `json:"marks"`
	CreatedAt   time.Time `json:"created_time"`
}

//Marks is used to
type Marks struct {
	ID        uint   `json:"id"`
	Field     string `json:"field"`
	Value     string `json:"value"`
	StudentID *uint  `json:"entityID"`
}

//Globalinfo is used to
type Globalinfo struct {
	ID        uint   `json:"id"`
	Field     string `json:"field"`
	Value     string `json:"value"`
	Entity    string `json:"entity"`
	StudentID *uint  `json:"entityid"`
}

//EventImages - Used to store Event Images
type EventImages struct {
	ID          uint   `json:"id" default:"1000"`
	Description string `json:"description"`
	Path        string `gorm:"not null" json:"thumbnail"`
	EventID     *uint  `json:"eventid"`
}

//User is used to store used Data
type User struct {
	ID   uint   `json:"id"`
	Name string `gorm:"unique" json:"username"`
	Hash string `json:"password"`
}

//APIUser is used to change passwords
type APIUser struct {
	Name    string `json:"username"`
	Hash    string `json:"password"`
	NewHash string `json:"new_password"`
	Secret  string `json:"secret"`
}

//Calendar is used to change calendar
type Calendar struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Start       uint   `json:"start"`
	End         uint   `json:"end"`
	Color       string `json:"color"`
}

//Table is a interface
type Table interface {
	getTableName() string
}

func (mr *EventImages) getTableName() string {
	return "EventImages"
}

func (mr *Marks) getTableName() string {
	return "Marks"
}

//GetAllEntity used to get All Notification
func GetAllEntity(value interface{}) (*int64, *error) {
	result := db.Find(value)
	return &result.RowsAffected, &result.Error
}

//GetAllEntityFilter used to get All Entity with filter
func GetAllEntityFilter(model interface{}, value interface{}) (*int64, *error) {
	result := db.Where(model).Find(value)
	return &result.RowsAffected, &result.Error
}

//GetAllEntityFilterOrder used to get All Entity with filter
func GetAllEntityFilterOrder(model interface{}, value interface{}, order interface{}) (*int64, *error) {
	result := db.Where(model).Order(order).Find(value)
	return &result.RowsAffected, &result.Error
}

//PostEntity is to Post a Notification
func PostEntity(value interface{}) (*int64, *error) {

	result := db.Create(value)
	return &result.RowsAffected, &result.Error
}

//UpdateEntity is to Update a Notification
func UpdateEntity(value interface{}, update interface{}) (*int64, *error) {

	result := db.Model(value).Updates(update)
	return &result.RowsAffected, &result.Error
}

//DeleteEntity is to Delete a Notification
func DeleteEntity(value interface{}) (*int64, *error) {

	result := db.Delete(value)
	return &result.RowsAffected, &result.Error
}

//GetEntity - get Single Notification based on ID
func GetEntity(value interface{}, id *string) (*int64, *error) {

	result := db.Find(value, id)
	return &result.RowsAffected, &result.Error
}

//PostChildEntity - posting child entity
func PostChildEntity(parent interface{}, child Table) *error {
	result := db.Model(parent).Association(child.getTableName()).Append(child)
	return &result
}

//GetAllChildEntity - posting child entity
func GetAllChildEntity(parent interface{}, child interface{}, table string) *error {
	result := db.Model(parent).Association(table).Find(child)
	return &result
}

//DeleteChildEntity - delete child entity
func DeleteChildEntity(parent interface{}, child Table) *error {
	result := db.Model(parent).Association(child.getTableName()).Delete(child)
	return &result

}

//GetCountEntity is used to get the count of entities
func GetCountEntity(parent interface{}, count *int64) {
	db.Model(parent).Where(parent).Count(count)
}

//HashUserPassword is used to hash the user passwords
func HashUserPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash - checks whether password and hash same
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//IsValidUser is used to validate users
func IsValidUser(user *User) (bool, User) {
	var dbUser User
	db.Where(&User{Name: user.Name}).Find(&dbUser)
	return CheckPasswordHash(user.Hash, dbUser.Hash), dbUser
}

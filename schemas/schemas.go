package schemas

import (
	"time"

	"gorm.io/gorm"
)

// User representa la tabla Users en la base de datos
type User struct {
    gorm.Model
    AccountID        int64
    UserID           int64
    CreateDate       time.Time
    LastUpdated      time.Time
    RecordStatus     string `gorm:"default:app.status.Active"`
    Active           string `gorm:"default:app.status.Active"`
    Identification   string
    Password         string
    CompanyName      string
    FirstName        string
    LastName         string
    Email            string
    Phone            string
    EmergencyPhone   string
    I18n             string
    Address          string
    AuthMenu         string `gorm:"type:text;default:'[]'"`
    AuthKeys         string `gorm:"type:text;default:'[]'"`
    AuthGroups       string `gorm:"type:text;default:'[]'"`
    Sites            string `gorm:"type:text;default:'[]'"`
    Clients          string `gorm:"type:text;default:'[]'"`
}

// Connection representa la tabla Connections en la base de datos
type Connection struct {
    gorm.Model
    ConnectionID   string `gorm:"uniqueIndex"`
    AccountID      int64
    UserID         int64
    CreateDate     time.Time
    LastUpdated    time.Time
    RecordStatus   string `gorm:"default:app.status.Active"`
    Active         string `gorm:"default:app.status.Active"`
    Connected      string
    Disconnected   string
    UserData       string `gorm:"type:text;default:'{}'"`
    AccountData    string `gorm:"type:text;default:'{}'"`
    AuthMenu       string `gorm:"type:text;default:'[]'"`
    AuthKeys       string `gorm:"type:text;default:'[]'"`
    AuthGroups     string `gorm:"type:text;default:'[]'"`
    Sites          string `gorm:"type:text;default:'[]'"`
    Clients        string `gorm:"type:text;default:'[]'"`
}

// Site representa la tabla Sites en la base de datos
type Site struct {
    gorm.Model
    AccountID         int64
    SiteID            int64
    SiteName          string
    CreateDate        time.Time
    LastUpdated       time.Time
    RecordStatus      string `gorm:"default:app.status.Active"`
    Active            string `gorm:"default:app.status.Active"`
    Des               string
    Description       string
    OperateBy         string
    Logo              string
    RulesDocuments    string `gorm:"type:text;default:'{}'"`
    ServicesAmenities string `gorm:"type:text;default:'{}'"`
    Type              string `gorm:"default:Yard"`
    Email             string
    Phone             string
    Address           string
    Website           string
    Geolocation       string
}

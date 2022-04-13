package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Card struct {
	bun.BaseModel `bun:"cc_card"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`    //
	Creationdate  time.Time `bun:"creationdate" json:"creationdate"` //
	// Firstusedate       time.Time `bun:"firstusedate" json:"firstusedate"`                //
	// Expirationdate     time.Time `bun:"expirationdate" json:"expirationdate"`            //
	Enableexpire int64   `bun:"enableexpire,nullzero" json:"enableexpire"` //
	Expiredays   int64   `bun:"expiredays,nullzero" json:"expiredays"`     //
	Username     string  `bun:"username" json:"username"`                  //
	Useralias    string  `bun:"useralias" json:"useralias"`                //
	Uipass       string  `bun:"uipass" json:"uipass"`                      //
	Credit       float64 `bun:"credit" json:"credit"`                      //
	Tariff       int64   `bun:"tariff,nullzero" json:"tariff"`             //
	IDDidgroup   int64   `bun:"id_didgroup,nullzero" json:"id_didgroup"`   //
	Activated    string  `bun:"activated" json:"activated"`                //
	Status       int     `bun:"status" json:"status"`                      //
	Lastname     string  `bun:"lastname" json:"lastname"`                  //
	Firstname    string  `bun:"firstname" json:"firstname"`                //
	Address      string  `bun:"address" json:"address"`                    //
	City         string  `bun:"city" json:"city"`                          //
	State        string  `bun:"state" json:"state"`                        //
	Country      string  `bun:"country" json:"country"`                    //
	Zipcode      string  `bun:"zipcode" json:"zipcode"`                    //
	Phone        string  `bun:"phone" json:"phone"`                        //
	Email        string  `bun:"email" json:"email"`                        //
	Fax          string  `bun:"fax" json:"fax"`                            //
	Inuse        int64   `bun:"inuse,nullzero" json:"inuse"`               //
	Simultaccess int64   `bun:"simultaccess,nullzero" json:"simultaccess"` //
	Currency     string  `bun:"currency,nullzero" json:"currency"`         //
	// Lastuse            time.Time `bun:"lastuse" json:"lastuse"`                          //
	Nbused        int64   `bun:"nbused,nullzero" json:"nbused"`                   //
	Typepaid      int64   `bun:"typepaid,nullzero" json:"typepaid"`               //
	Creditlimit   int64   `bun:"creditlimit,nullzero" json:"creditlimit"`         //
	Voipcall      int64   `bun:"voipcall,nullzero" json:"voipcall"`               //
	SipBuddy      int64   `bun:"sip_buddy,nullzero" json:"sip_buddy"`             //
	IaxBuddy      int64   `bun:"iax_buddy,nullzero" json:"iax_buddy"`             //
	Language      string  `bun:"language,nullzero" json:"language"`               //
	Redial        string  `bun:"redial" json:"redial"`                            //
	Runservice    int64   `bun:"runservice,nullzero" json:"runservice"`           //
	Nbservice     int64   `bun:"nbservice,nullzero" json:"nbservice"`             //
	IDCampaign    int64   `bun:"id_campaign,nullzero" json:"id_campaign"`         //
	NumTrialsDone int64   `bun:"num_trials_done,nullzero" json:"num_trials_done"` //
	Vat           float32 `bun:"vat" json:"vat"`                                  //
	// Servicelastrun     time.Time `bun:"servicelastrun" json:"servicelastrun"`            //
	Initialbalance     float64 `bun:"initialbalance" json:"initialbalance"`           //
	Invoiceday         int64   `bun:"invoiceday,nullzero" json:"invoiceday"`          //
	Autorefill         int64   `bun:"autorefill,nullzero" json:"autorefill"`          //
	Loginkey           string  `bun:"loginkey" json:"loginkey"`                       //
	MacAddr            string  `bun:"mac_addr" json:"mac_addr"`                       //
	IDTimezone         int64   `bun:"id_timezone,nullzero" json:"id_timezone"`        //
	Tag                string  `bun:"tag" json:"tag"`                                 //
	VoicemailPermitted int     `bun:"voicemail_permitted" json:"voicemail_permitted"` //
	VoicemailActivated int     `bun:"voicemail_activated" json:"voicemail_activated"` //
	// LastNotification   time.Time `bun:"last_notification" json:"last_notification"`      //
	EmailNotification  string  `bun:"email_notification" json:"email_notification"`   //
	NotifyEmail        int     `bun:"notify_email" json:"notify_email"`               //
	CreditNotification int     `bun:"credit_notification" json:"credit_notification"` //
	IDGroup            int     `bun:"id_group" json:"id_group"`                       //
	CompanyName        string  `bun:"company_name" json:"company_name"`               //
	CompanyWebsite     string  `bun:"company_website" json:"company_website"`         //
	VatRn              string  `bun:"vat_rn,nullzero" json:"vat_rn"`                  //
	Traffic            int64   `bun:"traffic,nullzero" json:"traffic"`                //
	TrafficTarget      string  `bun:"traffic_target" json:"traffic_target"`           //
	Discount           float64 `bun:"discount" json:"discount"`                       //
	Restriction        int     `bun:"restriction" json:"restriction"`                 //
	IDSeria            int64   `bun:"id_seria,nullzero" json:"id_seria"`              //
	Serial             int64   `bun:"serial,nullzero" json:"serial"`                  //
	Block              int     `bun:"block" json:"block"`                             //
	LockPin            string  `bun:"lock_pin,nullzero" json:"lock_pin"`              //
	// LockDate           time.Time `bun:"lock_date" json:"lock_date"`                      //
	MaxConcurrent int    `bun:"max_concurrent" json:"max_concurrent"` //
	APIKey        string `bun:"api_key,nullzero" json:"api_key"`      //
	// CallerIds          *CallerId `bun:"rel:belongs-to" json:"caller_ids"`
}

type CardInfo struct {
	bun.BaseModel `bun:"cc_card,alias:c"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`    //
	Creationdate  time.Time `bun:"creationdate" json:"creationdate"` //
	// Firstusedate   time.Time `bun:"firstusedate" json:"firstusedate"`          //
	// Expirationdate time.Time `bun:"expirationdate" json:"expirationdate"`      //
	Enableexpire int64   `bun:"enableexpire,nullzero" json:"enableexpire"` //
	Expiredays   int64   `bun:"expiredays,nullzero" json:"expiredays"`     //
	Username     string  `bun:"username" json:"username"`                  //
	Useralias    string  `bun:"useralias" json:"useralias"`                //
	Credit       float64 `bun:"credit" json:"credit"`                      //
	Activated    string  `bun:"activated" json:"activated"`                //
	Status       int     `bun:"status" json:"status"`                      //
	// Lastuse        time.Time `bun:"lastuse" json:"lastuse"`                    //
	// Creditlimit    int64     `bun:"creditlimit,nullzero" json:"creditlimit"`   //
	IDGroup int   `bun:"id_group" json:"id_group"`         //
	Tariff  int64 `bun:"tariff,nullzero" json:"call_plan"` //
	// CallerIds      []CallerId `bun:"foreignKey:IDCcCard ,references:ID" json:"caller_ids"` //
	CallerId []*CallerId `bun:"rel:has-many,join:id=id_cc_card" json:"caller_ids"`
}

// TableName sets the insert table name for this struct type

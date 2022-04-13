package model

import (
	"github.com/uptrace/bun"
)

type SipBuddies struct {
	bun.BaseModel     `bun:"cc_sip_buddies"`
	ID                int    `bun:"id,pk" json:"id"`                                     //
	IDCcCard          int    `bun:"id_cc_card" json:"id_cc_card"`                        //
	Name              string `bun:"name" json:"name"`                                    //
	Accountcode       string `bun:"accountcode" json:"accountcode"`                      //
	Regexten          string `bun:"regexten" json:"regexten"`                            //
	Amaflags          string `bun:"amaflags,nullzero" json:"amaflags"`                   //
	Callgroup         string `bun:"callgroup,nullzero" json:"callgroup"`                 //
	Callerid          string `bun:"callerid" json:"callerid"`                            //
	Canreinvite       string `bun:"canreinvite" json:"canreinvite"`                      //
	Context           string `bun:"context" json:"context"`                              //
	DEFAULTip         string `bun:"DEFAULTip,nullzero" json:"DEFAULTip"`                 //
	Dtmfmode          string `bun:"dtmfmode" json:"dtmfmode"`                            //
	Fromuser          string `bun:"fromuser" json:"fromuser"`                            //
	Fromdomain        string `bun:"fromdomain" json:"fromdomain"`                        //
	Host              string `bun:"host" json:"host"`                                    //
	Insecure          string `bun:"insecure" json:"insecure"`                            //
	Language          string `bun:"language,nullzero" json:"language"`                   //
	Mailbox           string `bun:"mailbox" json:"mailbox"`                              //
	Md5secret         string `bun:"md5secret" json:"md5secret"`                          //
	Nat               string `bun:"nat,nullzero" json:"nat"`                             //
	Deny              string `bun:"deny" json:"deny"`                                    //
	Permit            string `bun:"permit,nullzero" json:"permit"`                       //
	Mask              string `bun:"mask" json:"mask"`                                    //
	Pickupgroup       string `bun:"pickupgroup,nullzero" json:"pickupgroup"`             //
	Port              string `bun:"port" json:"port"`                                    //
	Qualify           string `bun:"qualify,nullzero" json:"qualify"`                     //
	Restrictcid       string `bun:"restrictcid,nullzero" json:"restrictcid"`             //
	Rtptimeout        string `bun:"rtptimeout,nullzero" json:"rtptimeout"`               //
	Rtpholdtimeout    string `bun:"rtpholdtimeout,nullzero" json:"rtpholdtimeout"`       //
	Secret            string `bun:"secret" json:"secret"`                                //
	Type              string `bun:"type" json:"type"`                                    //
	Username          string `bun:"username" json:"username"`                            //
	Disallow          string `bun:"disallow" json:"disallow"`                            //
	Allow             string `bun:"allow" json:"allow"`                                  //
	Musiconhold       string `bun:"musiconhold" json:"musiconhold"`                      //
	Regseconds        int    `bun:"regseconds" json:"regseconds"`                        //
	Ipaddr            string `bun:"ipaddr" json:"ipaddr"`                                //
	Cancallforward    string `bun:"cancallforward,nullzero" json:"cancallforward"`       //
	Fullcontact       string `bun:"fullcontact" json:"fullcontact"`                      //
	Setvar            string `bun:"setvar" json:"setvar"`                                //
	Regserver         string `bun:"regserver,nullzero" json:"regserver"`                 //
	Lastms            string `bun:"lastms,nullzero" json:"lastms"`                       //
	Defaultuser       string `bun:"defaultuser" json:"defaultuser"`                      //
	Auth              string `bun:"auth" json:"auth"`                                    //
	Subscribemwi      string `bun:"subscribemwi" json:"subscribemwi"`                    //
	Vmexten           string `bun:"vmexten" json:"vmexten"`                              //
	CidNumber         string `bun:"cid_number" json:"cid_number"`                        //
	Callingpres       string `bun:"callingpres" json:"callingpres"`                      //
	Usereqphone       string `bun:"usereqphone" json:"usereqphone"`                      //
	Incominglimit     string `bun:"incominglimit" json:"incominglimit"`                  //
	Subscribecontext  string `bun:"subscribecontext" json:"subscribecontext"`            //
	Musicclass        string `bun:"musicclass" json:"musicclass"`                        //
	Mohsuggest        string `bun:"mohsuggest" json:"mohsuggest"`                        //
	Allowtransfer     string `bun:"allowtransfer" json:"allowtransfer"`                  //
	Autoframing       string `bun:"autoframing" json:"autoframing"`                      //
	Maxcallbitrate    string `bun:"maxcallbitrate" json:"maxcallbitrate"`                //
	Outboundproxy     string `bun:"outboundproxy" json:"outboundproxy"`                  //
	Rtpkeepalive      string `bun:"rtpkeepalive" json:"rtpkeepalive"`                    //
	Useragent         string `bun:"useragent,nullzero" json:"useragent"`                 //
	Callbackextension string `bun:"callbackextension,nullzero" json:"callbackextension"` //
}

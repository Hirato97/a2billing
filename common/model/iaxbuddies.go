package model

import (
	"github.com/uptrace/bun"
)

type IaxBuddies struct {
	bun.BaseModel              `bun:"cc_iax_buddies"`
	ID                         int    `bun:"id,pk" json:"id"`                                                //
	IDCcCard                   int    `bun:"id_cc_card" json:"id_cc_card"`                                   //
	Name                       string `bun:"name" json:"name"`                                               //
	Accountcode                string `bun:"accountcode" json:"accountcode"`                                 //
	Regexten                   string `bun:"regexten" json:"regexten"`                                       //
	Amaflags                   string `bun:"amaflags,nullzero" json:"amaflags"`                              //
	Callerid                   string `bun:"callerid" json:"callerid"`                                       //
	Context                    string `bun:"context" json:"context"`                                         //
	DEFAULTip                  string `bun:"DEFAULTip,nullzero" json:"DEFAULTip"`                            //
	Host                       string `bun:"host" json:"host"`                                               //
	Language                   string `bun:"language,nullzero" json:"language"`                              //
	Deny                       string `bun:"deny" json:"deny"`                                               //
	Permit                     string `bun:"permit,nullzero" json:"permit"`                                  //
	Mask                       string `bun:"mask" json:"mask"`                                               //
	Port                       string `bun:"port" json:"port"`                                               //
	Qualify                    string `bun:"qualify,nullzero" json:"qualify"`                                //
	Secret                     string `bun:"secret" json:"secret"`                                           //
	Type                       string `bun:"type" json:"type"`                                               //
	Username                   string `bun:"username" json:"username"`                                       //
	Disallow                   string `bun:"disallow" json:"disallow"`                                       //
	Allow                      string `bun:"allow" json:"allow"`                                             //
	Regseconds                 int    `bun:"regseconds" json:"regseconds"`                                   //
	Ipaddr                     string `bun:"ipaddr" json:"ipaddr"`                                           //
	Trunk                      string `bun:"trunk,nullzero" json:"trunk"`                                    //
	Dbsecret                   string `bun:"dbsecret" json:"dbsecret"`                                       //
	Regcontext                 string `bun:"regcontext" json:"regcontext"`                                   //
	Sourceaddress              string `bun:"sourceaddress" json:"sourceaddress"`                             //
	Mohinterpret               string `bun:"mohinterpret" json:"mohinterpret"`                               //
	Mohsuggest                 string `bun:"mohsuggest" json:"mohsuggest"`                                   //
	Inkeys                     string `bun:"inkeys" json:"inkeys"`                                           //
	Outkey                     string `bun:"outkey" json:"outkey"`                                           //
	CidNumber                  string `bun:"cid_number" json:"cid_number"`                                   //
	Sendani                    string `bun:"sendani" json:"sendani"`                                         //
	Fullname                   string `bun:"fullname" json:"fullname"`                                       //
	Auth                       string `bun:"auth" json:"auth"`                                               //
	Maxauthreq                 string `bun:"maxauthreq" json:"maxauthreq"`                                   //
	Encryption                 string `bun:"encryption" json:"encryption"`                                   //
	Transfer                   string `bun:"transfer" json:"transfer"`                                       //
	Jitterbuffer               string `bun:"jitterbuffer" json:"jitterbuffer"`                               //
	Forcejitterbuffer          string `bun:"forcejitterbuffer" json:"forcejitterbuffer"`                     //
	Codecpriority              string `bun:"codecpriority" json:"codecpriority"`                             //
	Qualifysmoothing           string `bun:"qualifysmoothing" json:"qualifysmoothing"`                       //
	Qualifyfreqok              string `bun:"qualifyfreqok" json:"qualifyfreqok"`                             //
	Qualifyfreqnotok           string `bun:"qualifyfreqnotok" json:"qualifyfreqnotok"`                       //
	Timezone                   string `bun:"timezone" json:"timezone"`                                       //
	Adsi                       string `bun:"adsi" json:"adsi"`                                               //
	Setvar                     string `bun:"setvar" json:"setvar"`                                           //
	Requirecalltoken           string `bun:"requirecalltoken" json:"requirecalltoken"`                       //
	Maxcallnumbers             string `bun:"maxcallnumbers" json:"maxcallnumbers"`                           //
	MaxcallnumbersNonvalidated string `bun:"maxcallnumbers_nonvalidated" json:"maxcallnumbers_nonvalidated"` //
}

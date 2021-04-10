package httpfinger

import "strings"

type keywordFinger []struct {
	Type    string `json:"type"`
	Cms     string `json:"cms"`
	Keyword string `json:"keyword"`
}

var KeywordFinger keywordFinger

func (k keywordFinger) Match(body string, header string) string {
	for _, kSub := range k {
		if kSub.Type == "body" {
			if strings.Contains(body, kSub.Keyword) {
				return kSub.Cms
			}
		}
		if kSub.Type == "header" {
			if strings.Contains(header, kSub.Keyword) {
				return kSub.Cms
			}
		}
	}
	return ""
}

var keywordFingerByte = []byte(`
[
    {
      "type": "body",
      "cms": "seeyon",
      "keyword": "/seeyon/USER-DATA/IMAGES/LOGIN/login.gif"
    },
    {
      "type": "body",
      "cms": "seeyon",
      "keyword": "/seeyon/common/"
    },
    {
      "type": "body",
      "cms": "Spring env",
      "keyword": "servletContextInitParams"
    },
    {
      "type": "body",
      "cms": "Spring env",
      "keyword": "logback"
    },
    {
      "type": "body",
      "cms": "Weblogic",
      "keyword": "Error 404--Not Found"
    },
    {
      "type": "body",
      "cms": "Weblogic",
      "keyword": "Error 403--"
    },
    {
      "type": "body",
      "cms": "Weblogic",
      "keyword": "/console/framework/skins/wlsconsole/images/login_WebLogic_branding.png"
    },
    {
      "type": "body",
      "cms": "Weblogic",
      "keyword": "Welcome to Weblogic Application Server"
    },
    {
      "type": "body",
      "cms": "Weblogic",
      "keyword": "<i>Hypertext Transfer Protocol -- HTTP/1.1</i>"
    },
    {
      "type": "body",
      "cms": "Sangfor SSL VPN",
      "keyword": "/por/login_psw.csp"
    },
    {
      "type": "body",
      "cms": "Sangfor SSL VPN",
      "keyword": "loginPageSP/loginPrivacy.js"
    },
    {
      "type": "body",
      "cms": "e-mobile",
      "keyword": "weaver,e-mobile"
    },
    {
      "type": "header",
      "cms": "ecology",
      "keyword": "ecology_JSessionid"
    },
    {
      "type": "header",
      "cms": "Shiro",
      "keyword": "rememberMe="
    },
    {
      "type": "header",
      "cms": "Shiro",
      "keyword": "=deleteMe"
    },
    {
      "type": "body",
      "cms": "e-Bridge",
      "keyword": "wx.weaver"
    },
    {
      "type": "body",
      "cms": "e-Bridge",
      "keyword": "e-Bridge"
    },
    {
      "type": "body",
      "cms": "Swagger UI",
      "keyword": "Swagger UI"
    },
    {
      "type": "body",
      "cms": "Ruijie",
      "keyword": "4008 111 000"
    },
    {
      "type": "body",
      "cms": "Huawei SMC",
      "keyword": "Script/SmcScript.js?version="
    },
    {
      "type": "body",
      "cms": "H3C Router",
      "keyword": "/wnm/ssl/web/frame/login.html"
    },
    {
      "type": "body",
      "cms": "Cisco SSLVPN",
      "keyword": "/+CSCOE+/logon.html"
    },
    {
      "type": "body",
      "cms": "\u901a\u8fbeOA",
      "keyword": "/images/tongda.ico"
    },
    {
      "type": "body",
      "cms": "\u901a\u8fbeOA",
      "keyword": "Office Anywhere"
    },
    {
      "type": "body",
      "cms": "\u901a\u8fbeOA",
      "keyword": "\u901a\u8fbeOA"
    },
    {
      "type": "body",
      "cms": "\u6df1\u4fe1\u670d waf",
      "keyword": "rsa.js"
    },
    {
      "type": "body",
      "cms": "\u6df1\u4fe1\u670d waf",
      "keyword": "Redirect to..."
    },
    {
      "type": "body",
      "cms": "\u7f51\u5fa1 vpn",
      "keyword": "/vpn/common/js/leadsec.js"
    },
    {
      "type": "body",
      "cms": "\u542f\u660e\u661f\u8fb0\u5929\u6e05\u6c49\u9a6cUSG\u9632\u706b\u5899",
      "keyword": "/cgi-bin/webui?op=get_product_model"
    },
    {
      "type": "body",
      "cms": "\u84dd\u51cc OA",
      "keyword": "sys/ui/extend/theme/default/style/icon.css"
    },
    {
      "type": "body",
      "cms": "\u6df1\u4fe1\u670d\u4e0a\u7f51\u884c\u4e3a\u7ba1\u7406\u7cfb\u7edf",
      "keyword": "utccjfaewjb = function(str, key)"
    },
    {
      "type": "body",
      "cms": "\u6df1\u4fe1\u670d\u4e0a\u7f51\u884c\u4e3a\u7ba1\u7406\u7cfb\u7edf",
      "keyword": "document.write(WRFWWCSFBXMIGKRKHXFJ"
    },
    {
      "type": "body",
      "cms": "\u6df1\u4fe1\u670d\u5e94\u7528\u4ea4\u4ed8\u62a5\u8868\u7cfb\u7edf",
      "keyword": "/reportCenter/index.php?cls_mode=cluster_mode_others"
    },
    {
      "type": "body",
      "cms": "\u91d1\u8776\u4e91\u661f\u7a7a",
      "keyword": "HTML5/content/themes/kdcss.min.css"
    },
    {
      "type": "body",
      "cms": "\u91d1\u8776\u4e91\u661f\u7a7a",
      "keyword": "/ClientBin/Kingdee.BOS.XPF.App.xap"
    },
    {
      "type": "body",
      "cms": "CoreMail",
      "keyword": "coremail/common"
    },
    {
      "type": "body",
      "cms": "\u542f\u660e\u661f\u8fb0\u5929\u6e05\u6c49\u9a6cUSG\u9632\u706b\u5899",
      "keyword": "\u5929\u6e05\u6c49\u9a6cUSG"
    },
    {
      "type": "body",
      "cms": "Jboss",
      "keyword": "jboss.css"
    },
    {
      "type": "body",
      "cms": "Gitlab",
      "keyword": "assets/gitlab_logo"
    },
    {
      "type": "body",
      "cms": "\u5b9d\u5854-BT.cn",
      "keyword": "\u5165\u53e3\u6821\u9a8c\u5931\u8d25"
    },
    {
      "type": "body",
      "cms": "\u7985\u9053",
      "keyword": "self.location"
    },
    {
      "type": "body",
      "cms": "\u7985\u9053",
      "keyword": "/theme/default/images/main/zt-logo.png"
    },
    {
      "type": "ADSL/Router",
      "cms": "\u7985\u9053",
      "keyword": "zentaosid"
    },
    {
      "type": "body",
      "cms": "\u7528\u53cb\u8f6f\u4ef6",
      "keyword": "UFIDA Software CO.LTD all rights reserved"
    },
    {
      "type": "body",
      "cms": "YONYOU NC",
      "keyword": "uclient.yonyou.com"
    },
    {
      "type": "body",
      "cms": "\u5b9d\u5854-BT.cn",
      "keyword": "\u5b9d\u5854Linux\u9762\u677f"
    },
    {
      "type": "body",
      "cms": "RabbitMQ",
      "keyword": "<title>RabbitMQ Management</title>"
    },
    {
      "type": "body",
      "cms": "Zabbix",
      "keyword": "zabbix"
    },
    {
      "type": "body",
      "cms": "\u8054\u8f6f\u51c6\u5165",
      "keyword": "\u7f51\u7edc\u51c6\u5165"
    },
    {
      "type": "body",
      "cms": "\u5217\u76ee\u5f55",
      "keyword": "Index of /"
    },
    {
      "type": "body",
      "cms": "\u5217\u76ee\u5f55",
      "keyword": " - /</title>"
    },
    {
      "type": "body",
      "cms": "\u6d6a\u6f6e\u670d\u52a1\u5668IPMI\u7ba1\u7406\u53e3",
      "keyword": "img/inspur_logo.png"
    },
    {
      "type": "body",
      "cms": "RegentApi_v2.0",
      "keyword": "RegentApi_v2.0"
    },
    {
      "type": "body",
      "cms": "Tomcat\u9ed8\u8ba4\u9875\u9762",
      "keyword": "/manager/status"
    }
]
`)

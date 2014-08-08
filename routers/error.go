package routers

import (
    "html/template"
    "net/http"
)

var errtpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <title>{{.Title}}</title>
        <link rel="shortcut icon" href="/static/img/doolp_favicon_01.png" />
        <style type="text/css">
            * {
                margin:0;
                padding:0;
            }

            body {
                background-color:#EFEFEF;
                font: .9em "Lucida Sans Unicode", "Lucida Grande", sans-serif;
            }

            #wrapper{
                width:600px;
                margin:40px auto 0;
                text-align:center;
                -moz-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
                -webkit-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
                box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
            }

            #wrapper h1{
                color:#FFF;
                text-align:center;
                margin-bottom:20px;
            }

            #wrapper a{
                display:block;
                font-size:.9em;
                padding-top:20px;
                color:#FFF;
                text-decoration:none;
                text-align:center;
            }

            #container {
                width:600px;
                padding-bottom:15px;
                background-color:#FFFFFF;
            }

            .navtop{
                height:40px;
                background-color:#24B2EB;
                padding:13px;
            }

            .content {
                padding:10px 10px 25px;
                background: #FFFFFF;
                margin:;
                color:#333;
            }

            a.button{
                color:white;
                padding:15px 20px;
                text-shadow:1px 1px 0 #00A5FF;
                font-weight:bold;
                text-align:center;
                border:1px solid #24B2EB;
                margin:0px 200px;
                clear:both;
                background-color: #24B2EB;
                border-radius:100px;
                -moz-border-radius:100px;
                -webkit-border-radius:100px;
            }

            a.button:hover{
                text-decoration:none;
                background-color: #24B2EB;
            }

        </style>
    </head>
    <body>
        <div id="wrapper">
            <div id="container">
                <div class="navtop">
                    <h1>{{.Title}}</h1>
                </div>
                <div id="content">
                    {{.Content}}
                    <a href="/" title="Home" class="button">Go Home</a><br />

                    <br>Powered by Doolp {{.DoolpVersion}}
                    <a class="hidden-xs logo" href="/">
                    <img style="height:38px;width:102px;" src="/static/img/doolp_logo_01.png">
                    </a>
                </div>
            </div>
        </div>
    </body>
</html>
`

func Page_not_found(rw http.ResponseWriter, r *http.Request) {
    t, _ := template.New("doolptemp").Parse(errtpl)
    data := make(map[string]interface{})
    data["Title"] = "Page Not Found"
    data["Content"] = template.HTML("<br>The page you have requested has flown the coop." +
        "<br>Perhaps you are here because:" +
        "<br><br><ul>" +
        "<br>The page has moved" +
        "<br>The page no longer exists" +
        "<br>You were looking for your puppy and got lost" +
        "<br>You like 404 pages" +
        "</ul>")
    data["DoolpVersion"] = "0.1.1"
    //rw.WriteHeader(http.StatusNotFound)
    t.Execute(rw, data)
}

var request = new XMLHttpRequest()
function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return null;
}
const parseJwt = (token) => {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return null; }
};
-
Cookie = getCookie("ID")
if (Cookie != null) {
    console.log("A")
    request.open('GET', `http://192.168.0.72:8080/user/isvalid/${Cookie}`, true)
    request.addEventListener("readystatechange", function() {
        if (request.status == 204) { //Change to 200 when ready
            window.location = `http://localhost:5500/Frontend/dash.html`;
        }
    })
    request.send()
}

function handleCredentialResponse(response) {
    var parsedJwt = parseJwt(response.credential)
    if (parsedJwt.hd == "redborne.com") {
        request.open('GET', `http://192.168.0.72:8080/user/isvalid/${parsedJwt.sub}`, true)
        request.addEventListener("readystatechange", function() {
            if (request.status == 204) { //Change to 200 when ready
                window.location = `http://localhost:5500/Frontend/signup.html?JWT=${response.credential}`;
            } else if (request.status == 200) {
                document.cookie = `ID=${parsedJwt.sub}`
                window.location = `http://localhost:5500/Frontend/dash.html`;
            } else {
                console.log("Error")
            }
        })
        request.send()
    } else {
        window.alert("Please use your school account to login");
    }   
}
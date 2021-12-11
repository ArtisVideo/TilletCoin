var request = new XMLHttpRequest()
const parseJwt = (token) => {
try {
    return JSON.parse(atob(token.split('.')[1]));
} catch (e) {
    return null; }
};

function handleCredentialResponse(response) {
    parsedJwt = parseJwt(response.credential)
    if (parsedJwt.hd == "redborne.com") {
        request.open('GET', `http://192.168.0.72:8080/user/isvalid/${parsedJwt.sub}`, true)
        console.log(request.status)
        request.addEventListener("readystatechange", function() {
            if (request.status == 204) {
                console.log("Create account")
            } else if (request.status == 200) {
                window.location = `http://localhost:5500/Frontend/dash.html?ID=${parsedJwt.sub}`;
            } else {
                console.log("Error")
            }
        })
        request.send()
    } else {
        window.alert("Please use your school account to login");
    }
    
}
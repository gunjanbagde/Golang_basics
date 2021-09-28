function login(){
    var usrname=document.getElementById("uname").value;
    var pass=document.getElementById("password").value;

    if(usrname=="gunjan"&&pass=="1234"){
        var path="file:///C:/WebDev/second.html"
        location.assign(path);
       // window.alert("Successfully Logged in")
    }
    else{
        window.alert("Login Failed Please Try Again")
    }
}       
function openNav() {
    document.getElementById("mySidebar").style.width = "250px";
    document.getElementById("main").style.marginLeft = "250px";
    }

function closeNav() {
    document.getElementById("mySidebar").style.width = "0";
    document.getElementById("main").style.marginLeft = "0";
    }


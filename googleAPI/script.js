const button = document.getElementById('inputForm');

var lat
var long
var cor
var k = 0

button.addEventListener("submit", (e)=>{

//prevent auto submission
    e.preventDefault()

    
    fetch("http://127.0.0.1:4000/start", {
        method:"POST",
    }).then(
        res => res.text()
    ).then(
        (data) => {
            console.log(data);
            cor = data
            var str = cor.split(' ')
            lat = str[0]
            long = str[1]

            localStorage.setItem("lat",lat);
            localStorage.setItem("long",long)
            location.replace("index.html")
            console.log(data);
            
        }
  
    ).catch(
    error => console.error(error)
    

    )



});
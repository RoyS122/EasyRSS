function sendFormNewRSS() {
    
    var Addr = document.getElementById("InputAddressRSSFeed")
    var Name = document.getElementById("InputNameRSSFeed")
    var Vers = document.getElementById("InputVersionRSSFeed")
console.log(Addr.value, Name.value, Vers.value) 
    fetch("/addRSS", {
        method: 'POST',
        headers: {
            'content-type': 'application/json'
        },
        body: JSON.stringify({
            Name: Name.value,
            Link: Addr.value,
            Version: Vers.value
        }),
    }).then((v) => {
        if(v.ok) {
            displayRSSFeeds()
        }
        
    })
    
}


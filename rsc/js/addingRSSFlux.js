function sendFormNewRSS() {
    
    var Addr = document.getElementById("InputAddressRSSFeed")
    var Name = document.getElementById("InputNameRSSFeed")
console.log(Addr.value, Name.value,)
    fetch("/addRSS", {
        method: 'POST',
        headers: {
            'content-type': 'application/json'
        },
        body: JSON.stringify({
            Name: Name.value,
            Link: Addr.value
        }),
    }).then((v) => {
        if(v.ok) {
            displayRSSFeeds()
        }
        
    })
    
}


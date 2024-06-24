function sendFormNewRSS() {
    
    let Addr = document.getElementById("InputAddressRSSFeed")
    let Name = document.getElementById("InputNameRSSFeed")
    let Vers = document.getElementById("InputVersionRSSFeed")


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


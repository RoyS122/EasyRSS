async function displayRSSFeeds() {
    var container = document.getElementsByClassName("containerRSSFeed")[0]
    let containers = container.getElementsByClassName("containerRSSFlux")
    for(let i = containers.length - 1; i >= 0; i --) {
        containers[i].remove()
    }
    fetch("/getAllRSSFeeds", {
        method: 'GET'
    }).then((r) => {
       
        //console.log(r.json())
       // console.log(r.json())
        return r.json()
    
     }).then((r) => {
        
        r.forEach((flux, i) => {
            let fContainer = document.createElement("div")
            fContainer.className = "card"
            fContainer.id = "containerRSSFlux"
            let fLink = document.createElement("a")
            let delButton = document.createElement("button")
            delButton.onclick = () => {removeRSS(i)}
            delButton.textContent = "Delete"
            
            fLink.href = `/viewRSS/${i}`
            fLink.textContent = flux.Name
            fContainer.appendChild(fLink)
            fContainer.appendChild(delButton)
            container.appendChild(fContainer)
            
        });
    })
}

function removeRSS(id) {
    console.log(id)
    fetch("/deleteRSSFeed", {
        method: 'POST',
        headers: {
            'content-type': 'application/json'
        },
        body: JSON.stringify({
            Id: id
        }),
    }).then((r) => {
        if (r.ok) {
            return r 
        }
    }).then((r) => {
        console.log(r)
        displayRSSFeeds()
    })
    
}

document.addEventListener('DOMContentLoaded', function () {
    displayRSSFeeds();
});
function addCategorie() {
      
    var Name = document.getElementById("inputCategorieName")
    var Col = document.getElementById("inputCategorieColor")
    var Description = document.getElementById("inputCategorieDescription")
console.log(Col.value, Description.value, Name.value,)
    fetch("/addCategorie", {
        method: 'POST',
        headers: {
            'content-type': 'application/json'
        },
        body: JSON.stringify({
            Title: Name.value,
            Description: Description.value,
            Color: Col.value
            
        }),
    })
}
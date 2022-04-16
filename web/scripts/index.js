var visitCountContainer = document.querySelector(".visit-counter");

fetch('http://localhost:8090/visits', { 
    method: 'GET',
    mode: 'cors',
    cache: 'no-cache'
})
  .then( response => {
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    return response.json();
  })
  .then( json => { visitCountContainer.innerHTML = json.Visits } )
  .catch( err => console.error(`Fetch problem: ${err.message}`) );

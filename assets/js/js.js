function deleteRecord(recordId) {
    const url = window.location.protocol + "//" + window.location.hostname +":" + window.location.port + "/records/" + recordId 
    console.log(url)
fetch(url, {
  method: 'DELETE',
})
.then(res => res.json()) 
.then((res) => {
    if(res["data"]) {
        $("#" + recordId).hide();
    } else {
    alert("Something went wrong, couldn't delete the record");
    }
        })
}

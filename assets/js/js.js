function listSelectedFiles() {
    var selectedFiles = "<br />";
    const files = $('#inputFiles').prop('files');
for (i = 0; i < files.length; i++) {
    selectedFiles += "<i>" + files[i].name + "</i><br />";
}
    $('#upload-file-info').html(selectedFiles);
}

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

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
fetch(url, {
  method: 'DELETE',
})
.then(res => res.json()) 
.then((res) => {
    if(res["deleted"]) {
        $("#" + recordId).hide();
    } else {
    alert("Something went wrong, couldn't delete the record");
    }
        })
}

function deleteFile(fileId) {
    const url = window.location.protocol + "//" + window.location.hostname +":" + window.location.port + "/files/" + fileId 
fetch(url, {
  method: 'DELETE',
})
.then(res => res.json()) 
.then((res) => {
    if(res["deleted"]) {
        $("#" + fileId).hide();
    } else {
    alert("Something went wrong, couldn't delete the file");
    }
        })
}

function editForm() {
  var data = new FormData($("#editForm")[0])
  fetch('/records/edit', {
    method: 'PUT',
    body: data
  })
        .then(res => {
            if (!res.ok) {
                throw new Error("Error during request.");
            }
            return res.json()})
    .then((res) => {
        alert("Record updated.");
        location.reload();
     })
        .catch((err) => {
        alert(err);
        })
};

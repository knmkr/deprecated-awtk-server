// This file is required by the index.html file and will
// be executed in the renderer process for that window.
// All of the Node.js APIs are available in this process.

const {dialog} = require('electron').remote;
var request = require('superagent');
var path = require("path");

// TODO: DRY
var addr = "localhost:1323";

document.getElementById('select-file').addEventListener('click',function(){
    console.log(dialog.showOpenDialog({properties: ['openFile', 'openDirectory', 'multiSelections']}));
});

document.getElementById('get-genotype-btn').addEventListener('click',function(){
    uri = path.join(addr, "/v1/genomes/1/genotypes");
    request
        .get(uri)
        .query({locations: "20-14369"})
        .end(function(err, res){
            prev = document.getElementById('get-genotype-res');
            if (prev) {
                prev.parentNode.removeChild(prev);
            }

            var div = document.createElement('div');
            div.className = 'well';
            div.id = 'get-genotype-res';
            div.textContent = res.text;
            node = document.getElementById('get-genotype-code');
            node.parentNode.insertBefore(div, node.nextSibling);
        });
});

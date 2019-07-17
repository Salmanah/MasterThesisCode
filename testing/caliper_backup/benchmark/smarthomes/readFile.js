'use strict'

var fs = require("fs")


function readFile(){
    let data, err = fs.readFileSync("file.txt","utf8");

    if(err){
        console.log(err); 
        process.exit(1); 
    }

    return data; 
}

module.exports.readFile = readFile; 
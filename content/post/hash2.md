---
title: "Hash v.2"
date: 2007-10-03T13:36:55-04:00
slug: hash2
tags: [tools]
written: ["2007","2007-10","2007-10-03"]
---


<script src="/js/jquery.min.js" ></script>
<script type="text/javascript" src="/js/hashes.min.js"></script>
<script src="/js/aes.js"></script>
<script src="/js/seedrandom.min.js">
</script>
<script>
var SHA512 = new Hashes.SHA512;
function display(text) {
    $('#form').fadeOut('slow', function() {
        $('#outputBox').text(text);
        $('#outputContainer').fadeIn('slow');
    });
};
function unpack(str) {
    var bytes = [];
    for(var i = 0; i < str.length; i++) {
        var char = str.charCodeAt(i);
        bytes.push(char >>> 8);
        bytes.push(char & 0xFF);
    }
    return bytes;
}
function pack(bytes) {
    var str = "";
    for(var i = 0; i < bytes.length; i += 2) {
        var char = bytes[i] << 8;
        if (bytes[i + 1])
            char |= bytes[i + 1];
        str += String.fromCharCode(char);
    }
    return str;
}
function makeid(seed)
{
    Math.seedrandom(seed);
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.!?";
    for( var i=0; i < 10; i++ )
        text += possible.charAt(Math.floor(Math.random() * possible.length));
    return text;
}
$(document).ready(function() {
    $('#outputContainer').hide();
    $('#outputContainer').focus(function () { // select text on focus
        $(this).select();
    });
    //needs to be inside here os jquery is defined?
    show = function(text) {
        $('#form').fadeOut('slow', function() {
            $('#outputBox').text(text);
            $('#outputContainer').fadeIn('slow');
        });
    };
    $(window).keypress(function(e) {
        if(e.which == 13) {
            $(this).blur();
            $('#submit').focus().click();
        }
    });
    $('#encode').click(function() {
        msg = $('#inputtext').val().toLowerCase();
        passphrase = $('#passphrase').val();
        enc = SHA512.hex(msg+passphrase);
        show(makeid(enc))
    });
    $('#another').click(function() {
        $('#inputtext').val("");
        $('#passphrase').val("");
        $('#outputContainer').hide();   
        $('#form').fadeIn('slow');  
    });
    });
</script>

<div class="container-fluid">
    <div id="form">
<center>
<form>
  <div class="form-group">
    <input style="border: 1px solid transparent; font-size:18pt;height:40px;width:100%;" type="text" class="form-control input-lg" id="inputtext" placeholder="Message">
  </div>
  <div class="form-group">
    <input type="password" style="border: 1px solid transparent; font-size:18pt;height:40px;width:100%;" class="form-control input-lg" id="passphrase" placeholder="Salt">
  </div>
                <button  style="font-size:18pt;height:40px;width:100%;" type="button" id="encode" value="encrypt" class="btn btn-default btn-lg btn-block">Hash</button>
</form>
</center>
</div>
<center>
<div style="font-size:18pt;height:40px;width:100%;" id="outputContainer" style="word-wrap: break-word;">
<div style="font-size:18pt;height:40px;width:100%;" id="outputBox" class="lead" style="word-wrap: break-word;"></div>
<button style="font-size:18pt;height:40px;width:100%;" type="button" id="another" class="btn btn-default btn-lg btn-block">Another</button>
<br>            <br>
</div>
</center>
</div>
<br>
<script>
document.getElementById("inputtext").focus();
</script>
</div>
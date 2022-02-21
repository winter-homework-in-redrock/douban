var wait=60;
        function time(o){
            if (wait==0) {
                o.removeAttribute("disabled");    
                o.innerHTML="获取验证码";
                wait=60;
            }else{
                o.setAttribute("disabled", true);
                o.innerHTML=wait+"秒后重新获取";
                wait--;
                setTimeout(function(){
                    time(o)
                },1000)
            }
        }
        $(".yzm").click(function(){
            time(this);
        });


// 选项卡
function openway(evt, way) {
    var i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }
    document.getElementById(way).style.display = "block";
    evt.currentTarget.className += " active";
}
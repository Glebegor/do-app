var btn_head = document.getElementsByClassName('header_btn')[0]
var header_btn_gear = document.getElementsByClassName('header_btn_gear')[0]
var header = document.getElementsByClassName('header_list')[0]

btn_head.addEventListener("click", function(){
    if(header.classList.contains("activeHeader")){
        header_btn_gear.classList.remove("activeHeaderBtn")
        header.classList.remove("activeHeader")
    }
    else{
        header_btn_gear.classList.add("activeHeaderBtn")
        header.classList.add("activeHeader")
    }
})
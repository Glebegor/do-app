var sing_in_inputs = document.getElementsByClassName("sing_in_password")[0]
var sing_in_checkbox = document.getElementsByClassName('sing_in_checkbox')[0]
sing_in_checkbox.addEventListener("click", function(){
    if(sing_in_inputs.type == "password"){
        sing_in_inputs.type = "text"
    }else if (sing_in_inputs.type == "text"){
        sing_in_inputs.type = "password"
    }
})
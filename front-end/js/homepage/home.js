var password_field = document.getElementsByClassName("home_form_form_input_password")[0]
var home_form_form_input_btn_wrapper_chechbox_input = document.getElementsByClassName('home_form_form_input_btn_wrapper_chechbox_input')[0]
home_form_form_input_btn_wrapper_chechbox_input.addEventListener("click", function(){
    if(password_field.type == "password"){
        password_field.type = "text"
    }else if (password_field.type == "text"){
        password_field.type = "password"
    }
})
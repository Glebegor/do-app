var reveal = () => {
    var reveals = document.querySelectorAll(".reveal")
    for(var i = 0; i < reveals.length; i++){
        var windowHeight = window.innerHeight;
        var elementTop = reveals[i].getBoundingClientRect().top;
        var elementVisible = 150;
        if( elementTop < windowHeight - elementVisible) {
            reveals[i].classList.add("active_rev");
        }else{
            reveals[i].classList.remove("active_rev");
        }
    }
}

window.addEventListener("scroll", reveal);

reveal();
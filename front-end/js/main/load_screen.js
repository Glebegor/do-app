
var load_wrapper = document.getElementsByClassName('load_wrapper')[0]

const onReady = (callback) => {
    var intervalId = window.setInterval(function() {
        if(load_wrapper !== undefined) {
            window.clearInterval(intervalId);
            callback.call(this);
        }
    }, 1000);
}

function setVisible(selector, visible) {
    

        document.querySelector(selector).style.opacity = visible ? '1' : '0';
    setTimeout(() => {
        document.querySelector(selector).style.display = visible ? 'block' : 'none';
    }, 300)

}
  
  onReady(function() {
    setVisible('main', true);
    setVisible('.load_wrapper', false);
  });
///
/// Back to Top
///
//Get the button
let mybutton = document.getElementById("btn-back-to-top");
// When the user scrolls down 20px from the top of the document, show the button
window.onscroll = function () {
  scrollFunction();
};
function scrollFunction() {
  if (
    document.body.scrollTop > 20 ||
    document.documentElement.scrollTop > 20
  ) {
    mybutton.style.display = "block";
  } else {
    mybutton.style.display = "none";
  }
}
// When the user clicks on the button, scroll to the top of the document
mybutton.addEventListener("click", backToTop);
function backToTop() {
  document.body.scrollTop = 0;
  document.documentElement.scrollTop = 0;
}


///
/// HTMX configuration
///
document.body.addEventListener('htmx:beforeSwap', function (evt) {
  if (evt.detail.xhr.status === 422) {
    evt.detail.shouldSwap = true;
    evt.detail.isError = false;
  } else if (evt.detail.xhr.status === 415) {
    vt.detail.shouldSwap = true;
    evt.detail.isError = false;
  }
});

///
/// Notification Toasts 
///
const toastElement = document.getElementById("toast")
const toastBody = document.getElementById("toast-body")
const toast = new bootstrap.Toast(toastElement, { delay: 2000 })
htmx.on("showMessage", (e) => {
  toastBody.innerText = e.detail.value
  toast.show()
})
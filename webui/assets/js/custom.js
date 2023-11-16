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
/// Load Images functions
///
function loadImagePng() {
  var url = "/api/images/png";
  generateGallery(url);
}
function loadImageGif() {
  var url = "/api/images/gif";
  generateGallery(url);
}
function loadImageAll() {
  var url = "/api/images";
  generateGallery(url);
}

function generateGallery(api) {
  let gallery = document.querySelector('#cardImages');
  let html = '';
  fetch(api)
    .then(response => response.json())
    .then(data => {
      data.forEach(image => {
        html +=
          `<div class="col col-3 mb-3">
            <div class="card h-100 text-center">
              <div class="card-header text-nowrap overflow-auto">
                ${image.Name}
              </div>
              <div class="ratio ratio-1x1">
              <img src="${image.Path}" class="card-img-top h-100 border-bottom border-1 rounded-0" alt="${image.Name}" />
              </div>
              <div class="card-body d-flex justify-content-center">
                <a class="btn btn-primary" >
                  Apply
                </a>
              </div>
            </div>
          </div>`;
      })
      gallery.innerHTML = html;
    })
    .catch(error => console.error(error));
}
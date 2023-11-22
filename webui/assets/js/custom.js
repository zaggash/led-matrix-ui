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
/// Load Images
///
function loadImages(img) {
  var url = "/api/images";
  switch (img) {
    case "PNG":
      url = "/api/images/png";
      break;
    case "GIF":
      url = "/api/images/gif";
      break;
  }
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
                <button class="btn btn-primary" hx-post="/api/draw" hx-trigger="click"  hx-vals='${JSON.stringify(image)}' >
                  Apply
                </button>
              </div>
            </div>
          </div>`;
      })
      gallery.innerHTML = html;
      // Re-process new content to enable htmx JS
      htmx.process(htmx.find('#cardImages'))
    })
    .catch(error => console.error(error));
}

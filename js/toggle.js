let hamburgerIsOpen = false;

function openHamburger() {
  let hamburgerNavContainer = document.getElementById("hamburger-container");

  if (!hamburgerIsOpen) {
    console.log(hamburgerIsOpen);
    hamburgerNavContainer.style.display = "flex";
    hamburgerIsOpen = true;
  } else {
    console.log(hamburgerIsOpen);
    hamburgerNavContainer.style.display = "none";
    hamburgerIsOpen = false;
  }
}

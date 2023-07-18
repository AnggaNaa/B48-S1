let hamburgerIsOpen = false;

function openHamburger() {
  let hamburgerNavContainer = document.getElementById(
    "hamburger-nav-container"
  );
  if (!hamburgerIsOpen) {
    hamburgerNavContainer.style.display = "block";
    hamburgerIsOpen = true;
    console.log(hamburgerIsOpen);
  } else {
    hamburgerNavContainer.style.display = "none";
    hamburgerIsOpen = false;
    console.log(hamburgerIsOpen);
  }
}

/* ---------------------------------------------------------------------------------------- */
/* --------------------------STYLING HUMBERGER NAV----------------------------------------- */
/* ---------------------------------------------------------------------------------------- */

// mengambil nilai dari kunci choice dari objek localStorage pada browser.
// Objek localStorage digunakan untuk menyimpan pasangan kunci-nilai secara lokal di browser web pengguna.

let choice = localStorage.getItem("choice");
const themeSwitchToggle = document.getElementById("theme-switch");

const darkMode = () => {
  document.body.classList.add("dark-mode"); //menambahkan kelas CSS "dark-mode" ke elemen <body>
  localStorage.setItem("choice", "on"); //menyimpan nilai "on" dengan kunci "choice" ke dalam objek localStorage pada browser.
}; //Metode setItem() pada objek localStorage digunakan untuk menyimpan pasangan kunci-nilai dalam penyimpanan lokal browser.

const lightMode = () => {
  document.body.classList.remove("dark-mode");
  localStorage.setItem("choice", "off");
};

if (choice === "on") {
  darkMode();
}

themeSwitchToggle.addEventListener("click", () => {
  choice = localStorage.getItem("choice");
  if (choice !== "on") {
    darkMode();
  } else {
    lightMode();
  }
});

// untuk menambahkan event listener pada elemen themeSwitchToggle. Ketika elemen tersebut diklik,
// kode akan menjalankan fungsi yang memeriksa nilai "choice" yang disimpan di localStorage dan
// menentukan apakah harus mengaktifkan mode gelap atau mode terang.

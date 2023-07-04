function emptyFormAlert() {
  let name = document.getElementById("input-name-project").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let description = document.getElementById("input-description").value;
  let multiInput = document.querySelectorAll(".multi-input:checked");
  let image = document.getElementById("input-image").value;

  if (name == "") {
    return alert("Please input your project name or title");
  } else if (startDate == "") {
    return alert("When did you start this project?");
  } else if (endDate == "") {
    return alert("When did you finish this project?");
  } else if (description == "") {
    return alert("Please describe this project.");
  } else if (multiInput.length === 0) {
    return alert("Choose your technologies.");
  } else if (image == "") {
    return alert("Please attach an image of your project.");
  }
}

let dataProject = [];

function addProject(event) {
  event.preventDefault();
  let name = document.getElementById("input-name-project").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-image").files;

  const nodeJsIcon = '<i class="fa-brands fa-node-js"></i>';
  const golangIcon = '<i class="fa-brands fa-golang"></i>';
  const reactJsIcon = '<i class="fa-brands fa-react"></i>';
  const javascriptIcon = '<i class="fa-brands fa-square-js"></i>';

  let nodejs = document.getElementById("node-js").checked ? nodeJsIcon : "";
  let golang = document.getElementById("golang").checked ? golangIcon : "";
  let reactjs = document.getElementById("react-js").checked ? reactJsIcon : "";
  let javascript = document.getElementById("javascript").checked
    ? javascriptIcon
    : "";

  let multiInput = document.querySelectorAll(".multi-input:checked");
  if (multiInput.length === 0) {
    return alert("Select at least one technology used.");
  }

  // membuat url gambar dan menampilkan gambar yg dipilih pertama
  image = URL.createObjectURL(image[0]);
  console.log(image);

  let inputStartDate = new Date(startDate);
  let inputEndDate = new Date(endDate);

  if (inputStartDate > inputEndDate) {
    return alert("Please input your dates correctly.");
  }

  let selisih = Math.abs(
    inputEndDate.getMonth() -
      inputStartDate.getMonth() +
      12 * (inputEndDate.getFullYear() - inputStartDate.getFullYear())
  );
  console.log(selisih);

  let data = {
    name,
    selisih,
    description,
    nodejs,
    golang,
    reactjs,
    javascript,
    image,
  };

  dataProject.push(data);
  console.log(dataProject);

  renderDataProject();
}

function renderDataProject() {
  document.getElementById("contents").innerHTML = "";

  for (let index = 0; index < dataProject.length; index++) {
    document.getElementById("contents").innerHTML += `
        <div class="card-project">
          <div class="image-project">
            <img src="${dataProject[index].image}" alt="" />
          </div>

          <div class="duration">
            <a href="blog.html" target="_blank"
              ><h4>${dataProject[index].name}</h4></a
            >
            <p>duration : ${dataProject[index].selisih} bulan</p>
          </div>

          <div class="description">
            <p>
              ${dataProject[index].description}
            </p>
          </div>

          <div class="icons">
            ${dataProject[index].nodejs}
            ${dataProject[index].golang}
            ${dataProject[index].reactjs}
            ${dataProject[index].javascript}
          </div>

          <div class="button-project">
            <button>edit</button>
            <button>delete</button>
          </div>
        </div>
    `;
  }
}

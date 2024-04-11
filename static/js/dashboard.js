function openNav() {
    document.getElementById("mySidenav").style.width = "360px";
}

function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
}

function filter_dropdownFunction() {
    document.getElementById("dropdown_content").classList.toggle("filter_show");
}

window.onclick = function(event) {
    if (!event.target.matches('.dropdown_button')) {
        var dropdowns = document.getElementsByClassName("dropdown_content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('filter_show')) {
                openDropdown.classList.remove('filter_show');
            }
        }
    }
}

const team_container = document.getElementById('team_container');


appendData(newData);


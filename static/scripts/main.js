(function () {
  var req =  new XMLHttpRequest();
  var id = document.querySelectorAll(".del");
  var dels = Array.from(id);
  var url = "http://localhost:3000/post/remove/";
  dels.forEach(del => {
    del.addEventListener("click", e => {
      e.preventDefault();
      var id = del.href.replace(url,"");
      if (confirm("Deseja Deletar ?")) {
        remove(id);
      }

    });
  });

  function remove(id) {
    req.open("GET",url+id,true);
    req.setRequestHeader("Content-Type", "aplication/json");
    req.send(null);
    req.addEventListener("load",function(){
      if(req.status < 300 && req.readyState === 4){
        window.location ="/";
      }
    },false);
  }

}());

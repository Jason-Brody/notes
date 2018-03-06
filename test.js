let a=false;
setInterval(function() {
    if(!a)
        console.log(1)
    a = true
  }, 0);
  new Promise(function executor(resolve) {
    console.log(2);
    for( var i=0 ; i<100000000 ; i++ ) {
      i == 9999 && resolve();
    }
    console.log(3);
  }).then(function() {
    console.log(4);
  });
  console.log(5);



function setProgress(value,id){
    var x=value;
    if (x==""){
        $(id).val(0);
        x=0;
    }
    if(x>=0 && x<=100){
        $(id).text(x.toString()+"%")
            $(id).css("width",x.toString()+"%");
    }else{
            $(id).css("width","0%");
    }
}

var $;
layui.use(['form','layer'],function(){
    var form = layui.form
        layer = parent.layer === undefined ? layui.layer : top.layer,
        $ = layui.jquery;

    form.on("submit(addUser)",function(data){
        //弹出loading
        var index = top.layer.msg('数据提交中，请稍候',{icon: 16,time:false,shade:0.8});
        // ajax post demo
        // $.post("上传路径",{
        //     userName : $(".userName").val(),
        //     userEmail : $(".userEmail").val(),
        //     userSex : data.field.sex,
        //     userGrade : data.field.userGrade,
        //     userStatus : data.field.userStatus,
        //     newsTime : submitTime,
        //     userDesc : $(".userDesc").text(),
        // },function(res){
        //
        // })
        var newOrder = "";
        newOrder += '{"exName":"'+ $(".exName").val() +'",';  // exchange name
        newOrder += '"tradingPair":"'+ $(".tradingPair").val() +'",';
        newOrder += '"orderAmount":'+ $(".orderAmount").val() +',';
        newOrder += '"orderId":'+ $(".orderId").val() +',';
        newOrder += '"buyPrice":'+ $(".price").val() +',';
        newOrder += '"sellPrice":'+ $(".price").val() +',';
        newOrder += '"orderType":'+ $(".orderType").val() +'}';
        post_url = data.field.targetUrl;
        console.log("get post url:" + post_url);
        console.log("new order:" + newOrder);
        $.ajax({
            type: 'POST',
            url: post_url,
            data: newOrder,
            dataType: 'JSON',
            success: function(res) {
                console.log("result:" + JSON.stringify(res));
                console.log(res.msg)

                top.layer.close(index);
                if (res.msg != 'ok') {
                    top.layer.msg("添加失败！");
                } else {
                    top.layer.msg("添加成功！");
                }
                 layer.closeAll("iframe");
                 //刷新父页面
                 parent.location.reload();
            },
            error: function(data) {
                alert("create new item error:" + JSON.stringify(data));
                top.layer.close(index);
                top.layer.msg("添加失败！");
                 layer.closeAll("iframe");
                 //刷新父页面
                 parent.location.reload();
            },
        });
        return false;
    })

    //格式化时间
    function filterTime(val){
        if(val < 10){
            return "0" + val;
        }else{
            return val;
        }
    }
    //定时发布
    var time = new Date();
    var submitTime = time.getFullYear()+'-'+filterTime(time.getMonth()+1)+'-'+filterTime(time.getDate())+' '+filterTime(time.getHours())+':'+filterTime(time.getMinutes())+':'+filterTime(time.getSeconds());

})

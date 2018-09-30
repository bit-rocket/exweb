layui.use(['form','layer','table','laytpl'],function(){
    var form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer,
        $ = layui.jquery,
        laytpl = layui.laytpl,
        table = layui.table;

    var digestTypeDict = {
            "1": "2buy",
            "2": "2sell"
    };
    // digest order list
    var tableIns = table.render({
        elem: '#digestList',
        url : '/order/digest',
        cellMinWidth : 95,
        page : true,
        height : "full-125",
        limits : [10,15,20,25],
        limit : 20,
        id : "digestListTable",
        cols : [[
            {type: "checkbox", fixed:"left", width:50},
            {field: 'orderId', title: 'orderId', align:"center", width:80},
            {field: 'parentId', title: 'pId', align:"center", width:80},
            {field: 'exName', title: 'ex', align:"center", width:80},
            {field: 'tradingPair', title: 'pair', align:'center', width:100},
            {field: 'orderAmount', title: 'orderAmount', align:'center', width:100},
            {field: 'dealAmount', title: 'dealAmount', align:'center', width:100},
            {field: 'price', title: 'price', align:'center', width:130},
            {field: 'orderType', title: 'type',  align:'center', width:100,
            templet: function(d) {
                return digestTypeDict[d.orderType];
            }},
            {field: 'status', title: 'status', align:'center', width:100},
            {field: 'create', title: 'create / finish', align:'center', minWidth:350,
            templet: function(d){
                return "" + d.createTime + " / " + d.finishTime;
            }},
            {title: 'op', templet:'#digestListBar',fixed:"right",align:"center", minWidth:170}
        ]]
    });

    //搜索【此功能需要后台配合，所以暂时没有动态效果演示】
    $(".search_btn").on("click",function(){
        if($(".searchVal").val() != ''){
            table.reload("newsListTable",{
                page: {
                    curr: 1 //重新从第 1 页开始
                },
                where: {
                    key: $(".searchVal").val()  //搜索的关键字
                }
            })
        }else{
            layer.msg("请输入搜索的内容");
        }
    });

    //添加用户
    function addUser(edit){
        var index = layui.layer.open({
            title : "添加用户",
            type : 2,
            content : "userAdd.html",
            success : function(layero, index){
                var body = layui.layer.getChildFrame('body', index);
                if(edit){
                    body.find(".exName").val(edit.exName);
                    body.find(".tradingPair").val(edit.tradingPair);
                    body.find(".orderAmount").val(edit.holding);
                    body.find(".price").val(edit.buyPrice);
                    body.find(".targetUrl").val("/order/digest");
                    body.find(".orderId").val(edit.orderId);
                    // $("#maxOrderAmount").innerHtml("max:" + edit.holding);
                    body.find(".maxOrderAmount").text("max:" + edit.holding);
                    body.find(".addUser").text("digest order");
                    // or body.find(".addUser").html("digest order");

                    form.render();
                }
                setTimeout(function(){
                    layui.layer.tips('点击此处返回订单列表', '.layui-layer-setwin .layui-layer-close', {
                        tips: 3
                    });
                },500)
            }
        })
        layui.layer.full(index);
        window.sessionStorage.setItem("index",index);
        //改变窗口大小时，重置弹窗的宽高，防止超出可视区域（如F12调出debug的操作）
        $(window).on("resize",function(){
            layui.layer.full(window.sessionStorage.getItem("index"));
        })
    }
    $(".addNews_btn").click(function(){
        addUser();
    })

    //批量删除
    $(".delAll_btn").click(function(){
        var checkStatus = table.checkStatus('userListTable'),
            data = checkStatus.data,
            newsId = [];
        if(data.length > 0) {
            for (var i in data) {
                newsId.push(data[i].newsId);
            }
            layer.confirm('确定删除选中的用户？', {icon: 3, title: '提示信息'}, function (index) {
                // $.get("删除文章接口",{
                //     newsId : newsId  //将需要删除的newsId作为参数传入
                // },function(data){
                tableIns.reload();
                layer.close(index);
                // })
            })
        }else{
            layer.msg("请选择需要删除的用户");
        }
    })

    //列表操作
    table.on('tool(digestList)', function(obj){
        var layEvent = obj.event,
            data = obj.data;

        if(layEvent === 'edit'){ //编辑
            addUser(data);
        }else if(layEvent === 'usable'){ //启用禁用
            var _this = $(this),
                usableText = "是否确定禁用此用户？",
                btnText = "已禁用";
            if(_this.text()=="已禁用"){
                usableText = "是否确定启用此用户？",
                btnText = "已启用";
            }
            layer.confirm(usableText,{
                icon: 3,
                title:'系统提示',
                cancel : function(index){
                    layer.close(index);
                }
            },function(index){
                _this.text(btnText);
                layer.close(index);
            },function(index){
                layer.close(index);
            });
        }else if(layEvent === 'del'){ //删除
            layer.confirm('确定删除此用户？',{icon:3, title:'提示信息'},function(index){
                // $.get("删除文章接口",{
                //     newsId : data.newsId  //将需要删除的newsId作为参数传入
                // },function(data){
                    tableIns.reload();
                    layer.close(index);
                // })
            });
        }
    });

})

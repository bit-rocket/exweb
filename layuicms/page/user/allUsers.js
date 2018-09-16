layui.config({
    base : "js/"
}).use(['form','layer','jquery','laypage'],function(){
    var form = layui.form(),
        layer = parent.layer === undefined ? layui.layer : parent.layer,
        laypage = layui.laypage,
        $ = layui.jquery;

    //加载页面数据
    var ordersData = '';
    $.get("/order/todeal", function(data){
        ordersData = data;
        // load order data
        ordersList();
    })

    //查询
    $(".search_btn").click(function(){
        var userArray = [];
        if($(".search_input").val() != ''){
            var index = layer.msg('查询中，请稍候',{icon: 16,time:false,shade:0.8});
            setTimeout(function(){
                $.ajax({
                    url : "/order/todeal",
                    type : "get",
                    dataType : "json",
                    success : function(data){
                        ordersData = data;
                        for(var i=0;i<ordersData.length;i++){
                            var ordersStr = ordersData[i];
                            var selectStr = $(".search_input").val();
                            function changeStr(data){
                                var dataStr = '';
                                var showNum = data.split(eval("/"+selectStr+"/ig")).length - 1;
                                if(showNum > 1){
                                    for (var j=0;j<showNum;j++) {
                                        dataStr += data.split(eval("/"+selectStr+"/ig"))[j] + "<i style='color:#03c339;font-weight:bold;'>" + selectStr + "</i>";
                                    }
                                    dataStr += data.split(eval("/"+selectStr+"/ig"))[showNum];
                                    return dataStr;
                                }else{
                                    dataStr = data.split(eval("/"+selectStr+"/ig"))[0] + "<i style='color:#03c339;font-weight:bold;'>" + selectStr + "</i>" + data.split(eval("/"+selectStr+"/ig"))[1];
                                    return dataStr;
                                }
                            }
                            // exchange name
                            if(ordersStr.exName.indexOf(selectStr) > -1){
                                ordersStr["exName"] = changeStr(ordersStr.exName);
                            }
                            // trading pair
                            if(ordersStr.tradingPair.indexOf(selectStr) > -1){
                                ordersStr["tradingPair"] = changeStr(ordersStr.tradingPair);
                            }
                            // buyPrice
                            if(ordersStr.buyPrice.indexOf(selectStr) > -1){
                                ordersStr["buyPrice"] = changeStr(ordersStr.buyPrice);
                            }
                            // earn percent
                            if(ordersStr.earnRate.indexOf(selectStr) > -1){
                                ordersStr["earnRate"] = changeStr(ordersStr.earnRate);
                            }
                            if(ordersStr.exName.indexOf(selectStr)>-1 || ordersStr.tradingPair.indexOf(selectStr)>-1 || ordersStr.buyPrice.indexOf(selectStr)>-1 || ordersStr.earnRate.indexOf(selectStr)>-1){
                                userArray.push(ordersStr);
                            }
                        }
                        ordersData = userArray;
                        ordersList(ordersData);
                    }
                })

                layer.close(index);
            },2000);
        }else{
            layer.msg("请输入需要查询的内容");
        }
    })

    //添加会员
    $(".usersAdd_btn").click(function(){
        var index = layui.layer.open({
            title : "新订单",
            type : 2,
            content : "addUser.html",
            success : function(layero, index){
                setTimeout(function(){
                    layui.layer.tips('点击此处返回会员列表', '.layui-layer-setwin .layui-layer-close', {
                        tips: 3
                    });
                },500)
            }
        })
        //改变窗口大小时，重置弹窗的高度，防止超出可视区域（如F12调出debug的操作）
        $(window).resize(function(){
            layui.layer.full(index);
        })
        layui.layer.full(index);
    })

    //批量删除
    $(".batchDel").click(function(){
        var $checkbox = $('.users_list tbody input[type="checkbox"][name="checked"]');
        var $checked = $('.users_list tbody input[type="checkbox"][name="checked"]:checked');
        if($checkbox.is(":checked")){
            layer.confirm('确定删除选中的信息？',{icon:3, title:'提示信息'},function(index){
                var index = layer.msg('删除中，请稍候',{icon: 16,time:false,shade:0.8});
                setTimeout(function(){
                    //删除数据
                    for(var j=0;j<$checked.length;j++){
                        for(var i=0;i<ordersData.length;i++){
                            if(ordersData[i].newsId == $checked.eq(j).parents("tr").find(".news_del").attr("data-id")){
                                ordersData.splice(i,1);
                                ordersList(ordersData);
                            }
                        }
                    }
                    $('.users_list thead input[type="checkbox"]').prop("checked",false);
                    form.render();
                    layer.close(index);
                    layer.msg("删除成功");
                },2000);
            })
        }else{
            layer.msg("请选择需要删除的文章");
        }
    })

    //全选
    form.on('checkbox(allChoose)', function(data){
        var child = $(data.elem).parents('table').find('tbody input[type="checkbox"]:not([name="show"])');
        child.each(function(index, item){
            item.checked = data.elem.checked;
        });
        form.render('checkbox');
    });

    //通过判断文章是否全部选中来确定全选按钮是否选中
    form.on("checkbox(choose)",function(data){
        var child = $(data.elem).parents('table').find('tbody input[type="checkbox"]:not([name="show"])');
        var childChecked = $(data.elem).parents('table').find('tbody input[type="checkbox"]:not([name="show"]):checked')
        if(childChecked.length == child.length){
            $(data.elem).parents('table').find('thead input#allChoose').get(0).checked = true;
        }else{
            $(data.elem).parents('table').find('thead input#allChoose').get(0).checked = false;
        }
        form.render('checkbox');
    })

    //操作
    $("body").on("click",".users_edit",function(){  //编辑
        // layer.alert('您点击了会员编辑按钮，由于是纯静态页面，所以暂时不存在编辑内容，后期会添加，敬请谅解。。。',{icon:6, title:'文章编辑'});
        var index = layui.layer.open({
            title : "新订单",
            type : 2,
            content : "addUser.html",
            success : function(layero, index){
                setTimeout(function(){
                    layui.layer.tips('点击此处返回会员列表', '.layui-layer-setwin .layui-layer-close', {
                        tips: 3
                    });
                },500)
            }
        })
        //改变窗口大小时，重置弹窗的高度，防止超出可视区域（如F12调出debug的操作）
        $(window).resize(function(){
            layui.layer.full(index);
        })
        layui.layer.full(index);
    })

    $("body").on("click",".users_del",function(){  //删除
        var _this = $(this);
        layer.confirm('确定删除此用户？',{icon:3, title:'提示信息'},function(index){
            //_this.parents("tr").remove();
            for(var i=0;i<ordersData.length;i++){
                if(ordersData[i].orderId == _this.attr("data-id")){
                    ordersData.splice(i,1);
                    ordersList(ordersData);
                }
            }
            layer.close(index);
        });
    })

    function ordersList(){
        //渲染数据
        function renderDate(data,curr){
            var dataHtml = '';
            currData = ordersData.concat().splice(curr*nums-nums, nums);
            if(currData.length != 0){
                for(var i=0;i<currData.length;i++){
                    dataHtml += '<tr>'
                    +  '<td><input type="checkbox" name="checked" lay-skin="primary" lay-filter="choose"></td>'
                    +  '<td>'+currData[i].exName+'</td>'
                    +  '<td>'+currData[i].tradingPair+'</td>'
                    +  '<td>'+currData[i].orderAmount+'</td>'
                    +  '<td>'+currData[i].orderType+'</td>'
                    +  '<td>'+currData[i].holding+'</td>'
                    +  '<td>'+currData[i].buyPrice+'</td>'
                    +  '<td>'+currData[i].sellPrice+'</td>'
                    +  '<td>'+currData[i].earnRate+ '/' + currData[i].earnAmount + '</td>'
                    +  '<td>'+currData[i].status+'</td>'
                    +  '<td>'+currData[i].createTime+'</td>'
                    +  '<td>'+currData[i].finishTime+'</td>'
                    +  '<td>'
                    +    '<a class="layui-btn layui-btn-mini users_edit"><i class="iconfont icon-edit"></i> 编辑</a>'
                    +    '<a class="layui-btn layui-btn-danger layui-btn-mini users_del" data-id="'
                    +   data[i].orderId+'"><i class="layui-icon">&#xe640;</i> 删除</a>'
                    +  '</td>'
                    +'</tr>';
                }
            }else{
                dataHtml = '<tr><td colspan="8">暂无数据</td></tr>';
            }
            return dataHtml;
        }

        //分页
        var nums = 13; //每页出现的数据量
        laypage({
            cont : "page",
            pages : Math.ceil(ordersData.length/nums),
            jump : function(obj){
                $(".users_content").html(renderDate(ordersData,obj.curr));
                $('.users_list thead input[type="checkbox"]').prop("checked",false);
                form.render();
            }
        })
    }

})

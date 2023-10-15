async function setArtileList() {
    // 1.2 获取文章列表数据
    var res = await axios({
      url: 'http://47.96.100.92:8080/api/public',
    })
    console.log(res)
    // 1.3 展示到指定的标签结构中
    for(let i=0;i<res.data.data.text_list.length;i++){
        if(res.data.data.text_list[i].unname!='0'){
            var htmlStr = res.data.data.text_list.map(function(ele){
                return`
                <ul>
                    <li>匿名:</li>
                    <li>${ele.text}</li>
                </ul>
                `})
        }else if(res.data.data.text_list[i].unname!='1'){
            var htmlStr = res.data.data.text_list.map(function(ele){
                return`
                <div>
                    <span>${ele.name}</span>
                    <span>${ele.text}</span>
                </div>
                `})
        }
        document.querySelector('.wall-list').innerHTML = htmlStr
    }
  }
  setArtileList()


// 获取并设置文章列表
async function setList() {
    // 1.2 获取文章列表数据
    var res = await axios({
      url: 'http://47.96.100.92:8080/api/personal',
    })
    console.log(res)
    // 1.3 展示到指定的标签结构中
    const htmlStr = res.data.results.map(item => `<tr>
    <td>${item.unname}</td>
    <td>
      <span>${item.text}</span>
    </td>
    <td data-id="${item.id}">
      <i class="bi bi-pencil-square edit"></i>
      <i class="bi bi-trash3 del"></i>
    </td>
  </tr>`).join('')
    document.querySelector('.art-list').innerHTML = htmlStr
  }
  setList()





document.querySelector('.art-list').addEventListener('click',async e=>{
    if(e.target.classList.contains('del')){
        const delId=e.target.parentNode.dataset.delId
        const res =await axios({
            url:`http://47.96.100.92:8080/api/personal${delId}`,
            method:'DELETE'
        })
        setList()
    }
})




document.querySelector('.art-list').addEventListener('click',e=>{
    if(e.target.classList.contains('edit')){
        const artId=e.target.parentNode.dataset.id
        location.href= `./index.html?id=${artId}`
    }
})

;(function(){
    const paramsStr=location.search
    const params=new URLSearchParams(paramsStr)
    params.forEach(async (value,key)=>{
        if(key==='id'){
            document.querySelector('.title span').innerHTML='修改表白'
            document.querySelector('.send').innerHTML='修改'
            const res=axios({
                url:`http://47.96.100.92:8080/api/personal/${value}`
            })
            console.log(res)
            const dataObj={
                name:res.data.name,
                text:res.data.text,
                unname:res.data.unname,
                id:res.data.id
            }
            Object.keys(dataObj).forEach(key=>{
                if(key==='text'){
                    editor.setHtml(dataObj[key])
                }else{
                    document.querySelector(`[name=${key}]`).value=dataObj[key]
                }
            })
        }
    })
})

document.querySelector('.send').addEventListener('click',e=>{
    if(e.target.innerHTML!=='修改') return
    const text=document.querySelector('.publish-content').value
    const name=document.querySelector('.form-control').value
    const unname=document.querySelector('.unname').value
            // const privacy=document.querySelector('').value
    console.log(text)
    axios({
        url:`http://47.96.100.92:8080/api/personal/${data.id}`,
        method:'post',
        data:{
            name,
            text,
            unname
                    // privacy
        }
    }).then(result=>{
        console.log(result)
        alert('修改成功')
    }).catch(error=>{
        alert('修改失败')
    })
})
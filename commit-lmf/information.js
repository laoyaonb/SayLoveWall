<script type="text/javascript" src="form-serialize.js"></script>
const creator=
axios({
    url:'http://http://47.96.100.92:8080/api/settings',
    params:{
        creator:''
    }
}).then(result=>{
    const userobj=result.data.data
    Object.keys(userobj).forEach(key=>{
        if(key==='gender'){
            const gradiolist=document.querySelectorAll('.gender')
            const gnum=userobj[key]
            gradiolist[gnum].checked=true
        }
        else {
            document.querySelector(`.${key}`).value=userobj[key]
        }
    })
})

document.querySelector('.sumbit').addEventListener('click',()=>{
    const userform=document.querySelector('.form')
    const userinformation=serialize(userform,{hash:true,empty:true})
    userinformation.creator=creator
    userinformation.gender=+userinformation.gender
    console.log(userinformation)
    axios({
        url:'http://http://47.96.100.92:8080/api/settings',
        method:'put',
        data:userinformation
    }).then(result=>{
        alert("操作成功")
    }).catch(Error=>{
        console.log(Error)
        console.log(error.response.data.message)
        alert("操作失败")
    })
})
